#!/bin/bash

# Script to generate test coverage reports locally
# Usage: ./scripts/coverage.sh

set -e

echo "🧪 Running tests with coverage..."
# Exclude mocks, docs, cmd/api, and internal/app from coverage (infrastructure code)
go test $(go list ./... | grep -v '/mocks/' | grep -v '/docs' | grep -v 'cmd/api' | grep -v 'internal/app') -coverprofile=coverage.out -covermode=atomic

echo ""
echo "📊 Coverage summary:"
go tool cover -func=coverage.out

echo ""
echo "📈 Generating HTML coverage report..."
go tool cover -html=coverage.out -o coverage.html

echo ""
echo "✅ Coverage reports generated:"
echo "   - coverage.out (for SonarCloud)"
echo "   - coverage.html (for local viewing)"
echo ""
echo "💡 Open coverage.html in your browser to see detailed coverage"