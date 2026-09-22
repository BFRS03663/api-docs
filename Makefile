.PHONY: seed seed-url docker-prod help docker-up docker-down docker-logs dev-backend dev-frontend test test-backend test-frontend lint lint-backend lint-frontend build hash-password llms-full

help: ## Show targets
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-16s %s\n", $$1, $$2}'

llms-full: ## Rebuild llms-full.txt from llms.txt and docs/shiprocket-api/*.md
	node scripts/build-llms-full.mjs

docker-up: ## Start mongo + backend (air) + frontend (vite)
	docker compose up --build -d
	@echo "backend  http://localhost:$${BACKEND_PORT:-8080}/healthz"
	@echo "frontend http://localhost:$${FRONTEND_PORT:-5173}"

docker-down: ## Stop the stack
	docker compose down

docker-logs: ## Tail all logs
	docker compose logs -f

dev-backend: ## Run backend locally (needs .env exported and mongo reachable)
	cd backend && set -a && . ../.env && set +a && MONGO_URI=mongodb://localhost:27017 go run ./cmd/server

dev-frontend: ## Run frontend locally
	cd frontend && npm install && npm run dev

test: test-backend test-frontend ## Run all tests

test-backend:
	cd backend && go test ./... -race -cover

test-frontend:
	cd frontend && npm run test

lint: lint-backend lint-frontend ## Run all linters

lint-backend:
	cd backend && go vet ./... && (command -v golangci-lint >/dev/null && golangci-lint run ./... || echo "golangci-lint not installed, skipped")

lint-frontend:
	cd frontend && npm run typecheck

build: ## Build frontend, embed it, and build the backend binary
	cd frontend && npm run build
	rm -rf backend/internal/ui/dist && mkdir -p backend/internal/ui/dist && cp -r frontend/dist/. backend/internal/ui/dist/ && touch backend/internal/ui/dist/.gitkeep
	cd backend && CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o bin/server ./cmd/server

docker-prod: ## Build and run the production container (profile prod) on APP_PORT (default 8090)
	docker compose --profile prod up -d --build app
	@echo "app http://localhost:$${APP_PORT:-8090}"

seed: ## Import a spec into the local Mongo: make seed FILE=backend/testdata/petstore-v3.yaml SLUG=petstore
	cd backend && MONGO_URI=$${MONGO_URI:-mongodb://localhost:27017} go run ./cmd/seed -file ../$(FILE) -slug $(SLUG)

seed-url: ## Import from a URL: make seed-url URL=https://apidocs.example.com/ SLUG=example
	cd backend && MONGO_URI=$${MONGO_URI:-mongodb://localhost:27017} go run ./cmd/seed -url "$(URL)" -slug $(SLUG)

test-backend-db: ## Backend tests including Mongo repo tests (needs docker-up)
	cd backend && TEST_MONGO_URI=mongodb://localhost:27017 go test ./... -race -cover -count=1

hash-password: ## Print an ADMIN_PASSWORD_HASH line for .env (prompts for password)
	@cd backend && go run ./cmd/hashpw
