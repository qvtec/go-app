package main

import (
	"log"
	"net/http"

	"qvtec/go-app/internal/delivery"
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
	userHandler := delivery.NewUserHandler(userUseCase)

	router := gin.Default()
	delivery.SetupUserRouter(router, userHandler)

	http.ListenAndServe(":8080", router)
}
