# GO APP

GolangでクリーンアーキテクチャのAPIを実装してみる

## Table of Contents

- [GO APP](#go-app)
	- [Table of Contents](#table-of-contents)
	- [Installation](#installation)
	- [Usage](#usage)
		- [test](#test)
	- [Features](#features)
		- [Project structure](#project-structure)
	- [Packages](#packages)
	- [How to add new API](#how-to-add-new-api)
		- [1. migration](#1-migration)
		- [2. domain,usecase,repository,delivery](#2-domainusecaserepositorydelivery)
		- [3. main.go](#3-maingo)
	- [Todo](#todo)
	- [Reference](#reference)

## Installation

```
$ git clone git@github.com:qvtec/go-app.git
$ cd go-app
$ cp .env.example .env
$ make up
$ make migrate-up
```

## Usage

```
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

### test

```
$ make test
```

## Features

### Project structure

```
- build
  - go
    - Dockerfile
  - mysql
    - initdb.d
    - my.cnf
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
    - user_repository_test.go
    - album_repository.go
    - album_repository_test.go
  - delivery                    // Frameworks & Drivers
    - http
      - handler
        - user_handler.go
        - user_handler_test.go
        - album_handler.go
        - album_handler_test.go
      - router
        - user_router.go
        - album_router.go
- pkg
  - db
    - mysql
      -db_con.go
- db
  - migrations
      - 001_create_user_table.up.sql
      - 001_create_user_table.down.sql
- .env
- .env.example
- docker-compose.yml
- go.mod
- go.sum
- Makefile
- README.md
```

```
$ go mod tidy
```

## Packages

* gin-gonic/gin
* go-sql-driver/mysql

* hot reload: air
* auth: jwt
* test: testify
* mock: mockery

## How to add new API

### 1. migration

```
// Create up/down migration files
$ make migrate-create

// -- Add SQL

// Run migrate
$ make migrate
```

### 2. domain,usecase,repository,delivery

* delivery: ルーターとコントローラ呼び出し
* usecase: ビジネスロジック
* repository: DBデータ関連
* domain: データ構成

### 3. main.go

```
albumRepository := repository.NewAlbumRepository(db)
albumUseCase := usecase.NewAlbumUseCase(albumRepository)
albumHandler := httpHandler.NewAlbumHandler(albumUseCase)
httpRouter.SetupAlbumRouter(router, albumHandler)
```

## Todo

- [ ] validation
- [ ] log
- [ ] auth

## Reference

https://github.com/qvtec/go-clean-arch