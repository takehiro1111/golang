package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"go-api-arch-onion-template/application/domain"
	"go-api-arch-onion-template/pkg"
	"go-api-arch-onion-template/userinterface/web/gin/api"
)

type MockAlbumUseCase struct {
	mock.Mock
}

func NewMockAlbumUseCase() *MockAlbumUseCase {
	return &MockAlbumUseCase{}
}

func (m *MockAlbumUseCase) Create(album *domain.Album) (*domain.Album, error) {
	args := m.Called(album)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Album), args.Error(1)
}

func (m *MockAlbumUseCase) Get(ID int) (*domain.Album, error) {
	args := m.Called(ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Album), args.Error(1)
}
func (m *MockAlbumUseCase) Save(album *domain.Album) (*domain.Album, error) {
	args := m.Called(album)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Album), args.Error(1)
}

func (m *MockAlbumUseCase) Delete(ID int) error {
	args := m.Called(ID)
	if args.Get(0) == nil {
		return args.Error(1)
	}
	return nil
}

type AlbumHandlersSuite struct {
	suite.Suite
	albumHandler *AlbumHandler
}

func TestAlbumHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(AlbumHandlersSuite))
}

func (suite *AlbumHandlersSuite) TestCreate() {
	now := pkg.Str2time("2023-01-01")
	mockUseCase := NewMockAlbumUseCase()
	album := &domain.Album{
		Title:       "album",
		ReleaseDate: now,
		Category:    domain.Category{Name: "sports"},
	}

	mockUseCase.On("Create", album).Return(&domain.Album{
		ID:          1,
		Title:       "album",
		ReleaseDate: now,
		CategoryID:  1,
		Category: domain.Category{
			ID:   1,
			Name: "sports",
		},
	}, nil)
	suite.albumHandler = NewAlbumHandler(mockUseCase)

	request, _ := api.NewCreateAlbumRequest("/api/v1", api.CreateAlbumJSONRequestBody{
		Title:       "album",
		Category:    api.Category{Name: "sports"},
		ReleaseDate: api.ReleaseDate{Time: now},
	})
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request = request

	suite.albumHandler.CreateAlbum(ginContext)

	suite.Assert().Equal(http.StatusCreated, w.Code)
	bodyBytes, _ := io.ReadAll(w.Body)
	var albumGetResponse api.AlbumResponse
	err := json.Unmarshal(bodyBytes, &albumGetResponse)
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusCreated, w.Code)
	suite.Assert().Equal("album", albumGetResponse.Title)
	suite.Assert().Equal("sports", string(albumGetResponse.Category.Name))
	suite.Assert().NotNil(albumGetResponse.ReleaseDate)
}

func (suite *AlbumHandlersSuite) TestCreateRequestBodyFailure() {
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("POST", "/api/v1/album", nil)
	req.Header.Add("Content-Type", "application/json")
	ginContext.Request = req

	suite.albumHandler.CreateAlbum(ginContext)
	suite.Assert().Equal(http.StatusBadRequest, w.Code)
	suite.Assert().JSONEq(`{"message": "invalid request"}`, w.Body.String())
}

func (suite *AlbumHandlersSuite) TestCreateFailure() {
	now := pkg.Str2time("2023-01-01")
	mockUseCase := NewMockAlbumUseCase()
	album := &domain.Album{
		Title:       "album",
		ReleaseDate: now,
		Category:    domain.Category{Name: "sports"},
	}

	mockUseCase.On("Create", album).Return(nil,
		errors.New("invalid"))
	suite.albumHandler = NewAlbumHandler(mockUseCase)

	request, _ := api.NewCreateAlbumRequest("/api/v1", api.CreateAlbumJSONRequestBody{
		Title:       "album",
		Category:    api.Category{Name: "sports"},
		ReleaseDate: api.ReleaseDate{Time: now},
	})
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request = request

	suite.albumHandler.CreateAlbum(ginContext)

	suite.Assert().Equal(http.StatusInternalServerError, w.Code)
	suite.Assert().JSONEq(`{"message":"invalid"}`, w.Body.String())
}

func (suite *AlbumHandlersSuite) TestGet() {
	now := pkg.Str2time("2023-01-01")
	mockUseCase := NewMockAlbumUseCase()
	mockUseCase.On("Get", 1).Return(&domain.Album{
		ID:          1,
		Title:       "album",
		ReleaseDate: now,
		CategoryID:  1,
		Category: domain.Category{
			ID:   1,
			Name: "sports",
		},
	}, nil)
	suite.albumHandler = NewAlbumHandler(mockUseCase)

	request, _ := api.NewGetAlbumByIdRequest("/api/v1", 1)
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request = request

	suite.albumHandler.GetAlbumById(ginContext, 1)

	bodyBytes, _ := io.ReadAll(w.Body)
	var albumGetResponse api.AlbumResponse
	err := json.Unmarshal(bodyBytes, &albumGetResponse)
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusOK, w.Code)
	suite.Assert().Equal("album", albumGetResponse.Title)
	suite.Assert().Equal("sports", string(albumGetResponse.Category.Name))
	suite.Assert().NotNil(albumGetResponse.ReleaseDate)
}

func (suite *AlbumHandlersSuite) TestGetNoAlbumFailure() {
	mockUseCase := NewMockAlbumUseCase()
	mockUseCase.On("Get", 1111).Return(nil,
		errors.New("invalid"))
	suite.albumHandler = NewAlbumHandler(mockUseCase)

	request, _ := api.NewGetAlbumByIdRequest("/api/v1", 1111)
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request = request

	suite.albumHandler.GetAlbumById(ginContext, 1111)

	suite.Assert().Equal(http.StatusInternalServerError, w.Code)
	suite.Assert().JSONEq(`{"message": "invalid"}`, w.Body.String())
}

func (suite *AlbumHandlersSuite) TestUpdate() {

	now := pkg.Str2time("2023-01-01")
	mockUseCase := NewMockAlbumUseCase()
	title := "updated"
	categoryName := "food"

	album := &domain.Album{
		ID:       1,
		Title:    title,
		Category: domain.Category{Name: domain.CategoryName(categoryName)},
	}

	mockUseCase.On("Save", album).Return(&domain.Album{
		ID:          1,
		Title:       title,
		ReleaseDate: now,
		CategoryID:  1,
		Category: domain.Category{
			ID:   1,
			Name: domain.CategoryName(categoryName),
		},
	}, nil)

	suite.albumHandler = NewAlbumHandler(mockUseCase)

	category := api.Category{
		Name: api.CategoryName(categoryName),
	}
	request, _ := api.NewUpdateAlbumByIdRequest("/api/v1", 1,
		api.UpdateAlbumByIdJSONRequestBody{
			Title:    &title,
			Category: &category,
		},
	)
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request = request

	suite.albumHandler.UpdateAlbumById(ginContext, 1)

	bodyBytes, _ := io.ReadAll(w.Body)
	var albumGetResponse api.AlbumResponse
	err := json.Unmarshal(bodyBytes, &albumGetResponse)
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusOK, w.Code)
	suite.Assert().Equal("updated", albumGetResponse.Title)
	suite.Assert().Equal("food", string(albumGetResponse.Category.Name))
	suite.Assert().NotNil(albumGetResponse.ReleaseDate)
}

func (suite *AlbumHandlersSuite) TestUpdateRequestBodyFailure() {
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("PATCH", "/api/v1/album", nil)
	req.Header.Add("Content-Type", "application/json")
	ginContext.Request = req

	suite.albumHandler.CreateAlbum(ginContext)

	suite.Assert().Equal(http.StatusBadRequest, w.Code)
	suite.Assert().JSONEq(`{"message": "invalid request"}`, w.Body.String())
}

func (suite *AlbumHandlersSuite) TestUpdateAlbumFailure() {
	mockUseCase := NewMockAlbumUseCase()
	title := "updated"
	categoryName := "food"
	album := &domain.Album{
		ID:       1111,
		Title:    title,
		Category: domain.Category{Name: domain.CategoryName(categoryName)},
	}

	mockUseCase.On("Save", album).Return(nil,
		errors.New("invalid"))
	suite.albumHandler = NewAlbumHandler(mockUseCase)

	category := api.Category{
		Name: api.CategoryName(categoryName),
	}
	request, _ := api.NewUpdateAlbumByIdRequest("/api/v1", 1111,
		api.UpdateAlbumByIdJSONRequestBody{
			Title:    &title,
			Category: &category,
		},
	)
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request = request

	suite.albumHandler.UpdateAlbumById(ginContext, 1111)

	suite.Assert().Equal(http.StatusInternalServerError, w.Code)
	suite.Assert().JSONEq(`{"message": "invalid"}`, w.Body.String())
}

func (suite *AlbumHandlersSuite) TestDelete() {
	mockUseCase := NewMockAlbumUseCase()
	mockUseCase.On("Delete", 1).Return(nil, nil)
	suite.albumHandler = NewAlbumHandler(mockUseCase)

	request, _ := api.NewDeleteAlbumByIdRequest("/api/v1", 1)
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request = request

	suite.albumHandler.DeleteAlbumById(ginContext, 1)

	suite.Assert().Equal(http.StatusNoContent, w.Code)
}

func (suite *AlbumHandlersSuite) TestDeleteAlbumFailure() {
	mockUseCase := NewMockAlbumUseCase()
	mockUseCase.On("Delete", 1111).Return(nil, errors.New("invalid"))
	suite.albumHandler = NewAlbumHandler(mockUseCase)

	request, _ := api.NewDeleteAlbumByIdRequest("/api/v1", 1111)
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request = request

	suite.albumHandler.DeleteAlbumById(ginContext, 1111)

	suite.Assert().Equal(http.StatusInternalServerError, w.Code)
	suite.Assert().JSONEq(`{"message": "invalid"}`, w.Body.String())
}
