# load environment variables if .env exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

DB_CONTAINER ?= wry_postgres
DB_USER ?= $(DB_USERNAME)
ifeq ($(DB_USER),)
    DB_USER := postgres
endif
DB_NAME_VAL ?= $(DB_NAME)
ifeq ($(DB_NAME_VAL),)
    DB_NAME_VAL := postgres
endif

.PHONY: help up down db-wait migrate seed db-setup run dev certs sqlc

help: ## display available targets
	@echo "available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

certs: ## generate rsa keys if missing
	@mkdir -p certs
	@if [ ! -f certs/private.pem ] || [ ! -f certs/public.pem ]; then \
		echo "generating rsa keys in ./certs ..."; \
		openssl genrsa -out certs/private.pem 2048; \
		openssl rsa -in certs/private.pem -pubout -out certs/public.pem; \
		echo "rsa keys generated successfully."; \
	else \
		echo "rsa keys already exist in ./certs."; \
	fi

up: ## start database container
	@echo "starting database container..."
	@docker compose up -d db

db-wait: ## wait for postgres container readiness
	@echo "waiting for database to be ready..."
	@until docker exec $(DB_CONTAINER) pg_isready -U $(DB_USER) -d $(DB_NAME_VAL) >/dev/null 2>&1; do \
		sleep 1; \
	done
	@echo "database is ready."

migrate: db-wait ## apply ddl schema to database
	@echo "applying schema.sql..."
	@docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME_VAL) < internal/database/schema.sql
	@echo "schema applied successfully."

seed: db-wait ## populate database with seed data
	@echo "populating database with seed.sql..."
	@docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME_VAL) < internal/database/seed.sql
	@echo "seed applied successfully."

db-setup: up db-wait migrate seed ## start container, apply schema, and populate seed
	@echo "database setup completed."

sqlc: ## regenerate go code from sql queries and schema
	@echo "running sqlc generate..."
	@sqlc generate

run: ## run go application
	@go run ./cmd/api

dev: certs db-setup run ## full workflow: generate certs, setup db, and start app

down: ## stop and remove docker containers
	@docker compose down
