# Step 4: Configuration Files Setup

## Create Configuration YAML File
Create `configs/config.yaml`:

```yaml
# Server Configuration
server:
  port: 8080
  host: "0.0.0.0"
  read_timeout: 30s
  write_timeout: 30s
  max_header_bytes: 1048576

# Database Configuration
database:
  host: "localhost"
  port: 5432
  name: "gin_service"
  user: "postgres"
  password: "password"
  sslmode: "disable"
  max_open_conns: 25
  max_idle_conns: 5
  conn_max_lifetime: "5m"

# Logging Configuration
logging:
  level: "info"
  format: "json"
  output: "stdout"
  file:
    enabled: false
    path: "logs/app.log"
    max_size: 100
    max_age: 30
    max_backups: 10
    compress: true

# Environment
environment: "development"
```

## Create Environment File
Create `.env` file in the root directory:

```bash
# Server
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=gin_service
DB_USER=postgres
DB_PASSWORD=password
DB_SSLMODE=disable

# Logging
LOG_LEVEL=info
LOG_FORMAT=json

# Environment
ENV=development
```

## Create .env.example File
Create `.env.example` file for reference:

```bash
# Server
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=gin_service
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_SSLMODE=disable

# Logging
LOG_LEVEL=info
LOG_FORMAT=json

# Environment
ENV=development
```

## Update .gitignore
Add environment files to `.gitignore`:

```bash
# Add these lines to .gitignore
echo "
# Environment files
.env
.env.local
.env.production" >> .gitignore
```

## Verify Configuration Files
```bash
# Check if files are created
ls -la configs/
ls -la .env*

# Expected output should show:
# configs/config.yaml
# .env
# .env.example
```

## Next Steps
After creating configuration files, proceed to the next file to create the configuration management Go code.
