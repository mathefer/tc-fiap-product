package updateproduct

import "github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"

type UpdateProductUseCase interface {
	Execute(command *commands.UpdateProductCommand) error
}

