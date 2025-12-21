package updateproduct

import (
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/domain/repositories"
	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
)

var (
	_ UpdateProductUseCase = (*UpdateProductUseCaseImpl)(nil)
)

type UpdateProductUseCaseImpl struct {
	productRepository repositories.ProductRepository
}

func NewUpdateProductUseCaseImpl(productRepository repositories.ProductRepository) *UpdateProductUseCaseImpl {
	return &UpdateProductUseCaseImpl{productRepository: productRepository}
}

func (u *UpdateProductUseCaseImpl) Execute(command *commands.UpdateProductCommand) error {
	entity := entities.Product{
		ID:          command.ID,
		Name:        command.Name,
		Category:    command.Category,
		Price:       command.Price,
		Description: command.Description,
		ImageLink:   command.ImageLink,
	}

	return u.productRepository.Update(&entity)
}

