package product

import "context"

// ProductService defines the interface for product business logic
type ProductService interface {
	CreateProduct(ctx context.Context, req *CreateProductRequest) (*ProductResponse, error)
	GetProduct(ctx context.Context, id string) (*ProductResponse, error)
	GetAllProducts(ctx context.Context, req *GetProductsRequest) (*GetProductsResponse, error)
	UpdateProduct(ctx context.Context, id string, req *UpdateProductRequest) (*ProductResponse, error)
	DeleteProduct(ctx context.Context, id string) error
}

// ProductRepository defines the interface for product data access
type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}
