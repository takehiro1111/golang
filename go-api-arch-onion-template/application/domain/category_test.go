package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-api-arch-onion-template/application/domain"
)

func TestCategory(t *testing.T) {
	category := domain.Category{
		ID:   1,
		Name: "sports",
	}

	assert.Equal(t, 1, category.ID)
	assert.Equal(t, "sports", string(category.Name))
}
