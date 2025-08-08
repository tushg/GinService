.PHONY: build run test clean deps lint

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run linter
lint:
	golangci-lint run

# Run with hot reload (requires air: go install github.com/cosmtrek/air@latest)
dev:
	air

# Build for production
build-prod:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server cmd/server/main.go

# Docker build
docker-build:
	docker build -t gin-service .

# Docker run
docker-run:
	docker run -p 8080:8080 gin-service
