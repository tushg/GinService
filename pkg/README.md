# Pkg Directory - Utility Packages

This directory contains reusable utility packages that can be imported by other projects. These packages follow Go best practices and industry standards.

## 📁 Directory Structure

```
pkg/
├── README.md           # This file
├── utils/              # General utility functions
│   ├── string.go       # String manipulation utilities
│   └── time.go         # Time manipulation utilities
├── common/             # Common functionality
│   ├── errors.go       # Standardized error handling
│   └── responses.go    # Standardized API responses
└── constants/          # Application constants
    └── app.go          # Application-wide constants
```

## 🛠️ Available Packages

### 1. `pkg/utils/` - General Utilities

#### String Utilities (`pkg/utils/string.go`)
```go
import "gin-service/pkg/utils"

stringUtils := utils.NewStringUtils()

// Generate UUID
uuid := stringUtils.GenerateUUID()

// Generate random string
randomStr := stringUtils.GenerateRandomString(10)

// Check if string is empty
isEmpty := stringUtils.IsEmpty("")

// Convert to title case
titleCase := stringUtils.ToTitleCase("hello world")

// Truncate string
truncated := stringUtils.Truncate("very long string", 10)

// Remove special characters
clean := stringUtils.RemoveSpecialChars("Hello@World!")

// Create URL-friendly slug
slug := stringUtils.Slugify("Hello World!")
```

#### Time Utilities (`pkg/utils/time.go`)
```go
import "gin-service/pkg/utils"

timeUtils := utils.NewTimeUtils()

// Format datetime
formatted := timeUtils.FormatDateTime(time.Now())

// Get current timestamp
timestamp := timeUtils.GetCurrentTimestamp()

// Check if date is today
isToday := timeUtils.IsToday(someDate)

// Calculate age
age := timeUtils.GetAge(birthDate)

// Format duration
duration := timeUtils.FormatDuration(5 * time.Minute)

// Get start/end of day/week
startOfDay := timeUtils.GetStartOfDay(time.Now())
endOfWeek := timeUtils.GetEndOfWeek(time.Now())
```

### 2. `pkg/common/` - Common Functionality

#### Error Handling (`pkg/common/errors.go`)
```go
import "gin-service/pkg/common"

// Create different types of errors
validationErr := common.NewValidationError("Invalid input")
notFoundErr := common.NewNotFoundError("Resource not found")
internalErr := common.NewInternalError("Something went wrong")

// Create error with details
detailedErr := common.NewValidationErrorWithDetails("Invalid input", "Field 'email' is required")

// Create error with underlying error
dbErr := common.NewDatabaseErrorWithErr("Database operation failed", underlyingError)

// Check if error is AppError
if common.IsAppError(err) {
    appErr := common.GetAppError(err)
    statusCode := appErr.HTTPStatus
}
```

#### API Responses (`pkg/common/responses.go`)
```go
import "gin-service/pkg/common"

// Send success responses
common.SendSuccess(c, "Operation successful", data)
common.SendCreated(c, "Resource created", newResource)

// Send error responses
common.SendValidationError(c, "Invalid input")
common.SendNotFound(c, "Resource not found")
common.SendInternalError(c, "Something went wrong")

// Send paginated response
pagination := common.CalculatePagination(page, limit, total)
common.SendPaginatedSuccess(c, "Resources retrieved", data, pagination)
```

### 3. `pkg/constants/` - Application Constants

#### Application Constants (`pkg/constants/app.go`)
```go
import "gin-service/pkg/constants"

// Use application constants
appName := constants.AppName
apiVersion := constants.APIVersion

// Use environment constants
if env == constants.EnvProduction {
    // Production logic
}

// Use HTTP constants
contentType := constants.ContentTypeJSON
method := constants.MethodPOST

// Use validation constants
if len(input) > constants.MaxStringLength {
    // Handle validation
}
```

## 🎯 Usage Examples

### Example 1: Using Utilities in a Service
```go
package product

import (
    "gin-service/pkg/utils"
    "gin-service/pkg/common"
    "gin-service/pkg/constants"
)

type ProductService struct {
    stringUtils *utils.StringUtils
    timeUtils   *utils.TimeUtils
}

func NewProductService() *ProductService {
    return &ProductService{
        stringUtils: utils.NewStringUtils(),
        timeUtils:   utils.NewTimeUtils(),
    }
}

func (s *ProductService) CreateProduct(product *Product) error {
    // Generate UUID for product
    product.ID = s.stringUtils.GenerateUUID()
    
    // Set creation timestamp
    product.CreatedAt = s.timeUtils.GetCurrentTimestamp()
    
    // Create slug for URL
    product.Slug = s.stringUtils.Slugify(product.Name)
    
    return nil
}
```

### Example 2: Using Common Packages in Handler
```go
package product

import (
    "gin-service/pkg/common"
    "gin-service/pkg/constants"
)

func (h *ProductHandler) GetProduct(c *gin.Context) {
    id := c.Param("id")
    
    product, err := h.service.GetProduct(id)
    if err != nil {
        common.SendNotFound(c, "Product not found")
        return
    }
    
    common.SendSuccess(c, constants.MsgRetrieved, product)
}
```

### Example 3: Using Constants in Configuration
```go
package config

import "gin-service/pkg/constants"

func Load() *Config {
    return &Config{
        AppName:    constants.AppName,
        AppVersion: constants.AppVersion,
        Port:       constants.DefaultPort,
        Environment: constants.EnvDevelopment,
    }
}
```

## 🔧 Best Practices

### 1. Package Organization
- **`pkg/`** for reusable, public packages
- **`internal/`** for application-specific code
- Keep packages focused and single-purpose

### 2. Error Handling
- Use standardized error types from `pkg/common/errors.go`
- Always include appropriate HTTP status codes
- Provide meaningful error messages

### 3. API Responses
- Use standardized response format from `pkg/common/responses.go`
- Include consistent metadata (timestamp, path, method)
- Handle pagination properly

### 4. Constants
- Use constants from `pkg/constants/app.go` instead of magic numbers
- Keep constants organized by category
- Use descriptive names

### 5. Utilities
- Keep utility functions pure and stateless
- Provide clear documentation
- Include examples in comments

## 🚀 Benefits

1. **Reusability** - Packages can be imported by other projects
2. **Consistency** - Standardized error handling and responses
3. **Maintainability** - Centralized utility functions
4. **Testability** - Easy to unit test utility functions
5. **Documentation** - Clear API documentation
6. **Industry Standard** - Follows Go project conventions

## 📝 Adding New Utilities

When adding new utilities:

1. **Choose the right package** - `utils/`, `common/`, or `constants/`
2. **Follow naming conventions** - Use descriptive names
3. **Add documentation** - Include examples and usage
4. **Write tests** - Ensure good test coverage
5. **Update this README** - Document new functionality

## 🔗 Related Documentation

- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Go Package Guidelines](https://golang.org/doc/effective_go.html#package-names)
- [Go Error Handling](https://golang.org/doc/effective_go.html#errors)
