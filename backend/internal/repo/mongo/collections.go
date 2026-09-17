package mongo

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/shiprocket/apidocs/internal/domain"
)

// collectionDoc is the BSON shape of a collection. Spec and RawUpload are
// gzip-compressed JSON so arbitrary OpenAPI keys ($ref, dotted media types)
// never hit BSON field-name rules and large specs stay well under 16 MB.
type collectionDoc struct {
	ID             bson.ObjectID `bson:"_id,omitempty"`
	Slug           string        `bson:"slug"`
	Name           string        `bson:"name"`
	Description    string        `bson:"description"`
	Version        string        `bson:"version"`
	Source         sourceDoc     `bson:"source"`
	Servers        []string      `bson:"servers"`
	OperationCount int           `bson:"operationCount"`
	Spec           []byte        `bson:"spec,omitempty"`
	RawUpload      []byte        `bson:"rawUpload,omitempty"`
	CreatedAt      time.Time     `bson:"createdAt"`
	UpdatedAt      time.Time     `bson:"updatedAt"`
}

type sourceDoc struct {
	Type       string    `bson:"type"`
	Filename   string    `bson:"filename"`
	ImportedAt time.Time `bson:"importedAt"`
}

// metadataProjection excludes the large blobs for list/get calls.
var metadataProjection = bson.D{{Key: "spec", Value: 0}, {Key: "rawUpload", Value: 0}}

// CollectionRepo persists collections.
type CollectionRepo struct {
	col *mongo.Collection
}

// NewCollectionRepo returns a repo bound to the "collections" collection.
func NewCollectionRepo(db *mongo.Database) *CollectionRepo {
	return &CollectionRepo{col: db.Collection("collections")}
}

// EnsureIndexes creates the unique slug index. Safe to call on every startup.
func (r *CollectionRepo) EnsureIndexes(ctx context.Context) error {
	_, err := r.col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "slug", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

// Upsert inserts or replaces the collection identified by c.Slug and fills
// c.ID, c.CreatedAt and c.UpdatedAt. It reports whether a new record was created.
func (r *CollectionRepo) Upsert(ctx context.Context, c *domain.Collection) (bool, error) {
	spec, err := gzipBytes(c.Spec)
	if err != nil {
		return false, err
	}
	raw, err := gzipBytes(c.RawUpload)
	if err != nil {
		return false, err
	}
	now := time.Now().UTC().Truncate(time.Millisecond)

	set := bson.D{
		{Key: "name", Value: c.Name},
		{Key: "description", Value: c.Description},
		{Key: "version", Value: c.Version},
		{Key: "source", Value: sourceDoc{Type: string(c.Source.Type), Filename: c.Source.Filename, ImportedAt: c.Source.ImportedAt}},
		{Key: "servers", Value: nonNil(c.Servers)},
		{Key: "operationCount", Value: c.OperationCount},
		{Key: "spec", Value: spec},
		{Key: "rawUpload", Value: raw},
		{Key: "updatedAt", Value: now},
	}
	update := bson.D{
		{Key: "$set", Value: set},
		{Key: "$setOnInsert", Value: bson.D{{Key: "slug", Value: c.Slug}, {Key: "createdAt", Value: now}}},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After).SetProjection(metadataProjection)

	var doc collectionDoc
	if err := r.col.FindOneAndUpdate(ctx, bson.D{{Key: "slug", Value: c.Slug}}, update, opts).Decode(&doc); err != nil {
		return false, fmt.Errorf("upsert collection: %w", err)
	}
	c.ID = doc.ID.Hex()
	c.CreatedAt = doc.CreatedAt
	c.UpdatedAt = doc.UpdatedAt
	return doc.CreatedAt.Equal(now), nil
}

// List returns every collection's metadata sorted by name.
func (r *CollectionRepo) List(ctx context.Context) ([]domain.Collection, error) {
	opts := options.Find().SetProjection(metadataProjection).SetSort(bson.D{{Key: "name", Value: 1}})
	cur, err := r.col.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, fmt.Errorf("list collections: %w", err)
	}
	var docs []collectionDoc
	if err := cur.All(ctx, &docs); err != nil {
		return nil, fmt.Errorf("decode collections: %w", err)
	}
	out := make([]domain.Collection, 0, len(docs))
	for i := range docs {
		out = append(out, toDomain(&docs[i]))
	}
	return out, nil
}

// GetBySlug returns one collection's metadata.
func (r *CollectionRepo) GetBySlug(ctx context.Context, slug string) (*domain.Collection, error) {
	var doc collectionDoc
	err := r.col.FindOne(ctx, bson.D{{Key: "slug", Value: slug}}, options.FindOne().SetProjection(metadataProjection)).Decode(&doc)
	if err != nil {
		return nil, mapErr("get collection", err)
	}
	c := toDomain(&doc)
	return &c, nil
}

// GetSpec returns the canonical OpenAPI JSON for a slug.
func (r *CollectionRepo) GetSpec(ctx context.Context, slug string) ([]byte, error) {
	return r.blob(ctx, slug, "spec")
}

// GetRawUpload returns the originally uploaded bytes for a slug.
func (r *CollectionRepo) GetRawUpload(ctx context.Context, slug string) ([]byte, error) {
	return r.blob(ctx, slug, "rawUpload")
}

func (r *CollectionRepo) blob(ctx context.Context, slug, field string) ([]byte, error) {
	var doc bson.M
	opts := options.FindOne().SetProjection(bson.D{{Key: field, Value: 1}})
	if err := r.col.FindOne(ctx, bson.D{{Key: "slug", Value: slug}}, opts).Decode(&doc); err != nil {
		return nil, mapErr("get "+field, err)
	}
	raw, _ := doc[field].(bson.Binary)
	return gunzipBytes(raw.Data)
}

// Delete removes a collection and returns its id so callers can cascade.
func (r *CollectionRepo) Delete(ctx context.Context, slug string) (string, error) {
	var doc collectionDoc
	opts := options.FindOneAndDelete().SetProjection(bson.D{{Key: "_id", Value: 1}})
	if err := r.col.FindOneAndDelete(ctx, bson.D{{Key: "slug", Value: slug}}, opts).Decode(&doc); err != nil {
		return "", mapErr("delete collection", err)
	}
	return doc.ID.Hex(), nil
}

func toDomain(d *collectionDoc) domain.Collection {
	return domain.Collection{
		ID:          d.ID.Hex(),
		Slug:        d.Slug,
		Name:        d.Name,
		Description: d.Description,
		Version:     d.Version,
		Source: domain.ImportSource{
			Type:       domain.SourceType(d.Source.Type),
			Filename:   d.Source.Filename,
			ImportedAt: d.Source.ImportedAt,
		},
		Servers:        nonNil(d.Servers),
		OperationCount: d.OperationCount,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
	}
}

func mapErr(op string, err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return domain.ErrNotFound
	}
	return fmt.Errorf("%s: %w", op, err)
}

func nonNil(s []string) []string {
	if s == nil {
		return []string{}
	}
	return s
}

func gzipBytes(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func gunzipBytes(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, domain.ErrNotFound
	}
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("gunzip: %w", err)
	}
	defer r.Close()
	return io.ReadAll(r)
}
