package commands_test

import (
	"testing"

	"github.com/mathefer/tc-fiap-product/internal/product/usecase/commands"
	"github.com/stretchr/testify/assert"
)

func TestNewAddProductCommand(t *testing.T) {
	// Arrange
	name := "Hamburguer"
	category := 1
	price := 34.99
	description := "Hamburguer com salada"
	imageLink := "https://example.com/image.jpg"

	// Act
	cmd := commands.NewAddProductCommand(name, category, price, description, imageLink)

	// Assert
	assert.NotNil(t, cmd)
	assert.Equal(t, name, cmd.Name)
	assert.Equal(t, category, cmd.Category)
	assert.Equal(t, price, cmd.Price)
	assert.Equal(t, description, cmd.Description)
	assert.Equal(t, imageLink, cmd.ImageLink)
}

func TestNewAddProductCommand_WithEmptyValues(t *testing.T) {
	// Arrange & Act
	cmd := commands.NewAddProductCommand("", 0, 0.0, "", "")

	// Assert
	assert.NotNil(t, cmd)
	assert.Empty(t, cmd.Name)
	assert.Zero(t, cmd.Category)
	assert.Zero(t, cmd.Price)
	assert.Empty(t, cmd.Description)
	assert.Empty(t, cmd.ImageLink)
}

func TestNewGetProductCommand(t *testing.T) {
	// Arrange
	category := uint(1)

	// Act
	cmd := commands.NewGetProductCommand(category)

	// Assert
	assert.NotNil(t, cmd)
	assert.Equal(t, category, cmd.Category)
}

func TestNewGetProductCommand_WithZeroCategory(t *testing.T) {
	// Arrange & Act
	cmd := commands.NewGetProductCommand(0)

	// Assert
	assert.NotNil(t, cmd)
	assert.Zero(t, cmd.Category)
}

func TestNewUpdateProductCommand(t *testing.T) {
	// Arrange
	id := uint(1)
	name := "Hamburguer Atualizado"
	category := 1
	price := 39.99
	description := "Hamburguer com bacon"
	imageLink := "https://example.com/updated.jpg"

	// Act
	cmd := commands.NewUpdateProductCommand(id, name, category, price, description, imageLink)

	// Assert
	assert.NotNil(t, cmd)
	assert.Equal(t, id, cmd.ID)
	assert.Equal(t, name, cmd.Name)
	assert.Equal(t, category, cmd.Category)
	assert.Equal(t, price, cmd.Price)
	assert.Equal(t, description, cmd.Description)
	assert.Equal(t, imageLink, cmd.ImageLink)
}

func TestNewUpdateProductCommand_WithEmptyValues(t *testing.T) {
	// Arrange & Act
	cmd := commands.NewUpdateProductCommand(0, "", 0, 0.0, "", "")

	// Assert
	assert.NotNil(t, cmd)
	assert.Zero(t, cmd.ID)
	assert.Empty(t, cmd.Name)
	assert.Zero(t, cmd.Category)
	assert.Zero(t, cmd.Price)
	assert.Empty(t, cmd.Description)
	assert.Empty(t, cmd.ImageLink)
}

func TestNewDeleteProductCommand(t *testing.T) {
	// Arrange
	id := uint(1)

	// Act
	cmd := commands.NewDeleteProductCommand(id)

	// Assert
	assert.NotNil(t, cmd)
	assert.Equal(t, id, cmd.ID)
}

func TestNewDeleteProductCommand_WithZeroID(t *testing.T) {
	// Arrange & Act
	cmd := commands.NewDeleteProductCommand(0)

	// Assert
	assert.NotNil(t, cmd)
	assert.Zero(t, cmd.ID)
}
