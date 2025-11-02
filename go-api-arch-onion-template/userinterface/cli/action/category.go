package action

import (
	"go-api-arch-onion-template/application/domain"
	"go-api-arch-onion-template/application/service"
	"go-api-arch-onion-template/pkg/logger"
)

var CategoryName string

type CategoryAction struct {
	categoryService service.CategoryService
}

func NewCategoryAction(categoryService service.CategoryService) *CategoryAction {
	return &CategoryAction{
		categoryService: categoryService,
	}
}

func (a *CategoryAction) CreateCategory(name string) (*domain.Category, error) {
	category, err := domain.NewCategory(name)
	if err != nil {
		return nil, err
	}
	createdCategory, err := a.categoryService.GetOrCreate(category)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return createdCategory, nil
}
