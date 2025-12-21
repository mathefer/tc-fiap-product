package controller

import (
	"github.com/mathefer/tc-fiap-product/internal/product/infrastructure/api/dto"
)

type ProductController interface {
	Get(category uint) ([]*dto.GetProductResponseDto, error)
	Add(product *dto.AddProductRequestDto) error
	Update(id uint, product *dto.UpdateProductRequestDto) error
	Delete(id uint) error
}
