// Package importer turns uploaded API descriptions (OpenAPI 2.0/3.x, Postman)
// into a canonical OpenAPI 3.x document plus flattened operations.
package importer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oasdiff/yaml"

	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer/postman"
)

// Format is the detected upload format.
type Format string

// Detectable formats.
const (
	FormatUnknown  Format = ""
	FormatOpenAPI2 Format = "openapi2"
	FormatOpenAPI3 Format = "openapi3"
	FormatPostman  Format = "postman"
)

// Sentinel errors callers map to HTTP statuses.
var (
	ErrUnknownFormat     = errors.New("unrecognised file format")
	ErrUnsupportedFormat = errors.New("format not supported yet")
	ErrInvalidSpec       = errors.New("invalid specification")
)

// Result is the outcome of a successful import.
type Result struct {
	Doc         *openapi3.T
	SpecJSON    []byte
	Source      domain.SourceType
	Title       string
	Description string
	Version     string
	Servers     []string
	Operations  []domain.Operation
}

// Detect sniffs the upload to decide which parser to use. It accepts JSON or
// YAML and never returns FormatUnknown without an error.
func Detect(data []byte) (Format, error) {
	if len(strings.TrimSpace(string(data))) == 0 {
		return FormatUnknown, fmt.Errorf("%w: empty file", ErrUnknownFormat)
	}
	js, err := yaml.YAMLToJSON(data)
	if err != nil {
		return FormatUnknown, fmt.Errorf("%w: not valid JSON or YAML", ErrUnknownFormat)
	}
	var probe struct {
		OpenAPI string `json:"openapi"`
		Swagger string `json:"swagger"`
		Info    struct {
			Schema    string `json:"schema"`
			PostmanID string `json:"_postman_id"`
		} `json:"info"`
		Item     json.RawMessage `json:"item"`
		Requests json.RawMessage `json:"requests"` // Postman v1 marker
	}
	if err := json.Unmarshal(js, &probe); err != nil {
		return FormatUnknown, fmt.Errorf("%w: top level is not an object", ErrUnknownFormat)
	}
	switch {
	case strings.HasPrefix(probe.OpenAPI, "3."):
		return FormatOpenAPI3, nil
	case probe.Swagger == "2.0":
		return FormatOpenAPI2, nil
	case probe.Info.PostmanID != "" || strings.Contains(probe.Info.Schema, "getpostman.com/json/collection") || len(probe.Item) > 0 || len(probe.Requests) > 0:
		return FormatPostman, nil
	}
	return FormatUnknown, fmt.Errorf("%w: no openapi, swagger, or postman markers", ErrUnknownFormat)
}

// Import detects the format, parses it, and returns the canonical document
// with flattened operations.
func Import(ctx context.Context, data []byte) (*Result, error) {
	format, err := Detect(data)
	if err != nil {
		return nil, err
	}

	var doc *openapi3.T
	var source domain.SourceType
	switch format {
	case FormatOpenAPI2, FormatOpenAPI3:
		doc, source, err = parseOpenAPI(ctx, data, format)
	case FormatPostman:
		var js []byte
		js, err = postman.ToOpenAPIJSON(data)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrInvalidSpec, err)
		}
		doc, err = loadAndValidate(ctx, js)
		source = domain.SourcePostman
	default:
		return nil, ErrUnknownFormat
	}
	if err != nil {
		return nil, err
	}

	specJSON, err := json.Marshal(doc)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidSpec, err)
	}

	res := &Result{
		Doc:        doc,
		SpecJSON:   specJSON,
		Source:     source,
		Operations: Flatten(doc),
	}
	if doc.Info != nil {
		res.Title = doc.Info.Title
		res.Description = doc.Info.Description
		res.Version = doc.Info.Version
	}
	for _, s := range doc.Servers {
		if s != nil && s.URL != "" {
			res.Servers = append(res.Servers, s.URL)
		}
	}
	return res, nil
}
