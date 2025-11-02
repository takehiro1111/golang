package repository

import (
	"gorm.io/gorm"

	"go-api-arch-onion-template/application/domain"
)

type CategoryRepository interface {
	GetOrCreate(category *domain.Category) (*domain.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) GetOrCreate(category *domain.Category) (*domain.Category, error) {
	var getOrCreatedCategory domain.Category
	tx := c.db.FirstOrCreate(&getOrCreatedCategory, category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &getOrCreatedCategory, nil
}
