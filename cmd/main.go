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

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := httpHandler.NewUserHandler(userUseCase)

	router := gin.Default()
	httpRouter.SetupUserRouter(router, userHandler)

	http.ListenAndServe(":8080", router)
}
