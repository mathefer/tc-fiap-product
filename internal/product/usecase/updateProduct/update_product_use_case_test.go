package updateproduct_test

import (
	"errors"
	"testing"

	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
	updateproduct "github.com/mathefer/tc-fiap-product/internal/product/usecase/updateProduct"
	mockRepositories "github.com/mathefer/tc-fiap-product/mocks/product/domain/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UpdateProductUseCaseTestSuite struct {
	suite.Suite
	mockRepository *mockRepositories.MockProductRepository
	useCase        updateproduct.UpdateProductUseCase
}

func (suite *UpdateProductUseCaseTestSuite) SetupTest() {
	suite.mockRepository = mockRepositories.NewMockProductRepository(suite.T())
	suite.useCase = updateproduct.NewUpdateProductUseCaseImpl(suite.mockRepository)
}

func TestUpdateProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateProductUseCaseTestSuite))
}

func (suite *UpdateProductUseCaseTestSuite) TestExecute_Success() {
	// Arrange
	command := commands.NewUpdateProductCommand(1, "Hamburguer Atualizado", 1, 39.99, "Hamburguer com bacon", "https://example.com/updated.jpg")

	expectedProduct := &entities.Product{
		ID:          command.ID,
		Name:        command.Name,
		Category:    command.Category,
		Price:       command.Price,
		Description: command.Description,
		ImageLink:   command.ImageLink,
	}

	suite.mockRepository.EXPECT().
		Update(expectedProduct).
		Return(nil).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *UpdateProductUseCaseTestSuite) TestExecute_RepositoryError() {
	// Arrange
	command := commands.NewUpdateProductCommand(1, "Pizza", 1, 45.99, "Pizza margherita", "https://example.com/pizza.jpg")

	expectedProduct := &entities.Product{
		ID:          command.ID,
		Name:        command.Name,
		Category:    command.Category,
		Price:       command.Price,
		Description: command.Description,
		ImageLink:   command.ImageLink,
	}

	expectedError := errors.New("database error")

	suite.mockRepository.EXPECT().
		Update(expectedProduct).
		Return(expectedError).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *UpdateProductUseCaseTestSuite) TestExecute_ProductNotFound() {
	// Arrange
	command := commands.NewUpdateProductCommand(999, "Non-existent Product", 1, 10.0, "Description", "https://example.com/image.jpg")

	expectedProduct := &entities.Product{
		ID:          command.ID,
		Name:        command.Name,
		Category:    command.Category,
		Price:       command.Price,
		Description: command.Description,
		ImageLink:   command.ImageLink,
	}

	expectedError := errors.New("product not found")

	suite.mockRepository.EXPECT().
		Update(expectedProduct).
		Return(expectedError).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockRepository.AssertExpectations(suite.T())
}
