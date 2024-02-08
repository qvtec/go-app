# Docker ////////////////////////////////////////////////////////////////////////////
build:
	docker compose build mysql
run:
	docker compose up -d mysql
	docker compose up app
down:
	docker compose down
clean:
	docker compose down -v --rmi all

# Test //////////////////////////////////////////////////////////////////////////////
GO := docker compose run app go

test:
	$(GO) test -cover -coverprofile=coverage.out ./...
coverage:
	$(GO) tool cover -html=coverage.out -o coverage.html

# Database Migrations ////////////////////////////////////////////////////////////////
include .env
MYSQL_DSN := "mysql://${DB_USER}:${DB_PASS}@tcp(mysql:${DB_PORT})/${DB_NAME}"
MIGRATE := docker compose run migrate -path=/migrations/ -database $(MYSQL_DSN)

migrate-up:
	$(MIGRATE) up
migrate-down:
	$(MIGRATE) down
migrate-reset:
	$(MIGRATE) drop
	$(MIGRATE) up
migrate-create: ## Create a set of up/down migrations with a specified name.
	@ read -p "Enter the name of the new migration: " Name; \
	$(MIGRATE) create -ext sql -dir ./db/migrations/ $${Name}