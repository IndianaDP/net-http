.PHONY: build run clean test

# Binary name
BINARY_NAME=shortener

# Build the application
build:
	go build -o ../bin/$(BINARY_NAME) ./cmd/shortener

# Run the application
run:
	go run ./cmd/shortener/main.go

# Clean build artifacts
clean:
	rm -rf ../bin/
	go clean

# Run tests
test:
	go test -v ./...

# Build and run
dev: build
	../bin/$(BINARY_NAME)

# Get dependencies
deps:
	go mod download

# Run tests with coverage
test-coverage:
	go test -v -cover ./...