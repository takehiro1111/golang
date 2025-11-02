package service

import (
	"go-api-arch-onion-template/application/domain"
	"go-api-arch-onion-template/infrastructure/database/repository"
)

type AlbumService interface {
	Create(album *domain.Album) (*domain.Album, error)
	Get(ID int) (*domain.Album, error)
	Save(*domain.Album) (*domain.Album, error)
	Delete(ID int) error
}

type albumService struct {
	albumRepository repository.AlbumRepository
}

func NewAlbumService(albumRepository repository.AlbumRepository) *albumService {
	return &albumService{
		albumRepository: albumRepository,
	}
}

func (a *albumService) Create(album *domain.Album) (*domain.Album, error) {
	return a.albumRepository.Create(album)
}

func (a *albumService) Get(ID int) (*domain.Album, error) {
	return a.albumRepository.Get(ID)
}

func (a *albumService) Save(album *domain.Album) (*domain.Album, error) {
	return a.albumRepository.Save(album)
}

func (a *albumService) Delete(ID int) error {
	return a.albumRepository.Delete(ID)
}
