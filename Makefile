# Makefile for MyTodo - Task Management System

.PHONY: help dev test build clean migrate-up migrate-down migrate-create seed docker-up docker-down lint format

# Default target
.DEFAULT_GOAL := help

# Variables
APP_NAME=mytodo
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")
MIGRATION_DIR=apps/api/pkg/database/migrations
MAIN_FILE=apps/api/cmd/server/main.go

## help: Display this help message
help:
	@echo "Available commands:"
	@echo ""
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /' | column -t -s ':'

## dev: Start the development server with hot reload
dev:
	@echo "Starting development server..."
	air -c .air.toml || go run $(MAIN_FILE)

## test: Run all tests
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

## test-unit: Run unit tests only
test-unit:
	@echo "Running unit tests..."
	go test -v -short ./...

## test-integration: Run integration tests only
test-integration:
	@echo "Running integration tests..."
	go test -v -run Integration ./...

## test-e2e: Run end-to-end tests
test-e2e:
	@echo "Running E2E tests..."
	go test -v ./apps/api/tests/e2e/...

## build: Build the application for production
build:
	@echo "Building application..."
	go build -o bin/$(APP_NAME) $(MAIN_FILE)

## build-docker: Build Docker image
build-docker:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):latest -f infrastructure/docker/api/Dockerfile .

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean

## run: Run the application
run: build
	@echo "Running application..."
	./bin/$(APP_NAME)

## migrate-up: Run database migrations up
migrate-up:
	@echo "Running migrations up..."
	go run tools/migrations/main.go up

## migrate-down: Rollback last migration
migrate-down:
	@echo "Rolling back migration..."
	go run tools/migrations/main.go down

## migrate-create: Create a new migration (usage: make migrate-create name=add_users_table)
migrate-create:
	@echo "Creating migration: $(name)"
	go run tools/migrations/main.go create $(name)

## migrate-status: Check migration status
migrate-status:
	@echo "Checking migration status..."
	go run tools/migrations/main.go status

## seed: Seed the database with test data
seed:
	@echo "Seeding database..."
	go run tools/seed/main.go

## docker-up: Start all Docker services
docker-up:
	@echo "Starting Docker services..."
	docker-compose up -d

## docker-down: Stop all Docker services
docker-down:
	@echo "Stopping Docker services..."
	docker-compose down

## docker-logs: View Docker logs
docker-logs:
	docker-compose logs -f

## db-reset: Drop and recreate database (WARNING: destructive)
db-reset:
	@echo "Resetting database..."
	./infrastructure/scripts/init.sh

## lint: Run linters
lint:
	@echo "Running linters..."
	golangci-lint run ./...

## format: Format code
format:
	@echo "Formatting code..."
	gofmt -s -w $(GO_FILES)
	goimports -w $(GO_FILES)

## tidy: Tidy Go modules
tidy:
	@echo "Tidying Go modules..."
	go mod tidy

## install-tools: Install development tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/swaggo/swag/cmd/swag@latest

## deps: Install/update dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod verify

## swagger: Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	swag init -g $(MAIN_FILE) -o docs/swagger

## health: Check if API is healthy
health:
	@echo "Checking API health..."
	curl -f http://localhost:8080/health || echo "API is not responding"

## setup: Initial project setup
setup: deps install-tools docker-up migrate-up seed
	@echo "Project setup complete!"
	@echo "Run 'make dev' to start development server"

## deploy: Deploy to production (placeholder)
deploy:
	@echo "Deploying to production..."
	./infrastructure/scripts/deploy.sh

## backup: Backup database
backup:
	@echo "Backing up database..."
	./infrastructure/scripts/backup.sh

## restore: Restore database from backup
restore:
	@echo "Restoring database..."
	./infrastructure/scripts/restore.sh

## watch: Watch for changes and run tests
watch:
	@echo "Watching for changes..."
	go test -v ./... -watch

## coverage: Generate test coverage report
coverage:
	@echo "Generating coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## benchmark: Run benchmarks
benchmark:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

## profile: Profile the application
profile:
	@echo "Profiling application..."
	go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=. ./...
	go tool pprof cpu.prof

## generate: Run go generate
generate:
	@echo "Running go generate..."
	go generate ./...

## check: Run all checks (format, lint, test)
check: format lint test
	@echo "All checks passed!"

## ci: Run CI pipeline locally
ci: tidy format lint test build
	@echo "CI pipeline completed!"
