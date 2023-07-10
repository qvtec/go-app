package router

import (
	"qvtec/go-app/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupAlbumRouter(router *gin.Engine, albumHandler *handler.AlbumHandler) {
	v1 := router.Group("/api/v1")

	v1.GET("/albums", albumHandler.GetAll)
	v1.POST("/albums", albumHandler.Create)
	v1.GET("/albums/:id", albumHandler.Get)
	v1.PUT("/albums/:id", albumHandler.Update)
	v1.DELETE("/albums/:id", albumHandler.Delete)
}
