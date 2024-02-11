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
	- [Reference](#reference)

## Installation

```sh
$ git clone git@github.com:qvtec/go-app.git
$ cd go-app
$ cp .env.example .env
$ make up
$ make migrate-up
```

## Usage

```sh
# Login
curl -X POST -H "Content-Type: application/json" -d '{"email": "test@test.com", "password": "password"}' http://localhost:8080/api/v1/auth/login

# ALL
$ curl -X GET -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/v1/users
# CREATE
$ curl -X POST -H "Content-Type: application/json"  -H "Authorization: Bearer $TOKEN" -d '{"name": "John Doe", "email": "john@example.com", "password": "password"}' http://localhost:8080/api/v1/users
# GetByID
$ curl -X GET -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/v1/users/{user_id}
# UPDATE
$ curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" -d '{"name": "John Smith"}' http://localhost:8080/api/v1/users/{user_id}
# DELETE
$ curl -X DELETE -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/v1/users/{user_id}
```

### test

```
$ make test
```

## Features

### Project structure

```sh
├── build
│   ├── go
│   │   └── Dockerfile
│   └── mysql
│       ├── initdb.d
│       └── my.cnf
├── cmd
│   └── main.go
├── db
│   └── migrations
├── internal
│   ├── delivery                    // Frameworks & Drivers
│   │   └── http
│   │       ├── handler
│   │       ├── middleware
│   │       └── router
│   ├── domain                      // Entities
│   ├── repository                  // Interface
│   └── usecase                     // Use cases
├── mocks
├── pkg
│   └── db
├── go.mod
├── go.sum
├── docker-compose.yml
├── Makefile
└── README.md
```

```sh
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

```sh
# Create up/down migration files
$ make migrate-create

# -- Add SQL

# Run migrate
$ make migrate
```

### 2. domain,usecase,repository,delivery

* delivery: ルーターとコントローラ呼び出し
* usecase: ビジネスロジック
* repository: DBデータ関連
* domain: データ構成

### 3. main.go

```sh
albumRepository := repository.NewAlbumRepository(db)
albumUseCase := usecase.NewAlbumUseCase(albumRepository)
albumHandler := httpHandler.NewAlbumHandler(albumUseCase)
httpRouter.SetupAlbumRouter(router, albumHandler)
```

## Reference

https://github.com/qvtec/go-clean-arch