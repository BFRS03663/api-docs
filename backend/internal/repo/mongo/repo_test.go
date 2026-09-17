package mongo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/shiprocket/apidocs/internal/domain"
)

// testDB connects to TEST_MONGO_URI and returns a throwaway database that is
// dropped when the test ends. Tests are skipped when the variable is unset.
func testDB(t *testing.T) (context.Context, *Client) {
	t.Helper()
	uri := os.Getenv("TEST_MONGO_URI")
	if uri == "" {
		t.Skip("TEST_MONGO_URI not set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	t.Cleanup(cancel)
	name := fmt.Sprintf("apidocs_test_%d", time.Now().UnixNano())
	c, err := Connect(ctx, uri, name)
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	t.Cleanup(func() {
		_ = c.DB.Drop(context.Background())
		_ = c.Close(context.Background())
	})
	return ctx, c
}

func TestCollectionAndOperationRoundTrip(t *testing.T) {
	ctx, c := testDB(t)
	cols := NewCollectionRepo(c.DB)
	ops := NewOperationRepo(c.DB)
	if err := cols.EnsureIndexes(ctx); err != nil {
		t.Fatalf("collection indexes: %v", err)
	}
	if err := ops.EnsureIndexes(ctx); err != nil {
		t.Fatalf("operation indexes: %v", err)
	}

	spec := []byte(`{"openapi":"3.0.0","info":{"title":"T","version":"1"},"paths":{"/a":{"get":{"responses":{"200":{"description":"ok"}}}}}}`)
	col := &domain.Collection{
		Slug: "demo", Name: "Demo", Version: "1",
		Source:    domain.ImportSource{Type: domain.SourceOpenAPI3, Filename: "demo.json", ImportedAt: time.Now().UTC()},
		Spec:      spec,
		RawUpload: []byte("raw"),
	}
	created, err := cols.Upsert(ctx, col)
	if err != nil {
		t.Fatalf("upsert: %v", err)
	}
	if !created || col.ID == "" {
		t.Fatalf("first upsert should create; created=%v id=%q", created, col.ID)
	}
	firstID := col.ID

	col.Name = "Demo v2"
	created, err = cols.Upsert(ctx, col)
	if err != nil {
		t.Fatalf("second upsert: %v", err)
	}
	if created || col.ID != firstID {
		t.Fatalf("second upsert should update in place; created=%v id=%q want %q", created, col.ID, firstID)
	}

	list, err := cols.List(ctx)
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(list) != 1 || list[0].Name != "Demo v2" || list[0].Spec != nil {
		t.Fatalf("list = %+v", list)
	}

	got, err := cols.GetSpec(ctx, "demo")
	if err != nil {
		t.Fatalf("get spec: %v", err)
	}
	if string(got) != string(spec) {
		t.Fatalf("spec round trip mismatch: %s", got)
	}
	raw, err := cols.GetRawUpload(ctx, "demo")
	if err != nil || string(raw) != "raw" {
		t.Fatalf("raw upload = %q, %v", raw, err)
	}

	if err := ops.Replace(ctx, col.ID, []domain.Operation{
		{Method: "GET", Path: "/a", OperationID: "getA", Summary: "Get A", Order: 0},
		{Method: "POST", Path: "/a", OperationID: "postA", Summary: "Post A", Order: 1},
	}); err != nil {
		t.Fatalf("replace ops: %v", err)
	}
	if err := ops.Replace(ctx, col.ID, []domain.Operation{
		{Method: "GET", Path: "/a", OperationID: "getA", Summary: "Get A", Order: 0},
	}); err != nil {
		t.Fatalf("second replace must not hit unique index: %v", err)
	}
	listed, err := ops.ListByCollection(ctx, col.ID)
	if err != nil {
		t.Fatalf("list ops: %v", err)
	}
	if len(listed) != 1 || listed[0].OperationID != "getA" || listed[0].Tags == nil {
		t.Fatalf("ops = %+v", listed)
	}

	if err := ops.Replace(ctx, col.ID, []domain.Operation{
		{Method: "GET", Path: "/pet/findByStatus", OperationID: "findPetsByStatus", Summary: "Finds Pets by status", Tags: []string{"pet"}, Order: 0},
		{Method: "POST", Path: "/store/order", OperationID: "placeOrder", Summary: "Place an order", SearchText: "purchase a pet", Order: 1},
	}); err != nil {
		t.Fatalf("replace for search: %v", err)
	}
	hits, err := ops.Search(ctx, "status", "", 10)
	if err != nil {
		t.Fatalf("search: %v", err)
	}
	if len(hits) != 1 || hits[0].OperationID != "findPetsByStatus" || hits[0].Score <= 0 {
		t.Fatalf("search hits = %+v", hits)
	}
	if hits, _ := ops.Search(ctx, "purchase", col.ID, 10); len(hits) != 1 || hits[0].OperationID != "placeOrder" {
		t.Fatalf("searchText not indexed: %+v", hits)
	}
	if hits, _ := ops.Search(ctx, "zebra", "", 10); len(hits) != 0 {
		t.Fatalf("unexpected hits: %+v", hits)
	}

	if _, err := cols.GetBySlug(ctx, "missing"); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("missing slug err = %v", err)
	}

	id, err := cols.Delete(ctx, "demo")
	if err != nil || id != firstID {
		t.Fatalf("delete = %q, %v", id, err)
	}
	if err := ops.DeleteByCollection(ctx, id); err != nil {
		t.Fatalf("delete ops: %v", err)
	}
	if _, err := cols.GetSpec(ctx, "demo"); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("after delete err = %v", err)
	}
}
