include .env
MIGRATIONS_DIR = ./cmd/migrate/migrations

.PHONY: migrate-up
migrate-up:
	@goose -dir $(MIGRATIONS_DIR) postgres "$(DB_CONN_STRING)" up

.PHONY: migrate-down
migrate-down:
	@goose -dir $(MIGRATIONS_DIR) postgres "$(DB_CONN_STRING)" down

.PHONY: migrate-status
migrate-status:
	@goose -dir $(MIGRATIONS_DIR) postgres "$(DB_CONN_STRING)" status

.PHONY: migrate-create
migrate-create:
	@goose -dir $(MIGRATIONS_DIR) postgres "$(DB_CONN_STRING)" create $(name) -s sql
