package product

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// productRepository implements ProductRepository interface
type productRepository struct {
	products map[string]*Product
	mutex    sync.RWMutex
}

// NewProductRepository creates a new product repository instance
func NewProductRepository() ProductRepository {
	return &productRepository{
		products: make(map[string]*Product),
	}
}

// Create adds a new product to the repository
func (r *productRepository) Create(ctx context.Context, product *Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Generate ID if not provided
	if product.ID == "" {
		product.ID = uuid.New().String()
	}

	// Set timestamps
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	// Store the product
	r.products[product.ID] = product
	return nil
}

// GetByID retrieves a product by ID
func (r *productRepository) GetByID(ctx context.Context, id string) (*Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, fmt.Errorf("product not found: %s", id)
	}

	return product, nil
}

// GetAll retrieves all products with pagination
func (r *productRepository) GetAll(ctx context.Context, limit, offset int) ([]*Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	products := make([]*Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}

	// Apply pagination
	if offset >= len(products) {
		return []*Product{}, nil
	}

	end := offset + limit
	if end > len(products) {
		end = len(products)
	}

	return products[offset:end], nil
}

// Update updates an existing product
func (r *productRepository) Update(ctx context.Context, product *Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.products[product.ID]; !exists {
		return fmt.Errorf("product not found: %s", product.ID)
	}

	product.UpdatedAt = time.Now()
	r.products[product.ID] = product
	return nil
}

// Delete removes a product by ID
func (r *productRepository) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.products[id]; !exists {
		return fmt.Errorf("product not found: %s", id)
	}

	delete(r.products, id)
	return nil
}

// Count returns the total number of products
func (r *productRepository) Count(ctx context.Context) (int64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return int64(len(r.products)), nil
}
