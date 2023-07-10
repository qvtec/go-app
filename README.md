# GO API

Golangを使用してWebアプリケーション用のAPIを実装

## Table of Contents

- [GO API](#go-api)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Features](#features)
    - [Project structure](#project-structure)
  - [API追加方法](#api追加方法)
    - [1. migration追加](#1-migration追加)
    - [2. domain,usecase,repository,delivery追加](#2-domainusecaserepositorydelivery追加)
    - [3. main.go追加](#3-maingo追加)
  - [対応予定](#対応予定)
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

## API追加方法

### 1. migration追加

```
// マイグレーションファイル作成
$ make migrate-create

// -- テーブル指定とファイルの中身追加

// マイグレーション実行
$ make migrate-up
```

### 2. domain,usecase,repository,delivery追加

* delivery: ルーターとコントローラ呼び出し
* usecase: ビジネスロジック
* repository: DBデータ関連
* domain: データ構成

### 3. main.go追加

```
albumRepository := repository.NewAlbumRepository(db)
albumUseCase := usecase.NewAlbumUseCase(albumRepository)
albumHandler := httpHandler.NewAlbumHandler(albumUseCase)
httpRouter.SetupAlbumRouter(router, albumHandler)
```

## 対応予定

- [ ] バリデーション
- [ ] ログ
- [ ] テスト
- [ ] 認証

## Reference

https://github.com/qvtec/go-clean-arch