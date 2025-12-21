package dto

type AddProductRequestDto struct {
	Name        string  `json:"name" example:"Hamburguer"`
	Category    int     `json:"category" example:"1"`
	Price       float64 `json:"price" example:"34.99"`
	Description string  `json:"description" example:"Hamburguer com salada"`
	ImageLink   string  `json:"image_link" example:"https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"`
}

