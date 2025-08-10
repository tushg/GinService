package product

import (
	"context"
	"fmt"
)

// productService implements ProductService interface
type productService struct {
	repository ProductRepository
}

// NewProductService creates a new product service instance
func NewProductService(repository ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

// CreateProduct handles product creation business logic
func (s *productService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*ProductResponse, error) {
	// Validate business rules
	if req.Price <= 0 {
		return nil, fmt.Errorf("price must be greater than zero")
	}

	if req.Stock < 0 {
		return nil, fmt.Errorf("stock cannot be negative")
	}

	// Create product entity
	product := &Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Stock:       req.Stock,
	}

	// Save to repository
	if err := s.repository.Create(ctx, product); err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return &ProductResponse{
		Product: product,
		Message: "Product created successfully",
	}, nil
}

// GetProduct handles product retrieval business logic
func (s *productService) GetProduct(ctx context.Context, id string) (*ProductResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("product ID is required")
	}

	product, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return &ProductResponse{
		Product: product,
	}, nil
}

// GetAllProducts handles product listing business logic
func (s *productService) GetAllProducts(ctx context.Context, req *GetProductsRequest) (*GetProductsResponse, error) {
	// Set default pagination values
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	offset := req.Offset
	if offset < 0 {
		offset = 0
	}

	// Get products from repository
	products, err := s.repository.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	// Get total count
	total, err := s.repository.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get product count: %w", err)
	}

	return &GetProductsResponse{
		Products: products,
		Total:    total,
		Limit:    limit,
		Offset:   offset,
	}, nil
}

// UpdateProduct handles product update business logic
func (s *productService) UpdateProduct(ctx context.Context, id string, req *UpdateProductRequest) (*ProductResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("product ID is required")
	}

	// Get existing product
	existingProduct, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	// Apply updates
	if req.Name != nil {
		existingProduct.Name = *req.Name
	}
	if req.Description != nil {
		existingProduct.Description = *req.Description
	}
	if req.Price != nil {
		if *req.Price <= 0 {
			return nil, fmt.Errorf("price must be greater than zero")
		}
		existingProduct.Price = *req.Price
	}
	if req.Category != nil {
		existingProduct.Category = *req.Category
	}
	if req.Stock != nil {
		if *req.Stock < 0 {
			return nil, fmt.Errorf("stock cannot be negative")
		}
		existingProduct.Stock = *req.Stock
	}

	// Save to repository
	if err := s.repository.Update(ctx, existingProduct); err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return &ProductResponse{
		Product: existingProduct,
		Message: "Product updated successfully",
	}, nil
}

// DeleteProduct handles product deletion business logic
func (s *productService) DeleteProduct(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("product ID is required")
	}

	// Check if product exists
	if _, err := s.repository.GetByID(ctx, id); err != nil {
		return fmt.Errorf("failed to get product: %w", err)
	}

	// Delete from repository
	if err := s.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
