package getproduct

import (
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/domain/repositories"
	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
)

var (
	_ GetProductUseCase = (*GetProductUseCaseImpl)(nil)
)

type GetProductUseCaseImpl struct {
	productRepository repositories.ProductRepository
}

func NewGetProductUseCaseImpl(productRepository repositories.ProductRepository) *GetProductUseCaseImpl {
	return &GetProductUseCaseImpl{productRepository: productRepository}
}

func (u *GetProductUseCaseImpl) Execute(command *commands.GetProductCommand) ([]*entities.Product, error) {
	entities, err := u.productRepository.Get(command.Category)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

