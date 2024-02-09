include .env
MYSQL_DSN := "mysql://${DB_USER}:${DB_PASS}@tcp(mysql:${DB_PORT})"

# Docker ////////////////////////////////////////////////////////////////////////////
build:
	docker-compose build --no-cache
up:
	docker compose up -d
down:
	docker compose down --remove-orphans
destroy:
	docker compose down --rmi all --volumes --remove-orphans
ps:
	docker compose ps

.PHONY: build up down destroy ps

# Docker container ////////////////////////////////////////////////////////////////////////////
app:
	docker compose exec app sh
db:
	docker compose exec mysql sh
sql:
	docker compose exec mysql bash -c 'mysql -u ${DB_USER} -p${DB_PASS} ${DB_NAME}'

.PHONY: app db sql

# Log ////////////////////////////////////////////////////////////////////////////
logs:
	docker compose logs
logs-watch:
	docker compose logs --follow

# Database Migrations ////////////////////////////////////////////////////////////////
MIGRATE := docker compose run --rm migrate -path=/migrations/ -database $(MYSQL_DSN)/${DB_NAME}

migrate:
	$(MIGRATE) up
migrate-down:
	$(MIGRATE) down
migrate-reset:
	$(MIGRATE) drop
	$(MIGRATE) up
migrate-create: ## Create a set of up/down migrations with a specified name.
	@ read -p "Enter the name of the new migration: " Name; \
	$(MIGRATE) create -ext sql -dir ./db/migrations/ $${Name}

# Test //////////////////////////////////////////////////////////////////////////////
TEST_MIGRATE := docker compose run --rm migrate -path=/migrations/ -database $(MYSQL_DSN)/testing

mock:
	docker compose exec app mockery --all

test:
	@$(TEST_MIGRATE) drop
	@$(TEST_MIGRATE) up
	docker compose run --rm app go test -cover -coverprofile=coverage.out ./...
	docker compose run --rm app go tool cover -html=coverage.out -o coverage.html

test-debug:
ifdef TEST_PATH
	docker compose run --rm app go test -cover -coverprofile=coverage.out $(TEST_PATH)
else
	@echo "Usage: make test-debug TEST_PATH=<test path> ex: make test-debug TEST_PATH=./pkg/db/"
endif
