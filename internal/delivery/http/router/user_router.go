package router

import (
	"github.com/qvtec/go-app/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(router *gin.Engine, userHandler *handler.UserHandler) {
	v1 := router.Group("/api/v1")

	v1.GET("/users", userHandler.GetAll)
	v1.POST("/users", userHandler.Create)
	v1.GET("/users/:id", userHandler.Get)
	v1.PUT("/users/:id", userHandler.Update)
	v1.DELETE("/users/:id", userHandler.Delete)
}
