package importer

import (
	"regexp"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/shiprocket/apidocs/internal/domain"
)

// methodOrder fixes the display order of operations within one path.
var methodOrder = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "TRACE"}

var nonAlnum = regexp.MustCompile(`[^a-z0-9]+`)

// Flatten walks every path and method in the document and returns one
// Operation per HTTP operation, in a stable order. CollectionID and ID are
// left empty for the repository to fill.
func Flatten(doc *openapi3.T) []domain.Operation {
	if doc == nil || doc.Paths == nil {
		return nil
	}
	paths := doc.Paths.Keys()
	sort.Strings(paths)

	var ops []domain.Operation
	order := 0
	for _, path := range paths {
		item := doc.Paths.Value(path)
		if item == nil {
			continue
		}
		byMethod := item.Operations()
		for _, method := range methodOrder {
			op := byMethod[method]
			if op == nil {
				continue
			}
			ops = append(ops, flattenOne(method, path, op, order))
			order++
		}
	}
	return ops
}

func flattenOne(method, path string, op *openapi3.Operation, order int) domain.Operation {
	opID := op.OperationID
	if opID == "" {
		opID = fallbackOperationID(method, path)
	}
	tags := op.Tags
	if tags == nil {
		tags = []string{}
	}
	return domain.Operation{
		Method:      method,
		Path:        path,
		OperationID: opID,
		Summary:     op.Summary,
		Description: op.Description,
		Tags:        tags,
		Deprecated:  op.Deprecated,
		Order:       order,
		SearchText:  searchText(op),
	}
}

// fallbackOperationID derives a stable id such as get_pet_petid when the spec
// omits operationId.
func fallbackOperationID(method, path string) string {
	s := strings.ToLower(method + "_" + path)
	s = nonAlnum.ReplaceAllString(s, "_")
	return strings.Trim(s, "_")
}

// searchText gathers the free text a full-text index should see beyond the
// indexed fields: parameter names and descriptions, response descriptions.
func searchText(op *openapi3.Operation) string {
	var b strings.Builder
	for _, p := range op.Parameters {
		if p == nil || p.Value == nil {
			continue
		}
		b.WriteString(p.Value.Name)
		b.WriteByte(' ')
		b.WriteString(p.Value.Description)
		b.WriteByte(' ')
	}
	if op.RequestBody != nil && op.RequestBody.Value != nil {
		b.WriteString(op.RequestBody.Value.Description)
		b.WriteByte(' ')
	}
	if op.Responses != nil {
		for _, code := range op.Responses.Keys() {
			r := op.Responses.Value(code)
			if r == nil || r.Value == nil || r.Value.Description == nil {
				continue
			}
			b.WriteString(*r.Value.Description)
			b.WriteByte(' ')
		}
	}
	return strings.TrimSpace(b.String())
}
