# Validate environment variables
check-env:
	@if [ -z "$(POSTGRES_USER)" ] || [ -z "$(POSTGRES_PASSWORD)" ] || [ -z "$(POSTGRES_HOST)" ] || [ -z "$(POSTGRES_PORT)" ] || [ -z "$(POSTGRES_DB)" ]; then \
		echo "Error: Missing required environment variables (POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB)"; \
		exit 1; \
	fi

MIGRATE_DB_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

migration_up: check-env
	migrate -path migrations/ -database "$(MIGRATE_DB_URL)" -verbose up

migration_down: check-env
	migrate -path migrations/ -database "$(MIGRATE_DB_URL)" -verbose down

migration_fix: check-env
	migrate -path migrations/ -database "$(MIGRATE_DB_URL)" force $(VERSION)

run:
	air

build:
	docker-compose up -d --build
up:
	docker compose up -d 
down:
	docker compose down --remove-orphans
downrm:
	docker compose down -v
dev: 
	docker compose down && docker compose up -d --build && docker compose logs -f
