package main

import (
	"log"
	"net/http"
	"os"

	httpHandler "github.com/qvtec/go-app/internal/delivery/http/handler"
	httpRouter "github.com/qvtec/go-app/internal/delivery/http/router"
	"github.com/qvtec/go-app/internal/repository"
	"github.com/qvtec/go-app/internal/usecase"
	"github.com/qvtec/go-app/pkg/db"

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

	port := os.Getenv("SERVER_PORT")
	http.ListenAndServe(":" + port, router)
}
