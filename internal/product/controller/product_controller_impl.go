package controller

import (
	"github.com/mathefer/tc-fiap-product/internal/product/infrastructure/api/dto"
	productPresenter "github.com/mathefer/tc-fiap-product/internal/product/presenter"
	addProduct "github.com/mathefer/tc-fiap-product/internal/product/usecase/addProduct"
	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
	deleteProduct "github.com/mathefer/tc-fiap-product/internal/product/usecase/deleteProduct"
	getProduct "github.com/mathefer/tc-fiap-product/internal/product/usecase/getProduct"
	updateProduct "github.com/mathefer/tc-fiap-product/internal/product/usecase/updateProduct"
)

var (
	_ ProductController = (*ProductControllerImpl)(nil)
)

type ProductControllerImpl struct {
	presenter            productPresenter.ProductPresenter
	addProductUseCase    addProduct.AddProductUseCase
	getProductUseCase    getProduct.GetProductUseCase
	updateProductUseCase updateProduct.UpdateProductUseCase
	deleteProductUseCase deleteProduct.DeleteProductUseCase
}

func NewProductControllerImpl(
	presenter productPresenter.ProductPresenter,
	addProductUseCase addProduct.AddProductUseCase,
	getProductUseCase getProduct.GetProductUseCase,
	updateProductUseCase updateProduct.UpdateProductUseCase,
	deleteProductUseCase deleteProduct.DeleteProductUseCase) *ProductControllerImpl {
	return &ProductControllerImpl{
		presenter:            presenter,
		addProductUseCase:    addProductUseCase,
		getProductUseCase:    getProductUseCase,
		updateProductUseCase: updateProductUseCase,
		deleteProductUseCase: deleteProductUseCase,
	}
}

func (p *ProductControllerImpl) Get(category uint) ([]*dto.GetProductResponseDto, error) {
	products, err := p.getProductUseCase.Execute(commands.NewGetProductCommand(category))
	if err != nil {
		return nil, err
	}

	return p.presenter.Present(products), nil
}

func (p *ProductControllerImpl) Add(product *dto.AddProductRequestDto) error {
	command := commands.NewAddProductCommand(product.Name, product.Category, product.Price, product.Description, product.ImageLink)
	err := p.addProductUseCase.Execute(command)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductControllerImpl) Update(id uint, product *dto.UpdateProductRequestDto) error {
	command := commands.NewUpdateProductCommand(id, product.Name, product.Category, product.Price, product.Description, product.ImageLink)
	err := p.updateProductUseCase.Execute(command)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductControllerImpl) Delete(id uint) error {
	command := commands.NewDeleteProductCommand(id)
	err := p.deleteProductUseCase.Execute(command)
	if err != nil {
		return err
	}
	return nil
}
