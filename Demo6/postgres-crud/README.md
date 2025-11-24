# PostgreSQL CRUD Operations - Go REST API

A well-structured Go application demonstrating CRUD operations with PostgreSQL using GORM, following best practices and clean architecture principles.

## Project Structure

```
postgres-crud/
├── cmd/
│   └── api/
│       └── main.go          # REST API server entry point
├── internal/
│   ├── handler/             # HTTP handlers (controllers)
│   │   └── order_handler.go
│   ├── dto/                 # Data Transfer Objects
│   │   └── order_dto.go
│   ├── middleware/          # HTTP middleware
│   │   ├── logger.go
│   │   ├── recovery.go
│   │   └── cors.go
│   ├── router/              # Route definitions
│   │   └── router.go
│   └── errors/              # Error handling
│       └── errors.go
├── config/                  # Configuration management
│   └── config.go
├── database/                # Database connection and migration
│   └── database.go
├── model/                   # Data models
│   └── Order.go
├── repository/              # Data access layer
│   └── order_repository.go
├── service/                 # Business logic layer
│   └── order_service.go
├── go.mod                   # Go module dependencies
├── go.sum                   # Go module checksums
├── docker-compose.yml       # PostgreSQL Docker setup
├── Dockerfile               # Application Dockerfile
├── API.md                   # API documentation
├── Postman_Collection.json  # Postman API collection
├── Postman_Environment.json # Postman environment variables
├── POSTMAN_SETUP.md         # Postman setup guide
└── README.md                # Project documentation
```

## Architecture

This project follows a **layered architecture** pattern with REST API support:

1. **Handler Layer** (`internal/handler/`): HTTP request/response handling
2. **DTO Layer** (`internal/dto/`): Data Transfer Objects for API contracts
3. **Service Layer** (`service/`): Business logic and validation
4. **Repository Layer** (`repository/`): Data access operations
5. **Model Layer** (`model/`): Domain models and database schema
6. **Middleware** (`internal/middleware/`): HTTP middleware (logging, CORS, recovery)
7. **Router** (`internal/router/`): Route configuration
8. **Config Layer** (`config/`): Application configuration
9. **Database Layer** (`database/`): Database connection and migrations

## Features

- ✅ **REST API** with Gin framework
- ✅ Modern GORM (`gorm.io/gorm`) instead of deprecated `github.com/jinzhu/gorm`
- ✅ Clean separation of concerns (Handler, Service, Repository, Model layers)
- ✅ Interface-based design for testability
- ✅ Environment-based configuration
- ✅ Request validation with go-playground/validator
- ✅ Proper error handling with custom error types
- ✅ Database migrations
- ✅ Soft deletes support
- ✅ CORS middleware
- ✅ Request logging middleware
- ✅ Panic recovery middleware
- ✅ Docker support for PostgreSQL

## Prerequisites

- Go 1.23 or higher
- PostgreSQL database (or Docker for running PostgreSQL)
- Go modules enabled
- Docker and Docker Compose (optional, for containerized PostgreSQL)

## Configuration

The application uses environment variables for configuration. You can set these or use the defaults:

```bash
# Database Configuration
DB_HOST=localhost      # Default: localhost
DB_PORT=5432          # Default: 5432
DB_USER=postgres      # Default: postgres
DB_PASSWORD=password   # Default: password
DB_NAME=userdb        # Default: userdb
DB_SSLMODE=disable    # Default: disable

# Server Configuration
SERVER_HOST=0.0.0.0   # Default: 0.0.0.0
SERVER_PORT=8080      # Default: 8080
```

## Quick Start

1. **Start PostgreSQL** (using Docker):
```bash
docker-compose up -d
```

2. **Install dependencies:**
```bash
go mod download
```

3. **Run the REST API server:**
```bash
go run cmd/api/main.go
```

4. **Test the API:**
```bash
# Health check
curl http://localhost:8080/health

# Create an order
curl -X POST http://localhost:8080/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{"description": "Test Order"}'
```

The API will be available at `http://localhost:8080` (or the port specified in `SERVER_PORT` environment variable).

## Installation

### Option 1: Run from Source

1. **Install dependencies:**
```bash
go mod download
```

2. **Set up your PostgreSQL database** and update the configuration if needed (or use Docker Compose).

3. **Run the REST API server:**
```bash
go run cmd/api/main.go
```

### Option 2: Build Binary

```bash
# Build the application
go build -o postgres-crud.exe ./cmd/api

# Run the binary
./postgres-crud.exe
```

## Docker Setup

### Running PostgreSQL with Docker Compose

The easiest way to get started is using Docker Compose to run PostgreSQL:

1. **Start PostgreSQL container:**
```bash
docker-compose up -d
```

This will:
- Start a PostgreSQL 15 container
- Create a database named `userdb`
- Set up user `postgres` with password `password`
- Expose PostgreSQL on port `5432`
- Create a persistent volume for data

2. **Check if PostgreSQL is running:**
```bash
docker-compose ps
```

3. **View logs:**
```bash
docker-compose logs -f postgres
```

4. **Stop PostgreSQL:**
```bash
docker-compose down
```

5. **Stop and remove volumes (⚠️ deletes data):**
```bash
docker-compose down -v
```

### Running the Application in Docker

You can also run the entire application in Docker:

1. **Build the Docker image:**
```bash
docker build -t postgres-crud .
```

2. **Run with Docker Compose (database + app):**
   - Update `docker-compose.yml` to include the app service
   - Or run manually:
```bash
docker run --network host -e DB_HOST=localhost postgres-crud
```

### Docker Compose Configuration

The `docker-compose.yml` file includes:
- PostgreSQL 15 Alpine (lightweight)
- Health checks
- Persistent data volume
- Port mapping (5432:5432)
- Environment variables for database setup

Default credentials (matching config defaults):
- **Host**: `localhost`
- **Port**: `5432`
- **User**: `postgres`
- **Password**: `password`
- **Database**: `userdb`

## Usage

### REST API

The application provides a REST API with the following endpoints:

- **POST** `/api/v1/orders` - Create a new order
- **GET** `/api/v1/orders` - Get all orders
- **GET** `/api/v1/orders/:id` - Get order by ID
- **PUT** `/api/v1/orders/:id` - Update an order
- **DELETE** `/api/v1/orders/:id` - Delete an order
- **GET** `/health` - Health check endpoint

See [API.md](API.md) for detailed API documentation.

## Code Examples

### REST API Usage

**Create Order:**
```bash
curl -X POST http://localhost:8080/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{"description": "Laptop"}'
```

**Get All Orders:**
```bash
curl http://localhost:8080/api/v1/orders
```

**Get Order by ID:**
```bash
curl http://localhost:8080/api/v1/orders/1
```

**Update Order:**
```bash
curl -X PUT http://localhost:8080/api/v1/orders/1 \
  -H "Content-Type: application/json" \
  -d '{"description": "Updated Description"}'
```

**Delete Order:**
```bash
curl -X DELETE http://localhost:8080/api/v1/orders/1
```

### Programmatic Usage

```go
orderService := service.NewOrderService(repository.NewOrderRepository())
order, err := orderService.CreateOrder("Laptop")
```

## Testing

### Using Postman

1. **Import the Collection:**
   - Open Postman
   - Click "Import" button
   - Select `Postman_Collection.json` file
   - The collection will be imported with all endpoints

2. **Import the Environment (Optional):**
   - Click "Import" button
   - Select `Postman_Environment.json` file
   - Select the imported environment from the dropdown
   - This sets up variables like `base_url` and `order_id`

3. **Set Environment Variables:**
   - `base_url`: `http://localhost:8080` (default)
   - `order_id`: Will be automatically set when creating orders

4. **Run Requests:**
   - Start the API server: `go run cmd/api/main.go`
   - Use the collection to test all endpoints
   - The "Create Order" request automatically saves the order ID to the `order_id` variable

## License

This project is for educational purposes.

