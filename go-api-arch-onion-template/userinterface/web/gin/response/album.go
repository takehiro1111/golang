package response

import (
	"go-api-arch-onion-template/application/domain"
	"go-api-arch-onion-template/userinterface/web/gin/api"
)

func AlbumToResponse(album *domain.Album) *api.AlbumResponse {
	return &api.AlbumResponse{
		Id:          album.ID,
		Title:       album.Title,
		ReleaseDate: api.ReleaseDate{Time: album.ReleaseDate},
		Category: api.Category{
			Id:   &album.Category.ID,
			Name: api.CategoryName(album.Category.Name),
		},
	}
}
