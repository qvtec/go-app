package router

import (
	"github.com/qvtec/go-app/internal/delivery/http/handler"
	"github.com/qvtec/go-app/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(router *gin.RouterGroup, userHandler *handler.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.Use(middleware.AuthMiddleware)
		userGroup.GET("", userHandler.GetAll)
		userGroup.POST("", userHandler.Create)
		userGroup.GET("/:id", userHandler.GetByID)
		userGroup.PUT("/:id", userHandler.Update)
		userGroup.DELETE("/:id", userHandler.Delete)
	}
}
