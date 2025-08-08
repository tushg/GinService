.PHONY: build run test clean deps lint docker-build docker-run docker-stop docker-clean docker-dev

# Go commands
build:
	go build -o bin/server cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

deps:
	go mod download
	go mod tidy

lint:
	golangci-lint run

# Docker commands
docker-build:
	docker build -t gin-service:latest .

docker-run:
	docker run -d --name gin-service -p 8080:8080 gin-service:latest

docker-stop:
	docker stop gin-service || true
	docker rm gin-service || true

docker-clean:
	docker system prune -f
	docker image prune -f

docker-dev:
	docker-compose -f docker-compose.dev.yml up --build

docker-prod:
	docker-compose up --build -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f gin-service

# Development with hot reload (requires air)
dev:
	air

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Database commands (if you add a database later)
db-migrate:
	# Add database migration commands here
	@echo "Database migration commands will be added here"

db-seed:
	# Add database seeding commands here
	@echo "Database seeding commands will be added here"

# Utility commands
fmt:
	go fmt ./...

vet:
	go vet ./...

# Show help
help:
	@echo "Available commands:"
	@echo "  build          - Build the application"
	@echo "  run            - Run the application"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  clean          - Clean build artifacts"
	@echo "  deps           - Download and tidy dependencies"
	@echo "  lint           - Run linter"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  docker-stop    - Stop Docker container"
	@echo "  docker-clean   - Clean Docker resources"
	@echo "  docker-dev     - Run with docker-compose (development)"
	@echo "  docker-prod    - Run with docker-compose (production)"
	@echo "  docker-down    - Stop docker-compose services"
	@echo "  docker-logs    - Show docker-compose logs"
	@echo "  dev            - Run with hot reload (requires air)"
	@echo "  install-tools  - Install development tools"
	@echo "  fmt            - Format code"
	@echo "  vet            - Vet code"
