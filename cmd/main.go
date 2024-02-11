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

	authRepository := repository.NewAuthRepository(db)
	authUseCase := usecase.NewAuthUseCase(authRepository)
	authHandler := httpHandler.NewAuthHandler(authUseCase)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	httpRouter.SetupUserRouter(v1, userHandler)
	httpRouter.SetupAuthRouter(v1, authHandler)

	port := os.Getenv("SERVER_PORT")
	http.ListenAndServe(":"+port, router)
}
