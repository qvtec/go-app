package router

import (
	"github.com/qvtec/go-app/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(router *gin.RouterGroup, authHandler *handler.AuthHandler) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
	}
}
