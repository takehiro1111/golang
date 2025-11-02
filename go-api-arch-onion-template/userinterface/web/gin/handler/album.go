package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-api-arch-onion-template/application/domain"
	"go-api-arch-onion-template/application/service"
	"go-api-arch-onion-template/pkg/logger"
	"go-api-arch-onion-template/userinterface/web/gin/api"
	"go-api-arch-onion-template/userinterface/web/gin/response"
)

type AlbumHandler struct {
	albumService service.AlbumService
}

func NewAlbumHandler(albumService service.AlbumService) *AlbumHandler {
	return &AlbumHandler{
		albumService: albumService,
	}
}

func (a *AlbumHandler) CreateAlbum(c *gin.Context) {
	var requestBody api.CreateAlbumJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Warn(err.Error())
		c.JSON(http.StatusBadRequest, &api.ErrorResponse{Message: err.Error()})
		return
	}

	category, err := domain.NewCategory(string(requestBody.Category.Name))
	if err != nil {
		logger.Warn(err.Error())
		c.JSON(http.StatusBadRequest, &api.ErrorResponse{Message: err.Error()})
		return
	}

	album := &domain.Album{
		Title:       requestBody.Title,
		ReleaseDate: requestBody.ReleaseDate.Time,
		Category:    *category,
	}

	createdAlbum, err := a.albumService.Create(album)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, &api.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.AlbumToResponse(createdAlbum))
}

func (a *AlbumHandler) GetAlbumById(c *gin.Context, ID int) {
	album, err := a.albumService.Get(ID)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, &api.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.AlbumToResponse(album))
}

func (a *AlbumHandler) UpdateAlbumById(c *gin.Context, ID int) {
	var requestBody api.UpdateAlbumByIdJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Warn(err.Error())
		c.JSON(http.StatusBadRequest, &api.ErrorResponse{Message: err.Error()})
		return
	}

	category, err := domain.NewCategory(string(requestBody.Category.Name))
	if err != nil {
		logger.Warn(err.Error())
		c.JSON(http.StatusBadRequest, &api.ErrorResponse{Message: err.Error()})
		return
	}
	album := &domain.Album{ID: ID, Title: *requestBody.Title, Category: *category}

	updatedAlbum, err := a.albumService.Save(album)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, &api.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.AlbumToResponse(updatedAlbum))
}

func (a *AlbumHandler) DeleteAlbumById(c *gin.Context, ID int) {
	if err := a.albumService.Delete(ID); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, &api.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
