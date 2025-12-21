package deleteproduct

import (
	"github.com/mathefer/tc-fiap-product/internal/product/domain/repositories"
	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
)

var (
	_ DeleteProductUseCase = (*DeleteProductUseCaseImpl)(nil)
)

type DeleteProductUseCaseImpl struct {
	productRepository repositories.ProductRepository
}

func NewDeleteProductUseCaseImpl(productRepository repositories.ProductRepository) *DeleteProductUseCaseImpl {
	return &DeleteProductUseCaseImpl{productRepository: productRepository}
}

func (u *DeleteProductUseCaseImpl) Execute(command *commands.DeleteProductCommand) error {
	return u.productRepository.Delete(command.ID)
}

