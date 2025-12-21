package presenter

import (
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/infrastructure/api/dto"
)

type ProductPresenter interface {
	Present(products []*entities.Product) []*dto.GetProductResponseDto
}
