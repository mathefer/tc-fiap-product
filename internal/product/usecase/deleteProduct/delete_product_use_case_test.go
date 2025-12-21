package deleteproduct_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
	deleteproduct "github.com/mathefer/tc-fiap-product/internal/product/usecase/deleteProduct"
	mockRepositories "github.com/mathefer/tc-fiap-product/mocks/product/domain/repositories"
)

type DeleteProductUseCaseTestSuite struct {
	suite.Suite
	mockRepository *mockRepositories.MockProductRepository
	useCase        deleteproduct.DeleteProductUseCase
}

func (suite *DeleteProductUseCaseTestSuite) SetupTest() {
	suite.mockRepository = mockRepositories.NewMockProductRepository(suite.T())
	suite.useCase = deleteproduct.NewDeleteProductUseCaseImpl(suite.mockRepository)
}

func TestDeleteProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteProductUseCaseTestSuite))
}

func (suite *DeleteProductUseCaseTestSuite) TestExecute_Success() {
	// Arrange
	id := uint(1)
	command := commands.NewDeleteProductCommand(id)

	suite.mockRepository.EXPECT().
		Delete(id).
		Return(nil).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *DeleteProductUseCaseTestSuite) TestExecute_RepositoryError() {
	// Arrange
	id := uint(1)
	command := commands.NewDeleteProductCommand(id)

	expectedError := errors.New("database error")

	suite.mockRepository.EXPECT().
		Delete(id).
		Return(expectedError).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *DeleteProductUseCaseTestSuite) TestExecute_ProductNotFound() {
	// Arrange
	id := uint(999)
	command := commands.NewDeleteProductCommand(id)

	expectedError := errors.New("product not found")

	suite.mockRepository.EXPECT().
		Delete(id).
		Return(expectedError).
		Once()

	// Act
	err := suite.useCase.Execute(command)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockRepository.AssertExpectations(suite.T())
}

