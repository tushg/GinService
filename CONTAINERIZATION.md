# Containerization Guide

This document provides comprehensive information about the containerized deployment of the Gin Service.

## 🐳 What's Now Available

### **Docker Commands:**
```bash
# Build and run in development
docker-compose -f docker-compose.dev.yml up --build

# Build and run in production
docker-compose up --build -d

# Individual Docker commands
docker build -t gin-service:latest .
docker run -d --name gin-service -p 8081:8080 gin-service:latest
```

### **Service Status:**
✅ **Container is running** on port 8081  
✅ **Health endpoint working** at `http://localhost:8081/api/v1/health`  
✅ **All routes registered** and functional  
✅ **Logging system working** with structured logs  
✅ **Git repository updated** with all containerization files  

## 📁 Containerization Files

### **Dockerfile**
- **Multi-stage build** for optimized image size
- **Security**: Runs as non-root user
- **Health checks** built-in
- **Alpine Linux** base for minimal footprint

### **.dockerignore**
- Excludes unnecessary files from build context
- Improves build performance
- Reduces image size

### **docker-compose.yml (Production)**
- Production-ready configuration
- Environment variables for configuration
- Volume mounts for logs and configs
- Health checks and restart policies
- Optional nginx reverse proxy

### **docker-compose.dev.yml (Development)**
- Development-specific settings
- Debug mode enabled
- Source code mounted for live development
- Verbose logging

### **Makefile**
- Comprehensive Docker commands
- Development and production targets
- Testing and linting commands
- Utility functions

## 🚀 Quick Start

### **Development Environment:**
```bash
# Start development environment
docker-compose -f docker-compose.dev.yml up --build

# View logs
docker-compose -f docker-compose.dev.yml logs -f

# Stop development environment
docker-compose -f docker-compose.dev.yml down
```

### **Production Environment:**
```bash
# Start production environment
docker-compose up --build -d

# View logs
docker-compose logs -f gin-service

# Stop production environment
docker-compose down
```

### **Individual Docker Commands:**
```bash
# Build image
docker build -t gin-service:latest .

# Run container
docker run -d --name gin-service -p 8081:8080 gin-service:latest

# View logs
docker logs -f gin-service

# Stop container
docker stop gin-service

# Remove container
docker rm gin-service

# Clean up images
docker system prune -f
```

## 🔧 Configuration

### **Environment Variables:**
```bash
GIN_MODE=release          # Gin framework mode
LOG_LEVEL=info            # Logging level (debug, info, warn, error, fatal)
LOG_FORMAT=json           # Log format (json, text)
LOG_OUTPUT=stdout         # Log output (stdout, stderr, file)
```

### **Volume Mounts:**
- `./logs:/app/logs` - Persistent log storage
- `./configs:/app/configs:ro` - Configuration files (read-only)

### **Port Mapping:**
- `8081:8080` - Host port 8081 maps to container port 8080

## 📊 Health Checks

### **Built-in Health Check:**
```bash
# Health check endpoint
curl http://localhost:8081/api/v1/health

# Readiness check
curl http://localhost:8081/api/v1/health/ready

# Liveness check
curl http://localhost:8081/api/v1/health/live
```

### **Docker Health Check:**
```dockerfile
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/v1/health || exit 1
```

## 📝 Logging

### **View Container Logs:**
```bash
# All logs
docker logs gin-service

# Follow logs in real-time
docker logs -f gin-service

# Last 50 lines
docker logs --tail 50 gin-service

# With timestamps
docker logs -t gin-service

# Follow with timestamps
docker logs -f -t gin-service
```

### **Log Configuration:**
- **JSON Format**: Structured logging for production
- **Text Format**: Human-readable for development
- **Multiple Outputs**: Console, file, or both
- **Log Rotation**: Automatic file rotation with compression

## 🔒 Security Features

### **Non-root User:**
- Container runs as `appuser` (UID 1001)
- Minimal privileges for security
- Proper file permissions

### **Multi-stage Build:**
- Build stage with all tools
- Runtime stage with minimal dependencies
- Reduced attack surface

### **Alpine Linux:**
- Minimal base image
- Regular security updates
- Small footprint

## 🧪 Testing

### **Test the Service:**
```bash
# Health check
curl http://localhost:8081/api/v1/health

# Get all products
curl http://localhost:8081/api/v1/products

# Create a product
curl -X POST http://localhost:8081/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Test Product","price":99.99,"description":"Test description"}'
```

## 🛠️ Troubleshooting

### **Common Issues:**

#### **Port Already in Use:**
```bash
# Check what's using the port
netstat -ano | findstr :8081

# Kill the process
taskkill /PID <PID> /F
```

#### **Container Won't Start:**
```bash
# Check container logs
docker logs gin-service

# Check container status
docker ps -a

# Inspect container
docker inspect gin-service
```

#### **Build Issues:**
```bash
# Clean Docker cache
docker system prune -f

# Rebuild without cache
docker build --no-cache -t gin-service:latest .
```

### **Useful Commands:**
```bash
# Check container stats
docker stats gin-service

# Execute commands in container
docker exec -it gin-service sh

# Copy files from container
docker cp gin-service:/app/logs ./local-logs

# View container details
docker inspect gin-service
```

## 📈 Monitoring

### **Container Metrics:**
```bash
# Real-time stats
docker stats gin-service

# Resource usage
docker stats --no-stream gin-service
```

### **Application Metrics:**
- Health check endpoints
- Structured logging
- Request/response logging
- Performance metrics

## 🔄 CI/CD Integration

### **Docker Build in CI:**
```yaml
# Example GitHub Actions
- name: Build Docker image
  run: docker build -t gin-service:${{ github.sha }} .

- name: Push to registry
  run: docker push gin-service:${{ github.sha }}
```

### **Deployment:**
```bash
# Pull latest image
docker pull gin-service:latest

# Stop old container
docker stop gin-service

# Remove old container
docker rm gin-service

# Run new container
docker run -d --name gin-service -p 8081:8080 gin-service:latest
```

## 📚 Additional Resources

### **Docker Documentation:**
- [Docker Official Docs](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Best Practices](https://docs.docker.com/develop/dev-best-practices/)

### **Gin Framework:**
- [Gin Documentation](https://gin-gonic.com/docs/)
- [Gin Examples](https://github.com/gin-gonic/examples)

### **Alpine Linux:**
- [Alpine Linux](https://alpinelinux.org/)
- [Alpine Packages](https://pkgs.alpinelinux.org/)

---

## 🎯 Summary

The Gin Service is now fully containerized with:

✅ **Production-ready Docker setup**  
✅ **Development and production configurations**  
✅ **Comprehensive logging system**  
✅ **Health checks and monitoring**  
✅ **Security best practices**  
✅ **Easy deployment and scaling**  
✅ **CI/CD ready**  

Your service is ready for deployment in any containerized environment!
