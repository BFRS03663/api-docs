# API Docs Portal

Self-hosted API documentation site in the style of Postman's published docs. Upload an OpenAPI/Swagger spec or a Postman collection, browse it with a Try-it runner, and let LLMs and crawlers read the same content as Markdown via `llms.txt`.

## Stack

- **Backend:** Go 1.26, Gin, MongoDB driver v2
- **Frontend:** React 19, Vite, TypeScript, Tailwind, Scalar API reference
- **Database:** MongoDB 7 (Docker)

## Quick start

```bash
cp .env.example .env
make hash-password          # paste the printed line into .env
openssl rand -hex 32        # paste into JWT_SECRET='' in .env
make docker-up
```

- Frontend: http://localhost:5173
- Backend health: http://localhost:8080/healthz

## Development

| Command | What it does |
|---|---|
| `make docker-up` | Mongo, backend with hot reload, Vite dev server |
| `make test` | Go tests with race detector, Vitest |
| `make test-backend-db` | Go tests including Mongo repository tests (stack must be up) |
| `make seed FILE=... SLUG=...` | Import a spec file directly into Mongo (until the admin UI lands) |
| `make seed-url URL=... SLUG=...` | Import from a URL: a raw OpenAPI/Postman file, or a Postman published-docs page (the collection JSON is discovered from the page) |
| `make lint` | go vet, golangci-lint, tsc |
| `make build` | Production frontend bundle and backend binary |

## Try it (request runner)

The "Test Request" button in the reference sends requests through `ANY /api/v1/proxy?scalar_url=<target>`, so APIs that reject cross-origin browser calls still work. The proxy forwards the method, query, headers and body, relays the upstream status, headers and body, and:

- refuses private, loopback and link-local destinations on every redirect hop (unless `PROXY_ALLOW_PRIVATE=true`),
- restricts destinations to `PROXY_ALLOWED_HOSTS` when that list is set,
- never forwards cookies or hop-by-hop headers and never relays `Set-Cookie`,
- caps responses at 5 MB (`X-Proxy-Truncated: true` when cut), request bodies at 10 MB, and the exchange at 15 seconds,
- allows 60 requests per minute per client IP.

## LLM-friendly output and search

Every collection is also available as text so language models, crawlers and scripts can read it:

```
GET /llms.txt                          index (llmstxt.org convention): per collection its base URLs, auth scheme,
                                       intro, guide chapters and every endpoint grouped by section, each linked
                                       to its own Markdown page
GET /llms-full.txt                     the index followed by every collection's complete documentation
GET /docs/<slug>.md  |  .txt           one collection as Markdown or plain text
GET /docs/<slug>/<operationId>.md      one endpoint
GET /docs/<slug>/openapi.json          the canonical OpenAPI document
GET /api/v1/collections/<slug>         with `Accept: text/markdown` returns the Markdown rendering
GET /api/v1/search?q=<words>[&collection=<slug>][&limit=25]   full-text search over endpoints
```

The docs page has an "Export for LLMs" menu with the same links plus "Copy as Markdown". `SITE_NAME` in `.env` sets the title used in `llms.txt`.

## Admin

Open `/admin` and sign in with `ADMIN_USERNAME` and the password behind `ADMIN_PASSWORD_HASH`. From there you can import a collection by uploading an OpenAPI or Postman file or by pasting a URL (Postman published-docs pages are detected automatically), re-import an existing one, or delete it. The **Write spec** tab is a Swagger-Editor-style page: the OpenAPI document (YAML or JSON) on the left, the rendered reference on the right, re-rendered as you type; **Publish** creates the collection or replaces an existing one. **Edit** on the collections table opens a stored document in the same editor (Postman imports open as their converted OpenAPI document). Sessions are JWTs signed with `JWT_SECRET` and last 12 hours; login is limited to 5 attempts per minute per IP.

The same operations are available over HTTP with a bearer token:

```
POST   /api/v1/auth/login                {"username","password"} -> {"token","expiresAt"}
POST   /api/v1/admin/collections         multipart (file, slug, name) or JSON {"url","slug","name"}
PUT    /api/v1/admin/collections/:slug   same body, re-imports
GET    /api/v1/admin/collections/:slug/source   {"slug","filename","type","canonical","content"} — the editable text
DELETE /api/v1/admin/collections/:slug
```

Server-side fetches refuse private and loopback addresses unless `PROXY_ALLOW_PRIVATE=true`.

## Branding

Set in `.env` (read by docker compose and Vite): `VITE_SITE_NAME`, `VITE_SITE_LOGO` (a URL or a file under `frontend/public`), `VITE_ACCENT`. The Scalar reference is themed from the same accent and renders examples in a dark right-hand column.

## Deploy

`make build` produces a single static binary at `backend/bin/server` with the frontend embedded; it serves the UI, the API and the text surface on one port. The Dockerfile does the same in a distroless image:

```bash
docker build -t apidocs --build-arg VITE_SITE_NAME="Shiprocket API" --build-arg VITE_SITE_LOGO=/logo.jpg .
docker run --env-file .env -e MONGO_URI=mongodb://mongo:27017 -p 8080:8080 apidocs
```

or, with the bundled Mongo, `make docker-prod` (compose profile `prod`, port `APP_PORT`, default 8090). Branding env vars are baked into the bundle at build time. Behind a reverse proxy, forward `X-Forwarded-Proto` and `X-Forwarded-Host` so text renderings carry the public URL.

## Layout

```
backend/   Go service (cmd/server, internal/{config,httpapi,repo,importer,export,auth,proxy})
frontend/  React app
```

## Roadmap

Milestones M0 to M7 are tracked in the project plan: scaffold, OpenAPI import, Postman import, Scalar reader, admin auth and upload, search plus LLM text export, Try-it proxy, production image.
