#!/bin/bash
export DB_HOST=${DB_HOST:-localhost}
export DB_PORT=${DB_PORT:-5432}
export DB_USER=${DB_USER:-product_user}
export DB_PASSWORD=${DB_PASSWORD:-product_password}
export DB_NAME=${DB_NAME:-product_db}
export DB_SSLMODE=${DB_SSLMODE:-disable}
export PORT=${PORT:-8081}

echo "Starting Product Service..."
echo "Database: $DB_HOST:$DB_PORT/$DB_NAME"
echo "Port: $PORT"
go run cmd/api/main.go
