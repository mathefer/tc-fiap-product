package postgres_test

import (
	"os"
	"testing"

	"github.com/mathefer/tc-fiap-product/pkg/storage/postgres"
	"github.com/stretchr/testify/assert"
)

func TestBuildDSN_Success(t *testing.T) {
	// Arrange
	config := postgres.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "testuser",
		Password: "testpass",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	// Act
	dsn, err := postgres.BuildDSN(config)

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, dsn, "host=localhost")
	assert.Contains(t, dsn, "port=5432")
	assert.Contains(t, dsn, "user=testuser")
	assert.Contains(t, dsn, "password=testpass")
	assert.Contains(t, dsn, "dbname=testdb")
	assert.Contains(t, dsn, "sslmode=disable")
}

func TestBuildDSN_MissingHost(t *testing.T) {
	// Arrange
	config := postgres.DBConfig{
		Host:     "",
		Port:     "5432",
		User:     "testuser",
		Password: "testpass",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	// Act
	dsn, err := postgres.BuildDSN(config)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, postgres.ErrMissingEnvVars)
	assert.Empty(t, dsn)
}

func TestBuildDSN_MissingPort(t *testing.T) {
	// Arrange
	config := postgres.DBConfig{
		Host:     "localhost",
		Port:     "",
		User:     "testuser",
		Password: "testpass",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	// Act
	dsn, err := postgres.BuildDSN(config)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, postgres.ErrMissingEnvVars)
	assert.Empty(t, dsn)
}

func TestBuildDSN_MissingUser(t *testing.T) {
	// Arrange
	config := postgres.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "",
		Password: "testpass",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	// Act
	dsn, err := postgres.BuildDSN(config)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, postgres.ErrMissingEnvVars)
	assert.Empty(t, dsn)
}

func TestBuildDSN_MissingPassword(t *testing.T) {
	// Arrange
	config := postgres.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "testuser",
		Password: "",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	// Act
	dsn, err := postgres.BuildDSN(config)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, postgres.ErrMissingEnvVars)
	assert.Empty(t, dsn)
}

func TestBuildDSN_MissingDBName(t *testing.T) {
	// Arrange
	config := postgres.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "testuser",
		Password: "testpass",
		DBName:   "",
		SSLMode:  "disable",
	}

	// Act
	dsn, err := postgres.BuildDSN(config)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, postgres.ErrMissingEnvVars)
	assert.Empty(t, dsn)
}

func TestBuildDSN_MissingSSLMode(t *testing.T) {
	// Arrange
	config := postgres.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "testuser",
		Password: "testpass",
		DBName:   "testdb",
		SSLMode:  "",
	}

	// Act
	dsn, err := postgres.BuildDSN(config)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, postgres.ErrMissingEnvVars)
	assert.Empty(t, dsn)
}

func TestBuildDSN_AllFieldsEmpty(t *testing.T) {
	// Arrange
	config := postgres.DBConfig{}

	// Act
	dsn, err := postgres.BuildDSN(config)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, postgres.ErrMissingEnvVars)
	assert.Empty(t, dsn)
}

func TestNewDBConfigFromEnv_Success(t *testing.T) {
	// Arrange - Set environment variables
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpass")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_SSLMODE", "disable")
	defer func() {
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_PASSWORD")
		os.Unsetenv("DB_NAME")
		os.Unsetenv("DB_SSLMODE")
	}()

	// Act
	dsn, err := postgres.NewDBConfigFromEnv()

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, dsn, "host=localhost")
	assert.Contains(t, dsn, "port=5432")
	assert.Contains(t, dsn, "user=testuser")
	assert.Contains(t, dsn, "password=testpass")
	assert.Contains(t, dsn, "dbname=testdb")
	assert.Contains(t, dsn, "sslmode=disable")
}

func TestNewDBConfigFromEnv_MissingVariables(t *testing.T) {
	// Arrange - Ensure environment variables are not set
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_SSLMODE")

	// Act
	dsn, err := postgres.NewDBConfigFromEnv()

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, postgres.ErrMissingEnvVars)
	assert.Empty(t, dsn)
}

func TestNewDB_InvalidDSN(t *testing.T) {
	// Arrange - Use an invalid DSN that will fail to connect
	invalidDSN := "host=invalid_host_that_does_not_exist port=9999 user=fake password=fake dbname=fake sslmode=disable"

	// Act
	db, err := postgres.NewDB(invalidDSN)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, db)
	assert.Contains(t, err.Error(), "failed to connect to database")
}

func TestDBConfig_StructFields(t *testing.T) {
	// Arrange
	config := postgres.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "testuser",
		Password: "testpass",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	// Assert
	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, "5432", config.Port)
	assert.Equal(t, "testuser", config.User)
	assert.Equal(t, "testpass", config.Password)
	assert.Equal(t, "testdb", config.DBName)
	assert.Equal(t, "disable", config.SSLMode)
}
