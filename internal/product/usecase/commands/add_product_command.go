package commands

type AddProductCommand struct {
	Name        string
	Category    int
	Price       float64
	Description string
	ImageLink   string
}

func NewAddProductCommand(name string, category int, price float64, description string, imageLink string) *AddProductCommand {
	return &AddProductCommand{
		Name:        name,
		Category:    category,
		Price:       price,
		Description: description,
		ImageLink:   imageLink,
	}
}

