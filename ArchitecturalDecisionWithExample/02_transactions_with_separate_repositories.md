# Transactions with Separate Repositories: Implementation Guide

## Overview

This document demonstrates how to handle transactions, rollbacks, joins, and complex multi-table operations when using separate repositories per module. Contrary to common concerns, **separate repositories can handle all these scenarios effectively**.

## Key Question

> "If we go with separate repository for each module, can we still handle transactions/rollbacks, joins, or dependencies like inserting into multiple tables?"

**Answer: YES, absolutely!** Here are several proven approaches:

## Approach 1: Transaction Manager (Recommended)

Create a shared transaction manager that all repositories can use:

```go
// pkg/database/transaction.go
package database

import (
    "context"
    "database/sql"
    "fmt"
)

type TransactionManager interface {
    WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error
    BeginTx(ctx context.Context) (*sql.Tx, error)
}

type transactionManager struct {
    db *sql.DB
}

func NewTransactionManager(db *sql.DB) TransactionManager {
    return &transactionManager{db: db}
}

func (tm *transactionManager) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
    tx, err := tm.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        }
    }()
    
    if err := fn(tx); err != nil {
        if rbErr := tx.Rollback(); rbErr != nil {
            return fmt.Errorf("tx failed: %v, rollback failed: %v", err, rbErr)
        }
        return err
    }
    
    return tx.Commit()
}
```

## Approach 2: Repository Composition

Each repository can accept a transaction and fall back to the default connection:

```go
// internal/health/repository.go
type HealthRepository interface {
    GetHealth(ctx context.Context) (*Health, error)
    CreateHealth(ctx context.Context, health *Health) error
    CreateHealthWithTx(ctx context.Context, tx *sql.Tx, health *Health) error
}

type healthRepository struct {
    db     *sql.DB
    logger logger.Logger
}

func (r *healthRepository) CreateHealth(ctx context.Context, health *Health) error {
    return r.CreateHealthWithTx(ctx, nil, health)
}

func (r *healthRepository) CreateHealthWithTx(ctx context.Context, tx *sql.Tx, health *Health) error {
    query := `INSERT INTO health (id, status, timestamp) VALUES ($1, $2, $3)`
    
    if tx != nil {
        _, err := tx.ExecContext(ctx, query, health.ID, health.Status, health.Timestamp)
        return err
    }
    
    _, err := r.db.ExecContext(ctx, query, health.ID, health.Status, health.Timestamp)
    return err
}
```

## Approach 3: Service Layer Orchestration

Handle complex operations in the service layer using the transaction manager:

```go
// internal/product/service.go
type ProductService struct {
    productRepo  ProductRepository
    healthRepo   HealthRepository
    txManager    database.TransactionManager
    logger       logger.Logger
}

func (s *ProductService) CreateProductWithHealthCheck(ctx context.Context, product *Product) error {
    return s.txManager.WithTransaction(ctx, func(tx *sql.Tx) error {
        // Create product
        if err := s.productRepo.CreateProductWithTx(ctx, tx, product); err != nil {
            return fmt.Errorf("failed to create product: %w", err)
        }
        
        // Create health check record
        health := &Health{
            ID:        uuid.New().String(),
            Status:    "product_created",
            Timestamp: time.Now(),
        }
        
        if err := s.healthRepo.CreateHealthWithTx(ctx, tx, health); err != nil {
            return fmt.Errorf("failed to create health record: %w", err)
        }
        
        return nil
    })
}
```

## Approach 4: Cross-Module Queries

For complex joins, create a dedicated query service:

```go
// internal/query/service.go
type QueryService struct {
    db     *sql.DB
    logger logger.Logger
}

func (s *QueryService) GetProductWithHealthStatus(ctx context.Context, productID string) (*ProductWithHealth, error) {
    query := `
        SELECT 
            p.id, p.name, p.price,
            h.status, h.timestamp
        FROM products p
        LEFT JOIN health h ON h.product_id = p.id
        WHERE p.id = $1
    `
    
    var result ProductWithHealth
    err := s.db.QueryRowContext(ctx, query, productID).Scan(
        &result.ID, &result.Name, &result.Price,
        &result.HealthStatus, &result.HealthTimestamp,
    )
    
    if err != nil {
        return nil, err
    }
    
    return &result, nil
}
```

## Complete Example: Multi-Table Insert with Transaction

```go
// internal/order/service.go
type OrderService struct {
    orderRepo    OrderRepository
    productRepo  ProductRepository
    userRepo     UserRepository
    txManager    database.TransactionManager
    logger       logger.Logger
}

func (s *OrderService) CreateOrder(ctx context.Context, order *Order) error {
    return s.txManager.WithTransaction(ctx, func(tx *sql.Tx) error {
        // 1. Create order
        if err := s.orderRepo.CreateOrderWithTx(ctx, tx, order); err != nil {
            return fmt.Errorf("failed to create order: %w", err)
        }
        
        // 2. Update product inventory
        for _, item := range order.Items {
            if err := s.productRepo.UpdateInventoryWithTx(ctx, tx, item.ProductID, item.Quantity); err != nil {
                return fmt.Errorf("failed to update inventory: %w", err)
            }
        }
        
        // 3. Update user order count
        if err := s.userRepo.IncrementOrderCountWithTx(ctx, tx, order.UserID); err != nil {
            return fmt.Errorf("failed to update user: %w", err)
        }
        
        return nil
    })
}
```

## File Structure

```
pkg/database/
├── manager.go
├── transaction.go          # Transaction manager
└── postgresql/
    └── connection.go

internal/
├── health/
│   ├── repository.go       # Accepts tx parameter
│   └── service.go
├── product/
│   ├── repository.go       # Accepts tx parameter
│   └── service.go
├── order/
│   ├── repository.go       # Orchestrates multiple operations
│   └── service.go
└── query/
    └── service.go         # Complex cross-module queries
```

## Benefits of This Approach

1. **✅ Transactions Work Perfectly**: Full ACID compliance
2. **✅ Rollbacks Automatic**: Built into the transaction manager
3. **✅ Complex Operations**: Can span multiple repositories
4. **✅ Maintains Separation**: Each repository still owns its domain
5. **✅ Testable**: Easy to mock and test
6. **✅ Flexible**: Can use transactions or not as needed

## Transaction Flow Example

```
Service Layer
    ↓
Transaction Manager
    ↓
Begin Transaction
    ↓
Repository 1 (with tx) → Repository 2 (with tx) → Repository 3 (with tx)
    ↓
Commit (if all succeed) or Rollback (if any fail)
```

## Testing with Transactions

```go
func TestOrderService_CreateOrder(t *testing.T) {
    // Mock repositories
    mockOrderRepo := &MockOrderRepository{}
    mockProductRepo := &MockProductRepository{}
    mockUserRepo := &MockUserRepository{}
    
    // Mock transaction manager
    mockTxManager := &MockTransactionManager{}
    
    service := &OrderService{
        orderRepo:   mockOrderRepo,
        productRepo: mockProductRepo,
        userRepo:    mockUserRepo,
        txManager:   mockTxManager,
    }
    
    // Test transaction flow
    // ... test implementation
}
```

## Conclusion

**Separate repositories per module + Transaction Manager = Best of both worlds!** 

You get:
- ✅ Clean separation of concerns
- ✅ Full transaction support
- ✅ Complex multi-table operations
- ✅ Easy testing and maintenance
- ✅ Go-idiomatic design

This approach is actually **more flexible** than centralized repositories because each module can work independently OR participate in transactions when needed. It's the pattern used by many successful Go applications! 🚀

## Key Takeaways

1. **Transactions are fully supported** with separate repositories
2. **Rollbacks are automatic** through the transaction manager
3. **Complex operations** can span multiple repositories
4. **Joins and cross-module queries** are handled through dedicated services
5. **Testing remains simple** with proper mocking
6. **Maintains clean architecture** while providing flexibility
