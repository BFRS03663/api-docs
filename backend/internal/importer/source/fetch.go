// Package source retrieves API descriptions from remote URLs, including
// Postman published-documentation pages, which expose their collection JSON
// through a predictable endpoint advertised in the page's meta tags.
package source

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/shiprocket/apidocs/internal/netguard"
)

// MaxBytes caps how much we read from a remote source.
const MaxBytes = 25 << 20

// ErrUnsupported is returned when a page is HTML but not a recognised
// documentation host.
var ErrUnsupported = errors.New("url is an HTML page without a recognised collection or spec link")

var metaRe = regexp.MustCompile(`<meta\s+name="([^"]+)"\s+content="([^"]*)"`)

// Result is a downloaded description plus a filename hint for the record.
type Result struct {
	Data     []byte
	Filename string
	Source   string // the URL the bytes actually came from
}

// Fetcher downloads specs. Client may be customised for timeouts or proxies.
type Fetcher struct {
	Client *http.Client
}

// New returns a Fetcher whose connections are checked against the private
// address guard. allowPrivate is for local development and the seed CLI.
func New(allowPrivate bool) *Fetcher {
	return &Fetcher{Client: netguard.NewHTTPClient(60*time.Second, allowPrivate)}
}

// Fetch downloads rawURL. If it is JSON or YAML it is returned as-is. If it is
// a Postman published-docs page, the underlying collection JSON is fetched.
func (f *Fetcher) Fetch(ctx context.Context, rawURL string) (*Result, error) {
	u, err := url.Parse(rawURL)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return nil, fmt.Errorf("invalid url %q", rawURL)
	}
	body, ctype, err := f.get(ctx, u.String(), "application/json, application/yaml, text/yaml, text/html;q=0.8, */*;q=0.5")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(string(body[:min(len(body), 512)]))
	isHTML := strings.Contains(ctype, "text/html") || strings.HasPrefix(strings.ToLower(trimmed), "<!doctype html") || strings.HasPrefix(trimmed, "<html")
	if !isHTML {
		return &Result{Data: body, Filename: filenameFrom(u, ".json"), Source: u.String()}, nil
	}

	collURL, publishedID, ok := postmanCollectionURL(u, string(body))
	if !ok {
		return nil, ErrUnsupported
	}
	data, _, err := f.get(ctx, collURL, "application/json")
	if err != nil {
		return nil, fmt.Errorf("postman collection: %w", err)
	}
	return &Result{Data: data, Filename: publishedID + ".postman_collection.json", Source: collURL}, nil
}

func (f *Fetcher) get(ctx context.Context, rawURL, accept string) ([]byte, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Accept", accept)
	req.Header.Set("User-Agent", "apidocs-importer/1.0")
	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("fetch %s: %w", rawURL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("fetch %s: status %d", rawURL, resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, MaxBytes+1))
	if err != nil {
		return nil, "", fmt.Errorf("read %s: %w", rawURL, err)
	}
	if len(body) > MaxBytes {
		return nil, "", fmt.Errorf("fetch %s: response exceeds %d bytes", rawURL, MaxBytes)
	}
	return body, resp.Header.Get("Content-Type"), nil
}

// postmanCollectionURL builds the collection endpoint from the meta tags a
// Postman published-docs page carries (ownerId, publishedId, environmentUID).
// The endpoint lives on the same host as the page, so custom domains work.
func postmanCollectionURL(page *url.URL, html string) (collURL, publishedID string, ok bool) {
	meta := map[string]string{}
	for _, m := range metaRe.FindAllStringSubmatch(html, -1) {
		meta[m[1]] = m[2]
	}
	owner, pub := meta["ownerId"], meta["publishedId"]
	if owner == "" || pub == "" {
		return "", "", false
	}
	q := url.Values{"segregateAuth": {"true"}, "versionTag": {"latest"}}
	if env := meta["environmentUID"]; env != "" {
		q.Set("environment", env)
	}
	if tag := meta["versionTagId"]; tag != "" {
		q.Set("versionTag", tag)
	}
	coll := url.URL{Scheme: page.Scheme, Host: page.Host, Path: fmt.Sprintf("/api/collections/%s/%s", owner, pub), RawQuery: q.Encode()}
	return coll.String(), pub, true
}

func filenameFrom(u *url.URL, def string) string {
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	name := parts[len(parts)-1]
	if name == "" || !strings.Contains(name, ".") {
		return "spec" + def
	}
	return name
}
