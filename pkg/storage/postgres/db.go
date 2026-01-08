package postgres

import (
	"errors"
	"fmt"
	"log"
	"os"

	productEntities "github.com/mathefer/tc-fiap-product/internal/product/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConfig holds database configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// ErrMissingEnvVars is returned when required environment variables are not set
var ErrMissingEnvVars = errors.New("database environment variables are not properly set")

// NewPostgresDB creates a new PostgreSQL database connection.
// It reads configuration from environment variables and panics on failure.
// For production use.
func NewPostgresDB() *gorm.DB {
	dsn, err := NewDBConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	db, err := NewDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

// NewDBConfigFromEnv reads database configuration from environment variables.
// Returns error if any required variable is missing.
func NewDBConfigFromEnv() (string, error) {
	config := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return BuildDSN(config)
}

// BuildDSN constructs a PostgreSQL DSN from the given config.
// Returns error if any required field is empty.
func BuildDSN(config DBConfig) (string, error) {
	if config.Host == "" || config.User == "" || config.Password == "" ||
		config.DBName == "" || config.Port == "" || config.SSLMode == "" {
		return "", ErrMissingEnvVars
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode), nil
}

// NewDB creates a new GORM database connection with the given DSN.
// Returns error if connection fails.
func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

// Migrate runs database migrations for all entities.
// Returns error if migration fails.
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&productEntities.Product{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
