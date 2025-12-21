package presenter_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/presenter"
)

type ProductPresenterTestSuite struct {
	suite.Suite
	presenter presenter.ProductPresenter
}

func (suite *ProductPresenterTestSuite) SetupTest() {
	suite.presenter = presenter.NewProductPresenterImpl()
}

func TestProductPresenterTestSuite(t *testing.T) {
	suite.Run(t, new(ProductPresenterTestSuite))
}

func (suite *ProductPresenterTestSuite) TestPresent_Success() {
	// Arrange
	now := time.Now()
	products := []*entities.Product{
		{
			ID:          1,
			CreatedAt:   now,
			Name:        "Hamburguer",
			Category:    1,
			Price:       34.99,
			Description: "Hamburguer com salada",
			ImageLink:   "https://example.com/image.jpg",
		},
		{
			ID:          2,
			CreatedAt:   now,
			Name:        "Cheeseburguer",
			Category:    1,
			Price:       39.99,
			Description: "Hamburguer com queijo",
			ImageLink:   "https://example.com/cheese.jpg",
		},
	}

	// Act
	dtos := suite.presenter.Present(products)

	// Assert
	assert.NotNil(suite.T(), dtos)
	assert.Len(suite.T(), dtos, 2)
	assert.Equal(suite.T(), products[0].ID, dtos[0].ID)
	assert.Equal(suite.T(), products[0].Name, dtos[0].Name)
	assert.Equal(suite.T(), products[0].Category, dtos[0].Category)
	assert.Equal(suite.T(), products[0].Price, dtos[0].Price)
	assert.Equal(suite.T(), products[0].Description, dtos[0].Description)
	assert.Equal(suite.T(), products[0].ImageLink, dtos[0].ImageLink)
	assert.Equal(suite.T(), products[0].CreatedAt, dtos[0].CreatedAt)
	assert.Equal(suite.T(), products[1].ID, dtos[1].ID)
	assert.Equal(suite.T(), products[1].Name, dtos[1].Name)
}

func (suite *ProductPresenterTestSuite) TestPresent_EmptyList() {
	// Arrange
	products := []*entities.Product{}

	// Act
	dtos := suite.presenter.Present(products)

	// Assert
	assert.NotNil(suite.T(), dtos)
	assert.Len(suite.T(), dtos, 0)
}

func (suite *ProductPresenterTestSuite) TestPresent_WithEmptyFields() {
	// Arrange
	products := []*entities.Product{
		{
			ID:          0,
			CreatedAt:   time.Time{},
			Name:        "",
			Category:    0,
			Price:       0.0,
			Description: "",
			ImageLink:   "",
		},
	}

	// Act
	dtos := suite.presenter.Present(products)

	// Assert
	assert.NotNil(suite.T(), dtos)
	assert.Len(suite.T(), dtos, 1)
	assert.Equal(suite.T(), uint(0), dtos[0].ID)
	assert.Empty(suite.T(), dtos[0].Name)
	assert.Equal(suite.T(), 0, dtos[0].Category)
	assert.Equal(suite.T(), 0.0, dtos[0].Price)
	assert.Empty(suite.T(), dtos[0].Description)
	assert.Empty(suite.T(), dtos[0].ImageLink)
	assert.True(suite.T(), dtos[0].CreatedAt.IsZero())
}

func (suite *ProductPresenterTestSuite) TestPresent_PreservesAllData() {
	// Arrange
	now := time.Now()
	products := []*entities.Product{
		{
			ID:          123,
			CreatedAt:   now,
			Name:        "Pizza Margherita",
			Category:    1,
			Price:       45.99,
			Description: "Pizza com queijo e manjericão",
			ImageLink:   "https://example.com/pizza.jpg",
		},
	}

	// Act
	dtos := suite.presenter.Present(products)

	// Assert
	assert.NotNil(suite.T(), dtos)
	assert.Len(suite.T(), dtos, 1)
	assert.Equal(suite.T(), uint(123), dtos[0].ID)
	assert.Equal(suite.T(), "Pizza Margherita", dtos[0].Name)
	assert.Equal(suite.T(), 1, dtos[0].Category)
	assert.Equal(suite.T(), 45.99, dtos[0].Price)
	assert.Equal(suite.T(), "Pizza com queijo e manjericão", dtos[0].Description)
	assert.Equal(suite.T(), "https://example.com/pizza.jpg", dtos[0].ImageLink)
	assert.Equal(suite.T(), now, dtos[0].CreatedAt)
}

