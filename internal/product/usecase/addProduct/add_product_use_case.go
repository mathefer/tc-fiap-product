package addproduct

import "github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"

type AddProductUseCase interface {
	Execute(command *commands.AddProductCommand) error
}
