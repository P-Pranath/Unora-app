.PHONY: build run test clean migration-create migration-up migration-down migration-status help docker ent-generate ent-init lint fmt bootstrap swagger types generate-all

# App info
APP_NAME=github.com/UnoraApp/be

# Build and run settings
GO=$(shell which go)
GOFLAGS=-v
BUILD_DIR=./build
ENV_FILE=./.env

# Database configuration (loaded from .env file)
DB_HOST := 127.0.0.1
DB_PORT := 3306
DB_USER ?= $(shell grep DB_USER $(ENV_FILE) | cut -d '=' -f2 | tr -d '\r')
DB_PASS ?= $(shell grep DB_PASSWORD $(ENV_FILE) | cut -d '=' -f2 | tr -d '\r')
DB_NAME ?= $(shell grep DB_NAME $(ENV_FILE) | cut -d '=' -f2 | tr -d '\r')

# Migration settings
GOOSE=goose
MIGRATION_DIR=./migrations/versions

# Set default command
.DEFAULT_GOAL := help

DB_CONTAINER_NAME=mysql-local
DB_BACKUP_DIR=./backups
DB_BACKUP_FILE=$(DB_BACKUP_DIR)/dump_$(shell date +"%Y%m%d_%H%M%S").sql

help: ## Display available commands
	@echo "\033[1;34m=== Available commands ===\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2}'

# ========================================
# Development
# ========================================

run: ## Run the server locally (without Docker)
	@echo "\033[1;34m=== Running server ===\033[0m"
	@$(GO) run ./cmd/server

run-worker: ## Run the worker locally (without Docker)
	@echo "\033[1;34m=== Running worker ===\033[0m"
	@$(GO) run ./cmd/worker

run-cron: ## Run the cron locally (without Docker)
	@echo "\033[1;34m=== Running cron ===\033[0m"
	@$(GO) run ./cmd/cron

build: ## Build all binaries
	@echo "\033[1;34m=== Building binaries ===\033[0m"
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/server ./cmd/server
	@$(GO) build -o $(BUILD_DIR)/worker ./cmd/worker
	@$(GO) build -o $(BUILD_DIR)/cron ./cmd/cron
	@echo "\033[1;32m=== Build completed ===\033[0m"

test: ## Run Go unit tests
	@echo "\033[1;34m=== Running tests ===\033[0m"
	@$(GO) test -v ./... | sed 's/^/  /'
	@echo "\033[1;32m=== Tests completed ===\033[0m"

clean: ## Clean up build artifacts and Go cache
	@echo "\033[1;34m=== Cleaning build files ===\033[0m"
	@rm -rf $(BUILD_DIR)
	@rm -rf tmp/
	@$(GO) clean
	@echo "\033[1;32m=== Clean completed ===\033[0m"

# ========================================
# Docker
# ========================================

docker: ## Run docker script to choose env and action (up, down, build, logs)
	@echo "\033[1;34m=== Running Docker script ===\033[0m"
	@bash scripts/docker.sh

docker-up: ## Start all containers (local environment)
	@echo "\033[1;34m=== Starting containers ===\033[0m"
	@docker compose -f docker/docker-compose.local.yml up -d
	@echo "\033[1;32m=== Containers started ===\033[0m"

docker-down: ## Stop all containers (local environment)
	@echo "\033[1;34m=== Stopping containers ===\033[0m"
	@docker compose -f docker/docker-compose.local.yml down
	@echo "\033[1;32m=== Containers stopped ===\033[0m"

docker-build: ## Build all Docker images (local environment)
	@echo "\033[1;34m=== Building Docker images ===\033[0m"
	@docker compose -f docker/docker-compose.local.yml build
	@echo "\033[1;32m=== Build completed ===\033[0m"

docker-logs: ## Show logs for all containers
	@docker compose -f docker/docker-compose.local.yml logs -f

docker-infra: ## Start only infrastructure (MySQL, Redis, MinIO)
	@echo "\033[1;34m=== Starting infrastructure ===\033[0m"
	@docker compose -f docker/docker-compose.local.yml up -d mysql redis minio
	@echo "\033[1;32m=== Infrastructure started ===\033[0m"

# ========================================
# Database
# ========================================

backup-db: ## Backup MySQL database to ./backups folder
	@echo "\033[1;34m=== Backing up MySQL database ===\033[0m"
	@mkdir -p $(DB_BACKUP_DIR)
	@docker exec $(DB_CONTAINER_NAME) sh -c 'exec mysqldump -u$(DB_USER) -p$$MYSQL_ROOT_PASSWORD $(DB_NAME)' > $(DB_BACKUP_FILE)
	@echo "\033[1;32m=== Backup saved to $(DB_BACKUP_FILE) ===\033[0m"

restore-db: ## Restore latest MySQL backup
	@echo "\033[1;34m=== Restoring MySQL database from latest backup ===\033[0m"
	@BACKUP_FILE=$$(ls -t $(DB_BACKUP_DIR)/*.sql | head -n 1); \
	if [ -z "$$BACKUP_FILE" ]; then \
		echo "\033[1;31mNo backup file found in $(DB_BACKUP_DIR)\033[0m"; \
		exit 1; \
	fi; \
	docker exec -i $(DB_CONTAINER_NAME) sh -c 'exec mysql -u$(DB_USER) -p$$MYSQL_ROOT_PASSWORD $(DB_NAME)' < $$BACKUP_FILE
	@echo "\033[1;32m=== Restore completed ===\033[0m"

# ========================================
# Migrations (Goose)
# ========================================

check-goose: ## Ensure goose migration tool is installed
	@command -v $(GOOSE) >/dev/null 2>&1 || { \
		echo "\033[1;31m=== Error: Goose tool not found ===\033[0m"; \
		echo "  Install with: \033[1;33mgo install github.com/pressly/goose/v3/cmd/goose@latest\033[0m"; \
		exit 1; \
	}

DB_URL_LOCAL = $(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?parseTime=true

migration-create: check-goose ## Create a new migration (prompts for name)
	@echo "\033[1;34m=== Creating new migration ===\033[0m"
	@mkdir -p $(MIGRATION_DIR)
	@read -p "  Enter migration name: " name; \
	if [ -z "$$name" ]; then \
		echo "  \033[1;31mError: Migration name is required.\033[0m"; \
		exit 1; \
	fi; \
	$(GOOSE) -dir $(MIGRATION_DIR) create $$name sql

migration-up: check-goose ## Run all up migrations
	@echo "\033[1;34m=== Running migrations up ===\033[0m"
	@$(GOOSE) -dir $(MIGRATION_DIR) mysql "$(DB_URL_LOCAL)" up

migration-down: check-goose ## Rollback the last migration
	@echo "\033[1;34m=== Rolling back migration ===\033[0m"
	@$(GOOSE) -dir $(MIGRATION_DIR) mysql "$(DB_URL_LOCAL)" down

migration-status: check-goose ## Show the current migration status
	@echo "\033[1;34m=== Migration status ===\033[0m"
	@$(GOOSE) -dir $(MIGRATION_DIR) mysql "$(DB_URL_LOCAL)" status

migration-version: check-goose ## Print current DB version
	@echo "\033[1;34m=== Current DB Version ===\033[0m"
	@$(GOOSE) -dir $(MIGRATION_DIR) mysql "$(DB_URL_LOCAL)" version

migration-reset: check-goose ## Run full down and then up
	@echo "\033[1;34m=== Resetting migrations ===\033[0m"
	@$(GOOSE) -dir $(MIGRATION_DIR) mysql "$(DB_URL_LOCAL)" reset
	@$(GOOSE) -dir $(MIGRATION_DIR) mysql "$(DB_URL_LOCAL)" up
	@echo "\033[1;32m=== Reset completed ===\033[0m"

# ========================================
# Ent ORM
# ========================================

ent-generate: ## Generate Ent schema code
	@echo "\033[1;34m=== Generating Ent code ===\033[0m"
	@mkdir -p ./ent/generated
	@go run entgo.io/ent/cmd/ent generate ./ent/schema --target ./ent/generated --feature sql/upsert
	@echo "\033[1;32m=== Ent generation completed ===\033[0m"

ent-init: ## Create new Ent schema (prompts for name)
	@read -p "  Enter schema name: " schema; \
	if [ -z "$$schema" ]; then \
		echo "  \033[1;31mError: Schema name is required.\033[0m"; \
		exit 1; \
	fi; \
	go run entgo.io/ent/cmd/ent new $$schema

# ========================================
# Wire (Dependency Injection)
# ========================================

wire: ## Generate Wire dependency injection code
	@echo "\033[1;34m=== Generating Wire code ===\033[0m"
	@wire ./internal/di/...
	@echo "\033[1;32m=== Wire generation completed ===\033[0m"

wire-check: ## Check Wire setup without generating
	@echo "\033[1;34m=== Checking Wire setup ===\033[0m"
	@wire check ./internal/di/...

wire-show: ## Show Wire dependency graph
	@echo "\033[1;34m=== Wire dependency graph ===\033[0m"
	@wire show ./internal/di/...

# ========================================
# Swagger / OpenAPI
# ========================================

swagger: ## Generate Swagger/OpenAPI documentation
	@echo "\033[1;34m=== Generating Swagger docs ===\033[0m"
	@swag init -g cmd/server/main.go -o docs/swagger --parseDependency --parseInternal
	@echo "\033[1;32m=== Swagger docs generated at docs/swagger/ ===\033[0m"
	@echo "  View at: http://localhost:8080/swagger/index.html"

swagger-fmt: ## Format Swagger comments
	@echo "\033[1;34m=== Formatting Swagger comments ===\033[0m"
	@swag fmt

# ========================================
# TypeScript Types
# ========================================

types: swagger ## Generate TypeScript types from OpenAPI spec
	@echo "\033[1;34m=== Generating TypeScript types ===\033[0m"
	@bash scripts/generate-types.sh

# ========================================
# Code Generation (All)
# ========================================

generate: ent-generate wire ## Generate Ent + Wire code
	@echo "\033[1;32m=== Code generation completed ===\033[0m"

generate-all: ent-generate wire swagger types ## Generate all code (Ent + Wire + Swagger + TypeScript)
	@echo "\033[1;32m=== All code generation completed ===\033[0m"

# ========================================
# Code Quality
# ========================================

lint: ## Run GolangCI-Lint (excluding ent/ and docker/)
	@dirs=$$(find . -type f -name '*.go' \
		-not -path './docker/*' \
		-not -path './ent/*' \
		-exec dirname {} \; | sort -u); \
	golangci-lint run --config .golangci.yml $$dirs

fmt: ## Format Go code (gci + goimports), excluding docker/ and ent/
	@bash -c 'find . \( -path ./docker -o -path ./ent \) -prune -o -name "*.go" -type f -print \
	| xargs gci write -s standard -s default -s "Prefix(github.com/UnoraApp/be)"'
	@bash -c 'FILES=$$(find . \( -path ./docker -o -path ./ent \) -prune -o -name "*.go" -type f -print); goimports -w $$FILES'

gosec-scan: ## Run gosec static analyzer
	@GO111MODULE=on gosec ./...

govulncheck-scan: ## Run govulncheck on all modules
	@$(GO) run golang.org/x/vuln/cmd/govulncheck@latest ./...

security-scan: gosec-scan govulncheck-scan ## Run both gosec and govulncheck

pre-commit-run: ## Run all pre-commit hooks manually
	@pre-commit run --all-files

# ========================================
# Setup
# ========================================

bootstrap: ## Install development dependencies
	@echo "\033[1;34m=== Running bootstrap script ===\033[0m"
	@bash scripts/bootstrap.sh

scaffold-module: ## Generate module skeleton with handlers/services/etc.
	@bash scripts/scaffold.sh

seed: ## Run database seeder locally
	@echo "\033[1;34m=== Running Seeder ===\033[0m"
	@DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) go run cmd/seeder/main.go
	@echo "\033[1;32m=== Seeder completed ===\033[0m"

# ========================================
# Environment Setup
# ========================================

setup-env: ## Copy .env.example to .env
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "\033[1;32m=== .env created from .env.example ===\033[0m"; \
	else \
		echo "\033[1;33m=== .env already exists ===\033[0m"; \
	fi
