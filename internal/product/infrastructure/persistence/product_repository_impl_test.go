package persistence_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"github.com/mathefer/tc-fiap-product/internal/product/infrastructure/persistence"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProductRepositoryTestSuite struct {
	suite.Suite
	mockDB     sqlmock.Sqlmock
	db         *gorm.DB
	repository *persistence.ProductRepositoryImpl
}

func (suite *ProductRepositoryTestSuite) SetupTest() {
	var err error
	var sqlDB *sql.DB
	sqlDB, suite.mockDB, err = sqlmock.New()
	if err != nil {
		suite.T().Fatalf("Failed to open mock sql db, got error: %v", err)
	}

	if sqlDB == nil {
		suite.T().Fatal("mock db is null")
	}

	suite.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		suite.T().Fatalf("Failed to open gorm db, got error: %v", err)
	}

	suite.repository = persistence.NewProductRepositoryImpl(suite.db)
}

func (suite *ProductRepositoryTestSuite) TearDownTest() {
	suite.db = nil
	suite.mockDB = nil
	suite.repository = nil
}

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}

func (suite *ProductRepositoryTestSuite) TestGet_Success() {
	// Arrange
	category := uint(1)
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "created_at", "name", "category", "price", "description", "image_link"}).
		AddRow(1, now, "Hamburguer", 1, 34.99, "Hamburguer com salada", "https://example.com/image.jpg").
		AddRow(2, now, "Cheeseburguer", 1, 39.99, "Hamburguer com queijo", "https://example.com/cheese.jpg")

	suite.mockDB.ExpectQuery(`SELECT \* FROM "product" WHERE category = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	// Act
	products, err := suite.repository.Get(category)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), products)
	assert.Len(suite.T(), products, 2)
	assert.Equal(suite.T(), uint(1), products[0].ID)
	assert.Equal(suite.T(), "Hamburguer", products[0].Name)
	assert.Equal(suite.T(), 1, products[0].Category)
	assert.Equal(suite.T(), 34.99, products[0].Price)
	suite.mockDB.ExpectationsWereMet()
}

func (suite *ProductRepositoryTestSuite) TestGet_EmptyResult() {
	// Arrange
	category := uint(2)

	rows := sqlmock.NewRows([]string{"id", "created_at", "name", "category", "price", "description", "image_link"})

	suite.mockDB.ExpectQuery(`SELECT \* FROM "product" WHERE category = \$1`).
		WithArgs(2).
		WillReturnRows(rows)

	// Act
	products, err := suite.repository.Get(category)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), products)
	assert.Len(suite.T(), products, 0)
	suite.mockDB.ExpectationsWereMet()
}

func (suite *ProductRepositoryTestSuite) TestGet_DatabaseError() {
	// Arrange
	category := uint(1)
	expectedError := errors.New("database connection error")

	suite.mockDB.ExpectQuery(`SELECT \* FROM "product" WHERE category = \$1`).
		WithArgs(1).
		WillReturnError(expectedError)

	// Act
	products, err := suite.repository.Get(category)

	// Assert
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), products)
	assert.Len(suite.T(), products, 0)
	suite.mockDB.ExpectationsWereMet()
}

func (suite *ProductRepositoryTestSuite) TestAdd_Success() {
	// Arrange
	product := &entities.Product{
		Name:        "Hamburguer",
		Category:    1,
		Price:       34.99,
		Description: "Hamburguer com salada",
		ImageLink:   "https://example.com/image.jpg",
	}

	suite.mockDB.ExpectBegin()
	// GORM doesn't include created_at in INSERT - it's handled by database default
	// The RETURNING clause includes created_at and id
	now := time.Now()
	suite.mockDB.ExpectQuery(`INSERT INTO "product"`).
		WithArgs(product.Name, product.Category, product.Price, product.Description, product.ImageLink).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "id"}).AddRow(now, 1))
	suite.mockDB.ExpectCommit()

	// Act
	err := suite.repository.Add(product)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockDB.ExpectationsWereMet()
}

func (suite *ProductRepositoryTestSuite) TestAdd_DatabaseError() {
	// Arrange
	product := &entities.Product{
		Name:        "Hamburguer",
		Category:    1,
		Price:       34.99,
		Description: "Hamburguer com salada",
		ImageLink:   "https://example.com/image.jpg",
	}

	expectedError := errors.New("database write error")

	suite.mockDB.ExpectBegin()
	// GORM doesn't include created_at in INSERT - it's handled by database default
	suite.mockDB.ExpectQuery(`INSERT INTO "product"`).
		WithArgs(product.Name, product.Category, product.Price, product.Description, product.ImageLink).
		WillReturnError(expectedError)
	suite.mockDB.ExpectRollback()

	// Act
	err := suite.repository.Add(product)

	// Assert
	assert.Error(suite.T(), err)
	suite.mockDB.ExpectationsWereMet()
}

func (suite *ProductRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	product := &entities.Product{
		ID:          1,
		Name:        "Hamburguer Atualizado",
		Category:    1,
		Price:       39.99,
		Description: "Hamburguer com bacon",
		ImageLink:   "https://example.com/updated.jpg",
	}

	suite.mockDB.ExpectBegin()
	// GORM Updates() includes all fields including ID in SET, and ID in WHERE
	suite.mockDB.ExpectExec(`UPDATE "product" SET`).
		WithArgs(sqlmock.AnyArg(), product.Name, product.Category, product.Price, product.Description, product.ImageLink, product.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mockDB.ExpectCommit()

	// Act
	err := suite.repository.Update(product)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockDB.ExpectationsWereMet()
}

func (suite *ProductRepositoryTestSuite) TestUpdate_ProductNotFound() {
	// Arrange
	product := &entities.Product{
		ID:          999,
		Name:        "Non-existent Product",
		Category:    1,
		Price:       10.0,
		Description: "Description",
		ImageLink:   "https://example.com/image.jpg",
	}

	suite.mockDB.ExpectBegin()
	// GORM Updates() includes all fields including ID in SET, and ID in WHERE
	suite.mockDB.ExpectExec(`UPDATE "product" SET`).
		WithArgs(sqlmock.AnyArg(), product.Name, product.Category, product.Price, product.Description, product.ImageLink, product.ID).
		WillReturnResult(sqlmock.NewResult(0, 0))
	suite.mockDB.ExpectCommit()

	// Act
	err := suite.repository.Update(product)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), gorm.ErrRecordNotFound, err)
	suite.mockDB.ExpectationsWereMet()
}

func (suite *ProductRepositoryTestSuite) TestUpdate_DatabaseError() {
	// Arrange
	product := &entities.Product{
		ID:          1,
		Name:        "Hamburguer",
		Category:    1,
		Price:       34.99,
		Description: "Hamburguer com salada",
		ImageLink:   "https://example.com/image.jpg",
	}

	expectedError := errors.New("database update error")

	suite.mockDB.ExpectBegin()
	// GORM Updates() includes all fields including ID in SET, and ID in WHERE
	suite.mockDB.ExpectExec(`UPDATE "product" SET`).
		WithArgs(sqlmock.AnyArg(), product.Name, product.Category, product.Price, product.Description, product.ImageLink, product.ID).
		WillReturnError(expectedError)
	suite.mockDB.ExpectRollback()

	// Act
	err := suite.repository.Update(product)

	// Assert
	assert.Error(suite.T(), err)
	suite.mockDB.ExpectationsWereMet()
}

func (suite *ProductRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	id := uint(1)

	suite.mockDB.ExpectBegin()
	suite.mockDB.ExpectExec(`DELETE FROM "product"`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mockDB.ExpectCommit()

	// Act
	err := suite.repository.Delete(id)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockDB.ExpectationsWereMet()
}

func (suite *ProductRepositoryTestSuite) TestDelete_DatabaseError() {
	// Arrange
	id := uint(1)
	expectedError := errors.New("database delete error")

	suite.mockDB.ExpectBegin()
	suite.mockDB.ExpectExec(`DELETE FROM "product"`).
		WithArgs(id).
		WillReturnError(expectedError)
	suite.mockDB.ExpectRollback()

	// Act
	err := suite.repository.Delete(id)

	// Assert
	assert.Error(suite.T(), err)
	suite.mockDB.ExpectationsWereMet()
}
