package getproduct

import (
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
)

type GetProductUseCase interface {
	Execute(command *commands.GetProductCommand) ([]*entities.Product, error)
}
