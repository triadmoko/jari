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

url=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

migration-up:
	migrate -database "$(url)" -path ./migration/ up $(version)
	
migration-down:
	migrate -database "$(url)" -path ./migration/ down $(version)
	
migration-create:
	migrate create -ext sql -dir ./migration/ -seq $(name)

migration-force:
	migrate -database "$(url)" -path ./migration/ force $(version)

migration-version:
	migrate -database "$(url)" -path ./migration/ version
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
