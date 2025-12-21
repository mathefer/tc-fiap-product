package addproduct

import (
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/domain/repositories"
	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
)

var (
	_ AddProductUseCase = (*AddProductUseCaseImpl)(nil)
)

type AddProductUseCaseImpl struct {
	productRepository repositories.ProductRepository
}

func NewAddProductUseCaseImpl(productRepository repositories.ProductRepository) *AddProductUseCaseImpl {
	return &AddProductUseCaseImpl{productRepository: productRepository}
}

func (u *AddProductUseCaseImpl) Execute(command *commands.AddProductCommand) error {
	entity := entities.Product{
		Name:        command.Name,
		Category:    command.Category,
		Price:       command.Price,
		Description: command.Description,
		ImageLink:   command.ImageLink,
	}

	return u.productRepository.Add(&entity)
}
