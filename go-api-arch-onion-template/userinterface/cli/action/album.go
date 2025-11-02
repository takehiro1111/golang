package action

import (
	"time"

	"go-api-arch-onion-template/application/domain"
	"go-api-arch-onion-template/application/service"
	"go-api-arch-onion-template/pkg/logger"
)

var AlbumName string

type AlbumAction struct {
	albumService service.AlbumService
}

func NewAlbumAction(albumService service.AlbumService) *AlbumAction {
	return &AlbumAction{
		albumService: albumService,
	}
}

func (a *AlbumAction) CreateAlbum(title, categoryName string) (*domain.Album, error) {
	category, err := domain.NewCategory(categoryName)
	if err != nil {
		return nil, err
	}
	album := &domain.Album{
		Title:       title,
		ReleaseDate: time.Now(),
		Category:    *category,
	}

	createdAlbum, err := a.albumService.Create(album)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return createdAlbum, nil
}
