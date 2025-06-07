MIGRATE_DB_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

migration_up:
	migrate -path migrations/ -database "$(MIGRATE_DB_URL)" -verbose up

migration_down:
	migrate -path migrations/ -database "$(MIGRATE_DB_URL)" -verbose down

migration_fix:
	migrate -path migrations/ -database "$(MIGRATE_DB_URL)" force $(VERSION)
