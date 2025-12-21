package deleteproduct

import "github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"

type DeleteProductUseCase interface {
	Execute(command *commands.DeleteProductCommand) error
}

