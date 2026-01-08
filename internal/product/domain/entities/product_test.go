package entities_test

import (
	"testing"
	"time"

	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestProduct_TableName(t *testing.T) {
	// Arrange
	product := entities.Product{}

	// Act
	tableName := product.TableName()

	// Assert
	assert.Equal(t, "product", tableName)
}

func TestProduct_StructFields(t *testing.T) {
	// Arrange
	createdAt := time.Now()
	product := entities.Product{
		ID:          1,
		CreatedAt:   createdAt,
		Name:        "Hamburguer",
		Category:    1,
		Price:       34.99,
		Description: "Hamburguer com salada",
		ImageLink:   "https://example.com/image.jpg",
	}

	// Assert
	assert.Equal(t, uint(1), product.ID)
	assert.Equal(t, createdAt, product.CreatedAt)
	assert.Equal(t, "Hamburguer", product.Name)
	assert.Equal(t, 1, product.Category)
	assert.Equal(t, 34.99, product.Price)
	assert.Equal(t, "Hamburguer com salada", product.Description)
	assert.Equal(t, "https://example.com/image.jpg", product.ImageLink)
}

func TestProduct_ZeroValues(t *testing.T) {
	// Arrange
	product := entities.Product{}

	// Assert
	assert.Zero(t, product.ID)
	assert.True(t, product.CreatedAt.IsZero())
	assert.Empty(t, product.Name)
	assert.Zero(t, product.Category)
	assert.Zero(t, product.Price)
	assert.Empty(t, product.Description)
	assert.Empty(t, product.ImageLink)
}
