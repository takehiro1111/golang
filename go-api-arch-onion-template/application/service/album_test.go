package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"go-api-arch-onion-template/application/domain"
	"go-api-arch-onion-template/pkg"
)

type mockAlbumRepository struct {
	mock.Mock
}

func NewMockAlbumRepository() *mockAlbumRepository {
	return &mockAlbumRepository{}
}

func (m *mockAlbumRepository) Create(album *domain.Album) (*domain.Album, error) {
	args := m.Called(album)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Album), args.Error(1)
}

func (m *mockAlbumRepository) Get(ID int) (*domain.Album, error) {
	args := m.Called(ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Album), args.Error(1)
}
func (m *mockAlbumRepository) Save(album *domain.Album) (*domain.Album, error) {
	args := m.Called(album)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Album), args.Error(1)
}

func (m *mockAlbumRepository) Delete(ID int) error {
	args := m.Called(ID)
	if args.Get(0) == nil {
		return args.Error(1)
	}
	return nil
}

type AlbumServiceSuite struct {
	suite.Suite
	albumService *albumService
}

func TestAlbumServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AlbumServiceSuite))
}

func (suite *AlbumServiceSuite) SetupSuite() {
}

func (suite *AlbumServiceSuite) TestCreate() {
	now := pkg.Str2time("2023-01-01")
	title := "album"
	categoryName := "sports"
	mockAlbumRepository := NewMockAlbumRepository()
	suite.albumService = NewAlbumService(mockAlbumRepository)

	category := domain.Category{Name: domain.CategoryName(categoryName)}
	album := &domain.Album{
		Title:       title,
		ReleaseDate: now,
		Category:    category,
	}

	mockAlbumRepository.On("Create", album).Return(&domain.Album{
		ID:          1,
		Title:       title,
		ReleaseDate: now,
		CategoryID:  1,
		Category:    category,
	}, nil)

	album, err := suite.albumService.Create(album)
	suite.Assert().Nil(err)
	suite.Assert().Equal(title, album.Title)
	suite.Assert().Equal(now, album.ReleaseDate)
	suite.Assert().Equal(category.ID, album.Category.ID)
	suite.Assert().Equal(string(categoryName), string(album.Category.Name))
}

func (suite *AlbumServiceSuite) TestGet() {
	now := pkg.Str2time("2023-01-01")
	title := "album"
	categoryName := "sports"
	mockAlbumRepository := NewMockAlbumRepository()
	suite.albumService = NewAlbumService(mockAlbumRepository)
	category := &domain.Category{
		ID:   1,
		Name: domain.CategoryName(categoryName),
	}
	mockAlbumRepository.On("Get", 1).Return(&domain.Album{
		ID:          1,
		Title:       title,
		ReleaseDate: now,
		CategoryID:  category.ID,
		Category:    *category,
	}, nil)

	album, err := suite.albumService.Get(1)
	suite.Assert().Nil(err)
	suite.Assert().Equal(title, album.Title)
	suite.Assert().Equal(now, album.ReleaseDate)
	suite.Assert().Equal(category.ID, album.Category.ID)
	suite.Assert().Equal(string(categoryName), string(album.Category.Name))
}

func (suite *AlbumServiceSuite) TestUpdate() {
	now := pkg.Str2time("2023-01-01")
	title := "album"
	categoryName := "sports"
	category := domain.Category{Name: domain.CategoryName(categoryName)}
	album := &domain.Album{
		ID:          1,
		Title:       title,
		ReleaseDate: now,
		Category:    category,
	}

	mockAlbumRepository := NewMockAlbumRepository()
	suite.albumService = NewAlbumService(mockAlbumRepository)
	mockAlbumRepository.On("Save", album).Return(&domain.Album{
		ID:          1,
		Title:       title,
		ReleaseDate: now,
		CategoryID:  1,
		Category:    category,
	}, nil)

	album, err := suite.albumService.Save(album)
	suite.Assert().Nil(err)
	suite.Assert().Equal(title, album.Title)
	suite.Assert().Equal(now, album.ReleaseDate)
	suite.Assert().Equal(category.ID, album.Category.ID)
	suite.Assert().Equal(string(categoryName), string(album.Category.Name))
}

func (suite *AlbumServiceSuite) TestDelete() {
	mockAlbumRepository := NewMockAlbumRepository()
	suite.albumService = NewAlbumService(mockAlbumRepository)
	mockAlbumRepository.On("Delete", 1).Return(nil, nil)
	err := suite.albumService.Delete(1)
	suite.Assert().Nil(err)
}
