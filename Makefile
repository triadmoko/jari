include .env
env:
	 @echo $(MODE)

dev:
	sh -c 'set -a; . ./.env; set +a; gow run cmd/api/main.go'

# Frontend (SPA, embedded into the Go binary from webui/dist)
WEB_DIST=internal/shared/delivery/http/webui/dist

web-install:
	cd frontend && bun install --frozen-lockfile

web-dev:
	cd frontend && bun run dev

web-build: web-install
	cd frontend && bun run build
	find $(WEB_DIST) -mindepth 1 ! -name .gitkeep -delete
	cp -r frontend/build/. $(WEB_DIST)/

build: web-build
	CGO_ENABLED=0 go build -o main ./cmd/api

# Migrations: SQL files in ./migrations are embedded into the binary and applied
# by its `migrate` subcommand (same DB config as the server, loaded from .env).
MIGRATE=sh -c 'set -a; . ./.env; set +a; exec go run ./cmd/api migrate "$$@"' migrate

migration-up:
	$(MIGRATE) up

# n=N rolls back N migrations (default 1), n=all rolls back everything
migration-down:
	$(MIGRATE) down $(n)

migration-create:
	migrate create -ext sql -dir ./migrations -seq $(name)

migration-force:
	$(MIGRATE) force $(version)

migration-version:
	$(MIGRATE) version

swag:
	swag init --parseInternal -g cmd/api/main.go --output ./docs

# Testing
test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-coverage-report:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Mock generation
mock-gen:
	mockery

mock-clean:
	rm -rf internal/mocks
