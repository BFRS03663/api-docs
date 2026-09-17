package importer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oasdiff/yaml"

	"github.com/shiprocket/apidocs/internal/domain"
)

// parseOpenAPI loads a Swagger 2.0 or OpenAPI 3.x document, upgrading 2.0 to
// 3.0, resolving local $refs, and validating structure. External refs are
// never followed: uploads must be self-contained.
func parseOpenAPI(ctx context.Context, data []byte, format Format) (*openapi3.T, domain.SourceType, error) {
	js, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, "", fmt.Errorf("%w: %v", ErrInvalidSpec, err)
	}

	source := domain.SourceOpenAPI3
	if format == FormatOpenAPI2 {
		source = domain.SourceOpenAPI2
		var v2 openapi2.T
		if err := json.Unmarshal(js, &v2); err != nil {
			return nil, "", fmt.Errorf("%w: swagger 2.0 parse: %v", ErrInvalidSpec, err)
		}
		v3, err := openapi2conv.ToV3(&v2)
		if err != nil {
			return nil, "", fmt.Errorf("%w: swagger 2.0 upgrade: %v", ErrInvalidSpec, err)
		}
		// Round-trip through JSON so the loader resolves the converted $refs.
		if js, err = json.Marshal(v3); err != nil {
			return nil, "", fmt.Errorf("%w: %v", ErrInvalidSpec, err)
		}
	}

	doc, err := loadAndValidate(ctx, js)
	if err != nil {
		return nil, "", err
	}
	return doc, source, nil
}

// loadAndValidate parses OpenAPI 3.x JSON, resolves local $refs and applies
// structural validation. It is shared by every import path.
func loadAndValidate(ctx context.Context, js []byte) (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	loader.Context = ctx
	loader.IsExternalRefsAllowed = false
	doc, err := loader.LoadFromData(js)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidSpec, err)
	}

	// Structural validation only. Example, default, format and pattern checks
	// reject too many real-world specs that still render fine.
	err = doc.Validate(ctx,
		openapi3.DisableExamplesValidation(),
		openapi3.DisableSchemaDefaultsValidation(),
		openapi3.DisableSchemaFormatValidation(),
		openapi3.DisableSchemaPatternValidation(),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidSpec, err)
	}
	if doc.Paths == nil || doc.Paths.Len() == 0 {
		return nil, fmt.Errorf("%w: no paths defined", ErrInvalidSpec)
	}
	return doc, nil
}
