package controller_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/mathefer/tc-fiap-product/internal/product/controller"
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/infrastructure/api/dto"
	mockPresenter "github.com/mathefer/tc-fiap-product/mocks/product/presenter"
	mockAddProduct "github.com/mathefer/tc-fiap-product/mocks/product/usecase/addProduct"
	mockDeleteProduct "github.com/mathefer/tc-fiap-product/mocks/product/usecase/deleteProduct"
	mockGetProduct "github.com/mathefer/tc-fiap-product/mocks/product/usecase/getProduct"
	mockUpdateProduct "github.com/mathefer/tc-fiap-product/mocks/product/usecase/updateProduct"
)

type ProductControllerTestSuite struct {
	suite.Suite
	mockPresenter          *mockPresenter.MockProductPresenter
	mockAddProductUseCase  *mockAddProduct.MockAddProductUseCase
	mockGetProductUseCase  *mockGetProduct.MockGetProductUseCase
	mockUpdateProductUseCase *mockUpdateProduct.MockUpdateProductUseCase
	mockDeleteProductUseCase *mockDeleteProduct.MockDeleteProductUseCase
	productController      controller.ProductController
}

func (suite *ProductControllerTestSuite) SetupTest() {
	suite.mockPresenter = mockPresenter.NewMockProductPresenter(suite.T())
	suite.mockAddProductUseCase = mockAddProduct.NewMockAddProductUseCase(suite.T())
	suite.mockGetProductUseCase = mockGetProduct.NewMockGetProductUseCase(suite.T())
	suite.mockUpdateProductUseCase = mockUpdateProduct.NewMockUpdateProductUseCase(suite.T())
	suite.mockDeleteProductUseCase = mockDeleteProduct.NewMockDeleteProductUseCase(suite.T())

	suite.productController = controller.NewProductControllerImpl(
		suite.mockPresenter,
		suite.mockAddProductUseCase,
		suite.mockGetProductUseCase,
		suite.mockUpdateProductUseCase,
		suite.mockDeleteProductUseCase,
	)
}

func TestProductControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductControllerTestSuite))
}

func (suite *ProductControllerTestSuite) TestGet_Success() {
	// Arrange
	category := uint(1)
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
	}

	expectedDto := []*dto.GetProductResponseDto{
		{
			ID:          1,
			CreatedAt:   now,
			Name:        "Hamburguer",
			Category:    1,
			Price:       34.99,
			Description: "Hamburguer com salada",
			ImageLink:   "https://example.com/image.jpg",
		},
	}

	suite.mockGetProductUseCase.EXPECT().
		Execute(mock.Anything).
		Return(products, nil).
		Once()

	suite.mockPresenter.EXPECT().
		Present(products).
		Return(expectedDto).
		Once()

	// Act
	result, err := suite.productController.Get(category)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Len(suite.T(), result, 1)
	assert.Equal(suite.T(), expectedDto[0].ID, result[0].ID)
	assert.Equal(suite.T(), expectedDto[0].Name, result[0].Name)
	suite.mockGetProductUseCase.AssertExpectations(suite.T())
	suite.mockPresenter.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestGet_UseCaseError() {
	// Arrange
	category := uint(1)
	expectedError := errors.New("database error")

	suite.mockGetProductUseCase.EXPECT().
		Execute(mock.Anything).
		Return(nil, expectedError).
		Once()

	// Act
	result, err := suite.productController.Get(category)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockGetProductUseCase.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestAdd_Success() {
	// Arrange
	requestDto := &dto.AddProductRequestDto{
		Name:        "Pizza",
		Category:    1,
		Price:       45.99,
		Description: "Pizza margherita",
		ImageLink:   "https://example.com/pizza.jpg",
	}

	suite.mockAddProductUseCase.EXPECT().
		Execute(mock.Anything).
		Return(nil).
		Once()

	// Act
	err := suite.productController.Add(requestDto)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockAddProductUseCase.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestAdd_UseCaseError() {
	// Arrange
	requestDto := &dto.AddProductRequestDto{
		Name:        "Hamburguer",
		Category:    1,
		Price:       34.99,
		Description: "Hamburguer com salada",
		ImageLink:   "https://example.com/image.jpg",
	}

	expectedError := errors.New("database error")

	suite.mockAddProductUseCase.EXPECT().
		Execute(mock.Anything).
		Return(expectedError).
		Once()

	// Act
	err := suite.productController.Add(requestDto)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockAddProductUseCase.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestUpdate_Success() {
	// Arrange
	id := uint(1)
	requestDto := &dto.UpdateProductRequestDto{
		Name:        "Hamburguer Atualizado",
		Category:    1,
		Price:       39.99,
		Description: "Hamburguer com bacon",
		ImageLink:   "https://example.com/updated.jpg",
	}

	suite.mockUpdateProductUseCase.EXPECT().
		Execute(mock.Anything).
		Return(nil).
		Once()

	// Act
	err := suite.productController.Update(id, requestDto)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockUpdateProductUseCase.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestUpdate_UseCaseError() {
	// Arrange
	id := uint(1)
	requestDto := &dto.UpdateProductRequestDto{
		Name:        "Pizza",
		Category:    1,
		Price:       45.99,
		Description: "Pizza margherita",
		ImageLink:   "https://example.com/pizza.jpg",
	}

	expectedError := errors.New("product not found")

	suite.mockUpdateProductUseCase.EXPECT().
		Execute(mock.Anything).
		Return(expectedError).
		Once()

	// Act
	err := suite.productController.Update(id, requestDto)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockUpdateProductUseCase.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestDelete_Success() {
	// Arrange
	id := uint(1)

	suite.mockDeleteProductUseCase.EXPECT().
		Execute(mock.Anything).
		Return(nil).
		Once()

	// Act
	err := suite.productController.Delete(id)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockDeleteProductUseCase.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestDelete_UseCaseError() {
	// Arrange
	id := uint(1)
	expectedError := errors.New("database error")

	suite.mockDeleteProductUseCase.EXPECT().
		Execute(mock.Anything).
		Return(expectedError).
		Once()

	// Act
	err := suite.productController.Delete(id)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockDeleteProductUseCase.AssertExpectations(suite.T())
}

