package presenter

import (
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/infrastructure/api/dto"
)

var (
	_ ProductPresenter = (*ProductPresenterImpl)(nil)
)

type ProductPresenterImpl struct {
}

func NewProductPresenterImpl() *ProductPresenterImpl {
	return &ProductPresenterImpl{}
}

func (p *ProductPresenterImpl) Present(products []*entities.Product) []*dto.GetProductResponseDto {
	productDto := make([]*dto.GetProductResponseDto, len(products))

	for i, product := range products {
		productDto[i] = &dto.GetProductResponseDto{
			ID:          product.ID,
			CreatedAt:   product.CreatedAt,
			Name:        product.Name,
			Category:    product.Category,
			Price:       product.Price,
			Description: product.Description,
			ImageLink:   product.ImageLink,
		}
	}

	return productDto
}
