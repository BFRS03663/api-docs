// Package ui embeds the production frontend bundle. `make build` and the
// Dockerfile copy frontend/dist into ./dist before compiling; a development
// build without the bundle still compiles and simply reports the UI as not
// built.
package ui

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var bundle embed.FS

// FS returns the bundle rooted at the dist directory.
func FS() fs.FS {
	sub, err := fs.Sub(bundle, "dist")
	if err != nil {
		return bundle
	}
	return sub
}
