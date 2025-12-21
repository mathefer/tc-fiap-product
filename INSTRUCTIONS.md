# Running and Testing the Product Service

## Prerequisites

You need either:
- **Docker Desktop** running, OR
- **PostgreSQL** running locally

## Option 1: Using Docker Compose (Recommended)

1. Make sure Docker Desktop is running
2. Start the services:
   ```bash
   docker-compose up -d
   ```
3. Check logs:
   ```bash
   docker-compose logs -f app
   ```
4. The service will be available at `http://localhost:8081`

## Option 2: Using Local PostgreSQL

1. Make sure PostgreSQL is running locally
2. Create the database:
   ```bash
   createdb product_db
   # Or using psql:
   # psql -c "CREATE DATABASE product_db;"
   ```
3. Set environment variables:
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=your_postgres_user
   export DB_PASSWORD=your_postgres_password
   export DB_NAME=product_db
   export DB_SSLMODE=disable
   export PORT=8081
   ```
4. Run the service:
   ```bash
   ./run.sh
   # Or directly:
   go run cmd/api/main.go
   ```

## Testing the API

Once the service is running, you can test it:

### Using the test script:
```bash
./test-api.sh
```

### Using curl directly:
```bash
# Get products by category
curl http://localhost:8081/v1/product?category=1

# Add a product
curl -X POST http://localhost:8081/v1/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Hamburguer",
    "category": 1,
    "price": 34.99,
    "description": "Hamburguer com salada",
    "image_link": "https://example.com/burger.jpg"
  }'
```

### Using the HTTP file:
Open `http/product.http` in your IDE (VS Code with REST Client extension, or IntelliJ)

### Swagger Documentation:
Once running, visit: `http://localhost:8081/swagger/index.html`

## Category Values
- 1 - Lanche
- 2 - Acompanhamento  
- 3 - Bebida
- 4 - Sobremesa

## Troubleshooting

- **Port 8081 already in use**: Change `PORT` environment variable
- **Database connection error**: Check PostgreSQL is running and credentials are correct
- **Docker not running**: Start Docker Desktop or use Option 2

