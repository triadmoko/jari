-include .env
env:
	 @echo $(MODE)

dev:
	sh -c 'set -a; . ./.env; set +a; gow run ./cmd/api'

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

# Regenerate the frontend API client (frontend/src/lib/api/api.ts) from docs/swagger.json.
# Run after `make swag`; the generator version is pinned in frontend/package.json.
api-gen:
	cd frontend && bun run api-gen

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

# Dependency license reports (docs/licenses/) and policy check.
# The repo is Apache-2.0-licensed (see LICENSE): copyleft licenses (GPL, AGPL, LGPL, SSPL, ...) are not allowed.
# Tool versions are pinned so the reports are reproducible.
GO_LICENSES=go run github.com/google/go-licenses@v1.6.0
WEB_LICENSES=bunx license-checker-rseidelsohn@4.4.2 --excludePrivatePackages
# Licenses accepted in frontend dependencies (update docs/licenses/README.md when adding one)
WEB_ALLOWED_LICENSES=MIT;MIT*;ISC;0BSD;BSD-2-Clause;BSD-3-Clause;Apache-2.0;MPL-2.0;OFL-1.1;Python-2.0;Unlicense;BlueOak-1.0.0;MIT OR Apache-2.0;(MIT OR CC0-1.0);(MPL-2.0 OR Apache-2.0)

licenses:
	$(GO_LICENSES) report ./cmd/api --ignore app > docs/licenses/go.csv
	cd frontend && $(WEB_LICENSES) --csv > ../docs/licenses/frontend.csv

licenses-check:
	$(GO_LICENSES) check ./cmd/api --ignore app --disallowed_types=forbidden,restricted
	cd frontend && $(WEB_LICENSES) --onlyAllow '$(WEB_ALLOWED_LICENSES)' --summary
