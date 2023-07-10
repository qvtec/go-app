# Database
MYSQL_USER ?= user
MYSQL_PASSWORD ?= password
MYSQL_ADDRESS ?= 127.0.0.1:3306
MYSQL_DATABASE ?= docker

# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)

# --- Tooling & Variables ----------------------------------------------------------------
include ./build/make/tools.Makefile

# ~~~ Development Environment ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
build:
	@docker-compose -f build/docker-compose.yml build mysql

up:
	@docker-compose -f build/docker-compose.yml up mysql

down:
	@docker-compose -f build/docker-compose.yml down

clean:
	@docker-compose -f build/docker-compose.yml down -v --rmi all

run:
	go run ./cmd/main.go

.PHONY: build up down clean db run

# ~~~ Database Migrations ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
MYSQL_DSN := "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)"

migrate-up: $(MIGRATE)
	migrate  -database $(MYSQL_DSN) -path=./db/migrations up

migrate-down: $(MIGRATE)
	migrate  -database $(MYSQL_DSN) -path=misc/migrations down 1

migrate-drop: $(MIGRATE)
	migrate  -database $(MYSQL_DSN) -path=misc/migrations drop

migrate-create: $(MIGRATE) ## Create a set of up/down migrations with a specified name.
	@ read -p "Please provide name for the migration: " Name; \
	migrate create -ext sql -dir misc/migrations $${Name}