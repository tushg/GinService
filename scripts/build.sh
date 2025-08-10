#!/bin/bash
echo "Building gin-service..."
go build -o bin/server cmd/server/main.go
echo "Build complete!"
