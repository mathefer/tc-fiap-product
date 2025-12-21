package commands

type UpdateProductCommand struct {
	ID          uint
	Name        string
	Category    int
	Price       float64
	Description string
	ImageLink   string
}

func NewUpdateProductCommand(id uint, name string, category int, price float64, description string, imageLink string) *UpdateProductCommand {
	return &UpdateProductCommand{
		ID:          id,
		Name:        name,
		Category:    category,
		Price:       price,
		Description: description,
		ImageLink:   imageLink,
	}
}

