# GO API

Golangを使用してWebアプリケーション用のAPIを実装

## Table of Contents

- [GO API](#go-api)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Reference](#reference)

## Installation

```
$ git clone git@github.com:qvtec/go-app.git
$ cd go-app
$ make up
$ make migrate-up
```

## Usage

```
$ make run

// ALL
$ curl http://localhost:8080/api/v1/users

// CREATE
$ curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com"}' http://localhost:8080/api/v1/users

// GetByID
$ curl http://localhost:8080/api/v1/users/{user_id}

// UPDATE
$ curl -X PUT -H "Content-Type: application/json" -d '{"name": "John Smith", "email": "john.smith@example.com"}' http://localhost:8080/api/v1/users/{user_id}

// DELETE
$ curl -X DELETE http://localhost:8080/api/v1/users/{user_id}
```
```

## Features

### Project structure

```
- build
  - server
    - Dockerfile
  - docker-compose.yml
- cmd
  - main.go
- internal
  - domain                      // Entities
    - user.go
    - album.go
    - error.go
  - usecase                     // Use cases
    - user_usecase.go
    - album_usecase.go
  - repository                  // Interface
    - user_repository.go
    - album_repository.go
    - common_repository.go
  - delivery                    // Frameworks & Drivers
    - http
      - handler
        - user_handler.go
        - album_handler.go
      - router
        - user_router.go
- pkg
  - db
    - mysql
      -db_con.go
- db
  - migrations
      - 001_create_user_table.up.sql
      - 001_create_user_table.down.sql
- go.mod
- go.sum
- Makefile
- README.md
```

```
$ go mod tidy
```

## Reference

https://github.com/qvtec/go-clean-arch