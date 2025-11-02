package repository_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go-api-arch-onion-template/application/domain"
	"go-api-arch-onion-template/infrastructure/database/repository"
	"go-api-arch-onion-template/pkg/tester"
)

type CategoryRepositorySuite struct {
	tester.DBSQLiteSuite
	repository repository.CategoryRepository
}

func TestCategoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(CategoryRepositorySuite))
}

func (suite *CategoryRepositorySuite) SetupSuite() {
	suite.DBSQLiteSuite.SetupSuite()
	suite.repository = repository.NewCategoryRepository(suite.DBSQLiteSuite.DB)
}

func (suite *CategoryRepositorySuite) TestCategory() {
	paramCategory, err := domain.NewCategory("sports")
	suite.Assert().Nil(err)
	category, err := suite.repository.GetOrCreate(paramCategory)
	suite.Assert().Nil(err)
	suite.Assert().Equal(1, category.ID)
	suite.Assert().Equal("sports", string(category.Name))

	category, err = suite.repository.GetOrCreate(paramCategory)
	suite.Assert().Nil(err)
	suite.Assert().Equal(1, category.ID)
	suite.Assert().Equal("sports", string(category.Name))
}
