package commands

type GetProductCommand struct {
	Category uint
}

func NewGetProductCommand(category uint) *GetProductCommand {
	return &GetProductCommand{
		Category: category,
	}
}

