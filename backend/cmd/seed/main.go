// Command seed imports a spec file or URL straight into MongoDB. It shares
// the import path with the admin API and exists for scripted bootstrapping.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer/source"
	"github.com/shiprocket/apidocs/internal/importsvc"
	"github.com/shiprocket/apidocs/internal/repo/mongo"
)

func main() {
	file := flag.String("file", "", "path to an OpenAPI 2.0/3.x or Postman file (json or yaml)")
	src := flag.String("url", "", "URL of a spec, or of a Postman published-docs page (used instead of -file)")
	slug := flag.String("slug", "", "url slug for the collection, e.g. petstore")
	name := flag.String("name", "", "display name (defaults to the spec title)")
	uri := flag.String("mongo-uri", envOr("MONGO_URI", "mongodb://localhost:27017"), "mongo connection string")
	dbName := flag.String("mongo-db", envOr("MONGO_DB", "apidocs"), "mongo database name")
	flag.Parse()

	if (*file == "" && *src == "") || *slug == "" {
		fmt.Fprintln(os.Stderr, "usage: seed (-file spec.yaml | -url https://docs.example.com/) -slug my-api [-name \"My API\"]")
		os.Exit(2)
	}
	if !domain.ValidSlug(*slug) {
		fmt.Fprintln(os.Stderr, "slug must be lowercase letters, digits and hyphens")
		os.Exit(2)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	in := importsvc.Input{Slug: *slug, Name: *name, URL: *src}
	if *file != "" {
		data, err := os.ReadFile(*file)
		if err != nil {
			fail("read file", err)
		}
		in.Data, in.Filename = data, filepath.Base(*file)
	}

	db, err := mongo.Connect(ctx, *uri, *dbName)
	if err != nil {
		fail("mongo", err)
	}
	defer func() { _ = db.Close(context.Background()) }()

	cols := mongo.NewCollectionRepo(db.DB)
	ops := mongo.NewOperationRepo(db.DB)
	if err := cols.EnsureIndexes(ctx); err != nil {
		fail("indexes", err)
	}
	if err := ops.EnsureIndexes(ctx); err != nil {
		fail("indexes", err)
	}

	// The CLI runs on an operator's machine, so private URLs are allowed here.
	col, created, err := importsvc.New(cols, ops, source.New(true)).Import(ctx, in)
	if err != nil {
		fail("import", err)
	}
	action := "updated"
	if created {
		action = "created"
	}
	fmt.Printf("%s collection %q (%s): %d operations from %s source\n", action, col.Slug, col.ID, col.OperationCount, col.Source.Type)
}

func fail(step string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", step, err)
	os.Exit(1)
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
