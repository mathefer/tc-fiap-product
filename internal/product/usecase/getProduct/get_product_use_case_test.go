package getproduct_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
	getproduct "github.com/mathefer/tc-fiap-product/internal/product/usecase/getProduct"
	mockRepositories "github.com/mathefer/tc-fiap-product/mocks/product/domain/repositories"
)

type GetProductUseCaseTestSuite struct {
	suite.Suite
	mockRepository *mockRepositories.MockProductRepository
	useCase        getproduct.GetProductUseCase
}

func (suite *GetProductUseCaseTestSuite) SetupTest() {
	suite.mockRepository = mockRepositories.NewMockProductRepository(suite.T())
	suite.useCase = getproduct.NewGetProductUseCaseImpl(suite.mockRepository)
}

func TestGetProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GetProductUseCaseTestSuite))
}

func (suite *GetProductUseCaseTestSuite) TestExecute_Success() {
	// Arrange
	category := uint(1)
	command := commands.NewGetProductCommand(category)

	expectedProducts := []*entities.Product{
		{
			ID:          1,
			CreatedAt:   time.Now(),
			Name:        "Hamburguer",
			Category:    1,
			Price:       34.99,
			Description: "Hamburguer com salada",
			ImageLink:   "https://example.com/image.jpg",
		},
		{
			ID:          2,
			CreatedAt:   time.Now(),
			Name:        "Cheeseburguer",
			Category:    1,
			Price:       39.99,
			Description: "Hamburguer com queijo",
			ImageLink:   "https://example.com/cheese.jpg",
		},
	}

	suite.mockRepository.EXPECT().
		Get(category).
		Return(expectedProducts, nil).
		Once()

	// Act
	products, err := suite.useCase.Execute(command)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), products)
	assert.Len(suite.T(), products, 2)
	assert.Equal(suite.T(), expectedProducts[0].ID, products[0].ID)
	assert.Equal(suite.T(), expectedProducts[0].Name, products[0].Name)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *GetProductUseCaseTestSuite) TestExecute_EmptyResult() {
	// Arrange
	category := uint(2)
	command := commands.NewGetProductCommand(category)

	expectedProducts := []*entities.Product{}

	suite.mockRepository.EXPECT().
		Get(category).
		Return(expectedProducts, nil).
		Once()

	// Act
	products, err := suite.useCase.Execute(command)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), products)
	assert.Len(suite.T(), products, 0)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *GetProductUseCaseTestSuite) TestExecute_RepositoryError() {
	// Arrange
	category := uint(1)
	command := commands.NewGetProductCommand(category)

	expectedError := errors.New("database connection error")

	suite.mockRepository.EXPECT().
		Get(category).
		Return(nil, expectedError).
		Once()

	// Act
	products, err := suite.useCase.Execute(command)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), products)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockRepository.AssertExpectations(suite.T())
}

