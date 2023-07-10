package main

import (
	"log"
	"net/http"

	httpHandler "qvtec/go-app/internal/delivery/http/handler"
	httpRouter "qvtec/go-app/internal/delivery/http/router"
	"qvtec/go-app/internal/repository"
	"qvtec/go-app/internal/usecase"
	"qvtec/go-app/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.NewMySQLDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := httpHandler.NewUserHandler(userUseCase)
	httpRouter.SetupUserRouter(router, userHandler)

	albumRepository := repository.NewAlbumRepository(db)
	albumUseCase := usecase.NewAlbumUseCase(albumRepository)
	albumHandler := httpHandler.NewAlbumHandler(albumUseCase)
	httpRouter.SetupAlbumRouter(router, albumHandler)

	http.ListenAndServe(":8080", router)
}
