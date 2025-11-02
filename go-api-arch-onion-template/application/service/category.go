package service

import (
	"go-api-arch-onion-template/application/domain"
	"go-api-arch-onion-template/infrastructure/database/repository"
)

type (
	CategoryService interface {
		GetOrCreate(category *domain.Category) (*domain.Category, error)
	}
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *categoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (a *categoryService) GetOrCreate(category *domain.Category) (*domain.Category, error) {
	return a.categoryRepository.GetOrCreate(category)
}
