package product

import (
	"time"
)

// Product represents the product entity
type Product struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	Category    string    `json:"category" db:"category"`
	Stock       int       `json:"stock" db:"stock"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CreateProductRequest represents the request for creating a product
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Category    string  `json:"category" binding:"required"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
}

// UpdateProductRequest represents the request for updating a product
type UpdateProductRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price" binding:"omitempty,gt=0"`
	Category    *string  `json:"category"`
	Stock       *int     `json:"stock" binding:"omitempty,gte=0"`
}

// ProductResponse represents the response for product operations
type ProductResponse struct {
	Product *Product `json:"product"`
	Message string   `json:"message,omitempty"`
}

// GetProductsRequest represents the request for getting products with pagination
type GetProductsRequest struct {
	Limit  int `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset int `form:"offset" binding:"omitempty,min=0"`
}

// GetProductsResponse represents the response for getting multiple products
type GetProductsResponse struct {
	Products []*Product `json:"products"`
	Total    int64      `json:"total"`
	Limit    int        `json:"limit"`
	Offset   int        `json:"offset"`
}
