GO = go
GOFMT = gofmt
GOLINT = golint
GOCOVER = go test -coverprofile=coverage.out
GO_BIN = $(GO) build -o bin/app

PKG = github.com/sk-pathak/go-structure
BUILD_DIR = bin
SRC_DIR = cmd
APP_NAME = go-structure

.PHONY: all
all: help

.PHONY: help
help:
	@echo "Makefile for Go project"
	@echo ""
	@echo "Available targets:"
	@echo "  build         Build the Go application"
	@echo "  run           Run the Go application"
	@echo "  test          Run the tests"
	@echo "  lint          Run Go linter"
	@echo "  format        Run Go formatter"
	@echo "  migrate       Run the database migrations"
	@echo "  clean         Clean up the build artifacts"
	@echo "  docker-build  Build the Docker image"
	@echo "  docker-run    Run the Docker container"

.PHONY: build
build:
	$(GO) build -o $(BUILD_DIR)/app/${APP_NAME} $(SRC_DIR)/server/main.go

.PHONY: exec
exec:
	make build && $(BUILD_DIR)/app/${APP_NAME}

.PHONY: run
run:
	$(GO) run $(SRC_DIR)/server/main.go

.PHONY: test
test:
	$(GO) test ./...

.PHONY: lint
lint:
	$(GOLINT) ./...

.PHONY: format
format:
	$(GOFMT) -s -w $(SRC_DIR)

.PHONY: create-migration
create-migration:
	./scripts/create_migration.sh $(name)

.PHONY: apply-migration
apply-migration:
	./scripts/apply_migration.sh

.PHONY: sqlc
sqlc:
	./scripts/sqlc.sh

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)/app

.PHONY: docker
docker:
	docker compose up
