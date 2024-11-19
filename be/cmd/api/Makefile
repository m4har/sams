# Makefile for Windows
# ctrl + shift + p -> Type "Convert Indentation to Tabs"
POSTGRESQL_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATION_PATH=../../internal/infrastructure/migrations
MAIN_PATH=cmd/api/main.go
PROJECT_ROOT=../../

# Include environment variables
include .env

# Use CMD as shell
SHELL=cmd.exe
.SHELLFLAGS = /C

.PHONY: help create-db drop-db migrate-up migrate-down migrate-create migration-status install-air swagger

# Install Air for hot reload
install-air:
		@echo "Installing Air for hot reloading..."
		go install github.com/air-verse/air@latest
		@echo "Air installed successfully. Use 'air -v' to verify installation"

# Run commands
run-dev:
		air -c .air.toml

run-prod:
		go build -o bin/api.exe $(MAIN_PATH)
		./bin/api.exe

# Database management
create-db:
		psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "CREATE DATABASE $(DB_NAME);"

drop-db:
		psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "DROP DATABASE IF EXISTS $(DB_NAME);"

# Migration commands
migrate-up:
		migrate -database "$(POSTGRESQL_URL)" -path $(MIGRATION_PATH) up

migrate-down:
		migrate -database "$(POSTGRESQL_URL)" -path $(MIGRATION_PATH) down

migrate-create:
		@echo Enter migration name:
		@set /p NAME= && migrate create -ext sql -dir $(MIGRATION_PATH) -seq %NAME%

migration-status:
		migrate -database "$(POSTGRESQL_URL)" -path $(MIGRATION_PATH) version

# Generate swagger documentation
swagger:
		cd ../.. && swag init -g internal/interfaces/http/server.go --output ./docs --parseDependency --parseInternal

# Help command
help:
		@echo Available commands:
		@echo   make install-air     - Install Air for hot reloading
		@echo   make run-dev        - Run application in development mode with hot reload
		@echo   make run-prod       - Run application in production mode
		@echo   make create-db      - Create database
		@echo   make drop-db        - Drop database
		@echo   make migrate-up     - Run all migrations
		@echo   make migrate-down   - Rollback migrations
		@echo   make migrate-create - Create new migration