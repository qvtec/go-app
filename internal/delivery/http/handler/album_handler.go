package handler

import (
	"net/http"
	"qvtec/go-app/internal/domain"
	"qvtec/go-app/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	albumUseCase usecase.AlbumUseCase
}

func NewAlbumHandler(albumUseCase usecase.AlbumUseCase) *AlbumHandler {
	return &AlbumHandler{
		albumUseCase: albumUseCase,
	}
}


func (h *AlbumHandler) GetAll(c *gin.Context) {
	albums, err := h.albumUseCase.GetAll()
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get albums"})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, albums)
}


func (h *AlbumHandler) Create(c *gin.Context) {
	var album domain.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.albumUseCase.Create(&album)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create album"})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Album created successfully"})
}


func (h *AlbumHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	album, err := h.albumUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get album"})
		return
	}

	if album == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"album": album})
}


func (h *AlbumHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	var updateAlbum domain.Album
	if err := c.ShouldBindJSON(&updateAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updateAlbum.ID = id
	err = h.albumUseCase.Update(&updateAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update album"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album updated successfully"})
}


func (h *AlbumHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	err = h.albumUseCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete album"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}

