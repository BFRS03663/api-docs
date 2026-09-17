package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/shiprocket/apidocs/internal/domain"
)

type operationDoc struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	CollectionID bson.ObjectID `bson:"collectionId"`
	Method       string        `bson:"method"`
	Path         string        `bson:"path"`
	OperationID  string        `bson:"operationId"`
	Summary      string        `bson:"summary"`
	Description  string        `bson:"description"`
	Tags         []string      `bson:"tags"`
	Deprecated   bool          `bson:"deprecated"`
	Order        int           `bson:"order"`
	SearchText   string        `bson:"searchText"`
}

// OperationRepo persists the flattened operation index.
type OperationRepo struct {
	col *mongo.Collection
}

// NewOperationRepo returns a repo bound to the "operations" collection.
func NewOperationRepo(db *mongo.Database) *OperationRepo {
	return &OperationRepo{col: db.Collection("operations")}
}

// EnsureIndexes creates ordering, uniqueness, and full-text indexes.
func (r *OperationRepo) EnsureIndexes(ctx context.Context) error {
	_, err := r.col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "collectionId", Value: 1}, {Key: "order", Value: 1}}},
		{
			Keys:    bson.D{{Key: "collectionId", Value: 1}, {Key: "method", Value: 1}, {Key: "path", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{Key: "summary", Value: "text"},
				{Key: "description", Value: "text"},
				{Key: "path", Value: "text"},
				{Key: "tags", Value: "text"},
				{Key: "operationId", Value: "text"},
				{Key: "searchText", Value: "text"},
			},
			Options: options.Index().SetName("operations_text").SetWeights(bson.D{
				{Key: "summary", Value: 10},
				{Key: "operationId", Value: 8},
				{Key: "path", Value: 6},
				{Key: "tags", Value: 4},
				{Key: "description", Value: 2},
				{Key: "searchText", Value: 1},
			}),
		},
	})
	return err
}

// Replace deletes every operation for the collection and inserts ops. The two
// steps are not transactional (standalone Mongo has no transactions); a reader
// racing an import may briefly see an empty list.
func (r *OperationRepo) Replace(ctx context.Context, collectionID string, ops []domain.Operation) error {
	cid, err := bson.ObjectIDFromHex(collectionID)
	if err != nil {
		return fmt.Errorf("invalid collection id: %w", err)
	}
	if _, err := r.col.DeleteMany(ctx, bson.D{{Key: "collectionId", Value: cid}}); err != nil {
		return fmt.Errorf("clear operations: %w", err)
	}
	if len(ops) == 0 {
		return nil
	}
	docs := make([]any, 0, len(ops))
	for _, op := range ops {
		docs = append(docs, operationDoc{
			CollectionID: cid,
			Method:       op.Method,
			Path:         op.Path,
			OperationID:  op.OperationID,
			Summary:      op.Summary,
			Description:  op.Description,
			Tags:         nonNil(op.Tags),
			Deprecated:   op.Deprecated,
			Order:        op.Order,
			SearchText:   op.SearchText,
		})
	}
	if _, err := r.col.InsertMany(ctx, docs); err != nil {
		return fmt.Errorf("insert operations: %w", err)
	}
	return nil
}

// ListByCollection returns a collection's operations in display order.
func (r *OperationRepo) ListByCollection(ctx context.Context, collectionID string) ([]domain.Operation, error) {
	cid, err := bson.ObjectIDFromHex(collectionID)
	if err != nil {
		return nil, fmt.Errorf("invalid collection id: %w", err)
	}
	cur, err := r.col.Find(ctx, bson.D{{Key: "collectionId", Value: cid}}, options.Find().SetSort(bson.D{{Key: "order", Value: 1}}))
	if err != nil {
		return nil, fmt.Errorf("list operations: %w", err)
	}
	var docs []operationDoc
	if err := cur.All(ctx, &docs); err != nil {
		return nil, fmt.Errorf("decode operations: %w", err)
	}
	out := make([]domain.Operation, 0, len(docs))
	for _, d := range docs {
		out = append(out, domain.Operation{
			ID:           d.ID.Hex(),
			CollectionID: d.CollectionID.Hex(),
			Method:       d.Method,
			Path:         d.Path,
			OperationID:  d.OperationID,
			Summary:      d.Summary,
			Description:  d.Description,
			Tags:         nonNil(d.Tags),
			Deprecated:   d.Deprecated,
			Order:        d.Order,
			SearchText:   d.SearchText,
		})
	}
	return out, nil
}

// DeleteByCollection removes all operations for a collection.
func (r *OperationRepo) DeleteByCollection(ctx context.Context, collectionID string) error {
	cid, err := bson.ObjectIDFromHex(collectionID)
	if err != nil {
		return fmt.Errorf("invalid collection id: %w", err)
	}
	if _, err := r.col.DeleteMany(ctx, bson.D{{Key: "collectionId", Value: cid}}); err != nil {
		return fmt.Errorf("delete operations: %w", err)
	}
	return nil
}

// Search runs a $text query over the operation index, most relevant first.
// collectionID narrows the search when non-empty. CollectionSlug/Name are
// left for the caller to fill in.
func (r *OperationRepo) Search(ctx context.Context, q, collectionID string, limit int) ([]domain.SearchHit, error) {
	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: q}}}}
	if collectionID != "" {
		cid, err := bson.ObjectIDFromHex(collectionID)
		if err != nil {
			return nil, fmt.Errorf("invalid collection id: %w", err)
		}
		filter = append(filter, bson.E{Key: "collectionId", Value: cid})
	}
	score := bson.D{{Key: "$meta", Value: "textScore"}}
	// A $meta projection switches Mongo into inclusion mode, so every field
	// we decode has to be listed; searchText is deliberately left out.
	opts := options.Find().
		SetProjection(bson.D{
			{Key: "score", Value: score},
			{Key: "collectionId", Value: 1}, {Key: "method", Value: 1}, {Key: "path", Value: 1},
			{Key: "operationId", Value: 1}, {Key: "summary", Value: 1}, {Key: "description", Value: 1},
			{Key: "tags", Value: 1}, {Key: "deprecated", Value: 1}, {Key: "order", Value: 1},
		}).
		SetSort(bson.D{{Key: "score", Value: score}, {Key: "order", Value: 1}}).
		SetLimit(int64(limit))
	cur, err := r.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("search operations: %w", err)
	}
	// The embedded document must be an exported field or the decoder skips it.
	var docs []struct {
		Doc   operationDoc `bson:",inline"`
		Score float64      `bson:"score"`
	}
	if err := cur.All(ctx, &docs); err != nil {
		return nil, fmt.Errorf("decode search: %w", err)
	}
	out := make([]domain.SearchHit, 0, len(docs))
	for _, hit := range docs {
		d := hit.Doc
		out = append(out, domain.SearchHit{
			Operation: domain.Operation{
				ID:           d.ID.Hex(),
				CollectionID: d.CollectionID.Hex(),
				Method:       d.Method,
				Path:         d.Path,
				OperationID:  d.OperationID,
				Summary:      d.Summary,
				Description:  d.Description,
				Tags:         nonNil(d.Tags),
				Deprecated:   d.Deprecated,
				Order:        d.Order,
			},
			Score: hit.Score,
		})
	}
	return out, nil
}
