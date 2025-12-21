package persistence

import (
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/domain/repositories"
	"gorm.io/gorm"
)

var (
	_ repositories.ProductRepository = (*ProductRepositoryImpl)(nil)
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) Get(category uint) ([]*entities.Product, error) {
	var products []*entities.Product
	if err := r.db.Where("category = ?", category).Find(&products).Error; err != nil {
		return []*entities.Product{}, err
	}
	return products, nil
}

func (r *ProductRepositoryImpl) Add(product *entities.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) Update(product *entities.Product) error {
	result := r.db.Model(&entities.Product{}).Where("id = ?", product.ID).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *ProductRepositoryImpl) Delete(id uint) error {
	if err := r.db.Delete(&entities.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
