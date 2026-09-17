// Package importsvc turns an uploaded file or a URL into a stored collection.
// It is the single import path shared by the admin API and the seed CLI.
package importsvc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer"
	"github.com/shiprocket/apidocs/internal/importer/source"
)

// ErrInvalidSlug is returned for slugs that fail domain.ValidSlug.
var ErrInvalidSlug = errors.New("slug must be lowercase letters, digits and hyphens")

// CollectionStore is the write side the service needs.
type CollectionStore interface {
	Upsert(ctx context.Context, c *domain.Collection) (bool, error)
}

// OperationStore replaces a collection's operation index.
type OperationStore interface {
	Replace(ctx context.Context, collectionID string, ops []domain.Operation) error
}

// Service imports specs into the stores.
type Service struct {
	cols  CollectionStore
	ops   OperationStore
	fetch *source.Fetcher
	now   func() time.Time
}

// New wires the service. fetch may be nil when URL imports are not needed.
func New(cols CollectionStore, ops OperationStore, fetch *source.Fetcher) *Service {
	return &Service{cols: cols, ops: ops, fetch: fetch, now: time.Now}
}

// Input describes one import. Exactly one of Data or URL must be set.
type Input struct {
	Data     []byte
	Filename string
	URL      string
	Slug     string
	Name     string
}

// Import parses the input and upserts the collection under in.Slug. It
// reports whether the collection was newly created.
func (s *Service) Import(ctx context.Context, in Input) (*domain.Collection, bool, error) {
	if !domain.ValidSlug(in.Slug) {
		return nil, false, ErrInvalidSlug
	}
	data, filename := in.Data, in.Filename
	if in.URL != "" {
		if s.fetch == nil {
			return nil, false, errors.New("url imports are not enabled")
		}
		res, err := s.fetch.Fetch(ctx, in.URL)
		if err != nil {
			return nil, false, fmt.Errorf("%w: %w", importer.ErrInvalidSpec, err)
		}
		data, filename = res.Data, res.Filename
	}
	if len(data) == 0 {
		return nil, false, fmt.Errorf("%w: empty file", importer.ErrUnknownFormat)
	}

	res, err := importer.Import(ctx, data)
	if err != nil {
		return nil, false, err
	}

	name := in.Name
	if name == "" {
		name = res.Title
	}
	if name == "" {
		name = in.Slug
	}
	col := &domain.Collection{
		Slug:           in.Slug,
		Name:           name,
		Description:    res.Description,
		Version:        res.Version,
		Source:         domain.ImportSource{Type: res.Source, Filename: filename, ImportedAt: s.now().UTC()},
		Servers:        res.Servers,
		OperationCount: len(res.Operations),
		Spec:           res.SpecJSON,
		RawUpload:      data,
	}
	created, err := s.cols.Upsert(ctx, col)
	if err != nil {
		return nil, false, err
	}
	if err := s.ops.Replace(ctx, col.ID, res.Operations); err != nil {
		return nil, false, err
	}
	return col, created, nil
}
