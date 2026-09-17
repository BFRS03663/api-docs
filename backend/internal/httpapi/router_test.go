package httpapi

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/shiprocket/apidocs/internal/auth"
	"github.com/shiprocket/apidocs/internal/config"
	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer/source"
	"github.com/shiprocket/apidocs/internal/importsvc"
)

type fakePinger struct{ err error }

func (f fakePinger) Ping(context.Context) error { return f.err }

const testPassword = "correct horse battery"

func newTestRouter(p Pinger) http.Handler {
	return newTestRouterWith(p, &fakeCols{}, &fakeOps{})
}

func newTestRouterWith(p Pinger, cols CollectionStore, ops OperationStore) http.Handler {
	return newAdminRouter(p, cols, ops, true)
}

// newAdminRouter builds the full router with a real auth service (admin /
// testPassword) and an import service whose fetcher may or may not reach
// loopback addresses, which the URL-import tests rely on.
func newAdminRouter(p Pinger, cols CollectionStore, ops OperationStore, allowPrivate bool) http.Handler {
	return buildRouter(p, cols, ops, &fakeSearch{}, allowPrivate)
}

func newRouterWithSearch(cols CollectionStore, ops OperationStore, search Searcher) http.Handler {
	return buildRouter(fakePinger{}, cols, ops, search, true)
}

func buildRouter(p Pinger, cols CollectionStore, ops OperationStore, search Searcher, allowPrivate bool) http.Handler {
	hash, _ := bcrypt.GenerateFromPassword([]byte(testPassword), bcrypt.MinCost)
	cfg := &config.Config{GinMode: gin.TestMode, CORSOrigins: []string{"http://localhost:5173"}}
	return NewRouter(Deps{
		Config:      cfg,
		Mongo:       p,
		Logger:      slog.New(slog.DiscardHandler),
		Collections: cols,
		Operations:  ops,
		Auth:        auth.New("admin", string(hash), strings.Repeat("k", 32), time.Hour),
		Importer:    importsvc.New(cols, ops, source.New(allowPrivate)),
		Search:      search,
	})
}

// fakeSearch matches operations whose summary, path or operationId contain
// the query, case-insensitively, and records the arguments it was given.
type fakeSearch struct {
	ops            []domain.Operation
	lastCollection string
	lastLimit      int
}

func (f *fakeSearch) Search(_ context.Context, q, collectionID string, limit int) ([]domain.SearchHit, error) {
	f.lastCollection, f.lastLimit = collectionID, limit
	var hits []domain.SearchHit
	for _, op := range f.ops {
		hay := strings.ToLower(op.Summary + " " + op.Path + " " + op.OperationID)
		if strings.Contains(hay, strings.ToLower(q)) {
			op.CollectionID = "c1"
			hits = append(hits, domain.SearchHit{Operation: op, Score: 1})
		}
	}
	return hits, nil
}

func TestHealthzOK(t *testing.T) {
	r := newTestRouter(fakePinger{})
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body %s", rec.Code, rec.Body)
	}
	if !strings.Contains(rec.Body.String(), `"mongo":"ok"`) {
		t.Fatalf("body = %s", rec.Body)
	}
}

func TestHealthzMongoDown(t *testing.T) {
	r := newTestRouter(fakePinger{err: errors.New("down")})
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503; body %s", rec.Code, rec.Body)
	}
	if strings.Contains(rec.Body.String(), "down") {
		t.Fatalf("body leaks internal error: %s", rec.Body)
	}
}
