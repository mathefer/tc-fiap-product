package repositories

import "github.com/mathefer/tc-fiap-product/internal/product/domain/entities"

type ProductRepository interface {
	Get(category uint) ([]*entities.Product, error)
	Add(product *entities.Product) error
	Update(product *entities.Product) error
	Delete(id uint) error
}
