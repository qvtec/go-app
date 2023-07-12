# Docker ////////////////////////////////////////////////////////////////////////////
build:
	@docker-compose -f build/docker-compose.yml build mysql

run:
	@docker-compose -f build/docker-compose.yml up -d mysql
	@docker-compose -f build/docker-compose.yml up app

down:
	@docker-compose -f build/docker-compose.yml down

clean:
	@docker-compose -f build/docker-compose.yml down -v --rmi all

.PHONY: build run down clean

# Test //////////////////////////////////////////////////////////////////////////////
GO := docker-compose -f build/docker-compose.yml run app go

.PHONY: test

test:
	$(GO) test -cover -coverprofile=coverage.out ./...

coverage:
	$(GO) tool cover -html=coverage.out -o coverage.html

# Database Migrations ////////////////////////////////////////////////////////////////
include ./build/.env
MYSQL_DSN := "mysql://${DB_USER}:${DB_PASS}@tcp(mysql:${DB_PORT})/${DB_NAME}"
MIGRATE := docker-compose -f build/docker-compose.yml run migrate -path=/migrations/ -database $(MYSQL_DSN)

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