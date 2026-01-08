package dto_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mathefer/tc-fiap-product/internal/product/infrastructure/api/dto"
	"github.com/stretchr/testify/assert"
)

func TestAddProductRequestDto_JSONMarshal(t *testing.T) {
	// Arrange
	request := dto.AddProductRequestDto{
		Name:        "Hamburguer",
		Category:    1,
		Price:       34.99,
		Description: "Hamburguer com salada",
		ImageLink:   "https://example.com/image.jpg",
	}

	// Act
	data, err := json.Marshal(request)

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"name":"Hamburguer"`)
	assert.Contains(t, string(data), `"category":1`)
	assert.Contains(t, string(data), `"price":34.99`)
	assert.Contains(t, string(data), `"description":"Hamburguer com salada"`)
	assert.Contains(t, string(data), `"image_link":"https://example.com/image.jpg"`)
}

func TestAddProductRequestDto_JSONUnmarshal(t *testing.T) {
	// Arrange
	jsonData := `{"name":"Pizza","category":1,"price":45.99,"description":"Pizza margherita","image_link":"https://example.com/pizza.jpg"}`

	// Act
	var request dto.AddProductRequestDto
	err := json.Unmarshal([]byte(jsonData), &request)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Pizza", request.Name)
	assert.Equal(t, 1, request.Category)
	assert.Equal(t, 45.99, request.Price)
	assert.Equal(t, "Pizza margherita", request.Description)
	assert.Equal(t, "https://example.com/pizza.jpg", request.ImageLink)
}

func TestAddProductRequestDto_JSONUnmarshal_EmptyFields(t *testing.T) {
	// Arrange
	jsonData := `{}`

	// Act
	var request dto.AddProductRequestDto
	err := json.Unmarshal([]byte(jsonData), &request)

	// Assert
	assert.NoError(t, err)
	assert.Empty(t, request.Name)
	assert.Zero(t, request.Category)
	assert.Zero(t, request.Price)
	assert.Empty(t, request.Description)
	assert.Empty(t, request.ImageLink)
}

func TestUpdateProductRequestDto_JSONMarshal(t *testing.T) {
	// Arrange
	request := dto.UpdateProductRequestDto{
		Name:        "Hamburguer Atualizado",
		Category:    1,
		Price:       39.99,
		Description: "Hamburguer com bacon",
		ImageLink:   "https://example.com/updated.jpg",
	}

	// Act
	data, err := json.Marshal(request)

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"name":"Hamburguer Atualizado"`)
	assert.Contains(t, string(data), `"category":1`)
	assert.Contains(t, string(data), `"price":39.99`)
	assert.Contains(t, string(data), `"description":"Hamburguer com bacon"`)
	assert.Contains(t, string(data), `"image_link":"https://example.com/updated.jpg"`)
}

func TestUpdateProductRequestDto_JSONUnmarshal(t *testing.T) {
	// Arrange
	jsonData := `{"name":"Hamburguer Especial","category":1,"price":49.99,"description":"Hamburguer duplo","image_link":"https://example.com/special.jpg"}`

	// Act
	var request dto.UpdateProductRequestDto
	err := json.Unmarshal([]byte(jsonData), &request)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Hamburguer Especial", request.Name)
	assert.Equal(t, 1, request.Category)
	assert.Equal(t, 49.99, request.Price)
	assert.Equal(t, "Hamburguer duplo", request.Description)
	assert.Equal(t, "https://example.com/special.jpg", request.ImageLink)
}

func TestUpdateProductRequestDto_JSONUnmarshal_PartialFields(t *testing.T) {
	// Arrange
	jsonData := `{"name":"Hamburguer","price":34.99}`

	// Act
	var request dto.UpdateProductRequestDto
	err := json.Unmarshal([]byte(jsonData), &request)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Hamburguer", request.Name)
	assert.Zero(t, request.Category)
	assert.Equal(t, 34.99, request.Price)
	assert.Empty(t, request.Description)
	assert.Empty(t, request.ImageLink)
}

func TestGetProductResponseDto_JSONMarshal(t *testing.T) {
	// Arrange
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	response := dto.GetProductResponseDto{
		ID:          1,
		CreatedAt:   createdAt,
		Name:        "Hamburguer",
		Category:    1,
		Price:       34.99,
		Description: "Hamburguer com salada",
		ImageLink:   "https://example.com/image.jpg",
	}

	// Act
	data, err := json.Marshal(response)

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":1`)
	assert.Contains(t, string(data), `"created_at":`)
	assert.Contains(t, string(data), `"name":"Hamburguer"`)
	assert.Contains(t, string(data), `"category":1`)
	assert.Contains(t, string(data), `"price":34.99`)
	assert.Contains(t, string(data), `"description":"Hamburguer com salada"`)
	assert.Contains(t, string(data), `"image_link":"https://example.com/image.jpg"`)
}

func TestGetProductResponseDto_JSONUnmarshal(t *testing.T) {
	// Arrange
	jsonData := `{"id":1,"created_at":"2024-01-15T10:30:00Z","name":"Hamburguer","category":1,"price":34.99,"description":"Hamburguer com salada","image_link":"https://example.com/image.jpg"}`

	// Act
	var response dto.GetProductResponseDto
	err := json.Unmarshal([]byte(jsonData), &response)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, 2024, response.CreatedAt.Year())
	assert.Equal(t, time.January, response.CreatedAt.Month())
	assert.Equal(t, 15, response.CreatedAt.Day())
	assert.Equal(t, "Hamburguer", response.Name)
	assert.Equal(t, 1, response.Category)
	assert.Equal(t, 34.99, response.Price)
	assert.Equal(t, "Hamburguer com salada", response.Description)
	assert.Equal(t, "https://example.com/image.jpg", response.ImageLink)
}

func TestGetProductResponseDto_JSONMarshal_ZeroValues(t *testing.T) {
	// Arrange
	response := dto.GetProductResponseDto{}

	// Act
	data, err := json.Marshal(response)

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":0`)
	assert.Contains(t, string(data), `"name":""`)
	assert.Contains(t, string(data), `"category":0`)
	assert.Contains(t, string(data), `"price":0`)
}

func TestAddProductRequestDto_JSONMarshal_SpecialCharacters(t *testing.T) {
	// Arrange
	request := dto.AddProductRequestDto{
		Name:        `Hamburguer "Especial"`,
		Category:    1,
		Price:       34.99,
		Description: "Hamburguer com\nquebra de linha",
		ImageLink:   "https://example.com/image.jpg?param=value&other=123",
	}

	// Act
	data, err := json.Marshal(request)

	// Assert
	assert.NoError(t, err)

	// Unmarshal back to verify round-trip
	var decoded dto.AddProductRequestDto
	err = json.Unmarshal(data, &decoded)
	assert.NoError(t, err)
	assert.Equal(t, request.Name, decoded.Name)
	assert.Equal(t, request.Description, decoded.Description)
	assert.Equal(t, request.ImageLink, decoded.ImageLink)
}
