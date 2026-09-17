// Package domain holds the storage-agnostic types shared by importers,
// repositories, and HTTP handlers.
package domain

import (
	"errors"
	"regexp"
	"time"
)

// ErrNotFound is returned by repositories when a record does not exist.
var ErrNotFound = errors.New("not found")

// SourceType identifies which format a collection was imported from.
type SourceType string

// Supported import sources.
const (
	SourceOpenAPI2 SourceType = "openapi2"
	SourceOpenAPI3 SourceType = "openapi3"
	SourcePostman  SourceType = "postman"
)

// ImportSource records where a collection's spec came from.
type ImportSource struct {
	Type       SourceType `json:"type"`
	Filename   string     `json:"filename"`
	ImportedAt time.Time  `json:"importedAt"`
}

// Collection is one documented API. Spec holds the canonical OpenAPI 3.x JSON;
// RawUpload holds the original bytes the admin uploaded.
type Collection struct {
	ID             string       `json:"id"`
	Slug           string       `json:"slug"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Version        string       `json:"version"`
	Source         ImportSource `json:"source"`
	Servers        []string     `json:"servers"`
	OperationCount int          `json:"operationCount"`
	Spec           []byte       `json:"-"`
	RawUpload      []byte       `json:"-"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
}

// Operation is a flattened view of one HTTP operation, derived from the spec
// on import. It backs listing, search, and Markdown export.
type Operation struct {
	ID           string   `json:"id"`
	CollectionID string   `json:"collectionId"`
	Method       string   `json:"method"`
	Path         string   `json:"path"`
	OperationID  string   `json:"operationId"`
	Summary      string   `json:"summary"`
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	Deprecated   bool     `json:"deprecated"`
	Order        int      `json:"order"`
	SearchText   string   `json:"-"`
}

// SearchHit is one full-text search result with its relevance score.
type SearchHit struct {
	Operation
	CollectionSlug string  `json:"collectionSlug"`
	CollectionName string  `json:"collectionName"`
	Score          float64 `json:"score"`
}

var slugPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

// ValidSlug reports whether s is a URL-safe collection slug.
func ValidSlug(s string) bool {
	return len(s) <= 64 && slugPattern.MatchString(s)
}
