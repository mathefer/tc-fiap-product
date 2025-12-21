package addproduct_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	addproduct "github.com/mathefer/tc-fiap-product/internal/product/usecase/addProduct"
	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
	mockRepositories "github.com/mathefer/tc-fiap-product/mocks/product/domain/repositories"
)

type AddProductUseCaseTestSuite struct {
	suite.Suite
	mockRepository *mockRepositories.MockProductRepository
	useCase        addproduct.AddProductUseCase
}

func (suite *AddProductUseCaseTestSuite) SetupTest() {
	suite.mockRepository = mockRepositories.NewMockProductRepository(suite.T())
	suite.useCase = addproduct.NewAddProductUseCaseImpl(suite.mockRepository)
}

func TestAddProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AddProductUseCaseTestSuite))
}

func (suite *AddProductUseCaseTestSuite) TestExecute_Success() {
	// Arrange
	command := commands.NewAddProductCommand("Hamburguer", 1, 34.99, "Hamburguer com salada", "https://example.com/image.jpg")

	expectedProduct := &entities.Product{
		Name:        command.Name,
		Category:    command.Category,
		Price:       command.Price,
		Description: command.Description,
		ImageLink:   command.ImageLink,
	}

	suite.mockRepository.EXPECT().
		Add(expectedProduct).
		Return(nil).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *AddProductUseCaseTestSuite) TestExecute_RepositoryError() {
	// Arrange
	command := commands.NewAddProductCommand("Pizza", 1, 45.99, "Pizza margherita", "https://example.com/pizza.jpg")

	expectedProduct := &entities.Product{
		Name:        command.Name,
		Category:    command.Category,
		Price:       command.Price,
		Description: command.Description,
		ImageLink:   command.ImageLink,
	}

	expectedError := errors.New("database error")

	suite.mockRepository.EXPECT().
		Add(expectedProduct).
		Return(expectedError).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *AddProductUseCaseTestSuite) TestExecute_ValidatesProductData() {
	// Arrange
	command := commands.NewAddProductCommand("", 0, 0.0, "", "")

	expectedProduct := &entities.Product{
		Name:        command.Name,
		Category:    command.Category,
		Price:       command.Price,
		Description: command.Description,
		ImageLink:   command.ImageLink,
	}

	suite.mockRepository.EXPECT().
		Add(expectedProduct).
		Return(nil).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockRepository.AssertExpectations(suite.T())
}

