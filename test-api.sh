#!/bin/bash

# Test script for Product API
BASE_URL="http://localhost:8081"

echo "=== Testing Product API ==="
echo ""

# Test 1: Get products by category (should work even with empty DB)
echo "1. Testing GET /v1/product?category=1"
curl -s -X GET "${BASE_URL}/v1/product?category=1" | jq '.' || echo "Failed or no jq installed"
echo ""
echo ""

# Test 2: Add a product
echo "2. Testing POST /v1/product"
curl -s -X POST "${BASE_URL}/v1/product" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Hamburguer",
    "category": 1,
    "price": 34.99,
    "description": "Hamburguer com salada",
    "image_link": "https://example.com/burger.jpg"
  }' | jq '.' || echo "Failed"
echo ""
echo ""

# Test 3: Get products again
echo "3. Testing GET /v1/product?category=1 (after adding)"
curl -s -X GET "${BASE_URL}/v1/product?category=1" | jq '.' || echo "Failed"
echo ""
echo ""

# Test 4: Health check (if available)
echo "4. Testing Swagger endpoint"
curl -s -X GET "${BASE_URL}/swagger/index.html" | head -5 || echo "Swagger not available"
echo ""

