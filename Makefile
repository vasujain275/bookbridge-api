.PHONY: build run dev swag test clean

# Build the application
build:
	@echo "Building the application..."
	@go build -o bin/api cmd/api/main.go

# Run the built binary
run: build
	@echo "Running the application..."
	@./bin/api

# Start development mode using Air for live reload
dev:
	@echo "Starting development mode with Air..."
	@swag init -g cmd/api/main.go
	@air

# Generate Swagger docs; ensure swag is installed: go install github.com/swaggo/swag/cmd/swag@latest
swag:
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/api/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean build artifacts and Swagger files
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin
	@rm -rf tmp
	@rm -rf swagger*
