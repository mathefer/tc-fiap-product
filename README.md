# TC-FIAP Product Microservice

Product microservice extracted from the tc-fiap-50 monolith.

## Overview

This microservice handles all product-related operations including:
- Get products by category
- Add new products
- Update existing products
- Delete products

## API Endpoints

- `GET /v1/product?category={id}` - Get products by category
- `POST /v1/product` - Add a new product
- `PUT /v1/product/{id}` - Update a product
- `DELETE /v1/product/{id}` - Delete a product

## Category Values

- 1 - Lanche
- 2 - Acompanhamento
- 3 - Bebida
- 4 - Sobremesa

## Environment Variables

- `DB_HOST` - Database host
- `DB_PORT` - Database port (default: 5432)
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name (default: product_db)
- `DB_SSLMODE` - SSL mode (default: disable)
- `PORT` - Application port (default: 8081)

## Running Locally

### Using Docker Compose

```bash
docker-compose up
```

The service will be available at `http://localhost:8081`

### Using Go

```bash
go mod download
go run cmd/api/main.go
```

## Swagger Documentation

Once the service is running, Swagger documentation is available at:
`http://localhost:8081/swagger/index.html`

