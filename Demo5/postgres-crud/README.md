# PostgreSQL CRUD Operations - Go REST API

A well-structured Go application demonstrating CRUD operations with PostgreSQL using GORM, following best practices and clean architecture principles.

## Project Structure

```
postgres-crud/
├── config/          # Configuration management
│   └── config.go    # Database and application configuration
├── database/        # Database connection and migration
│   └── database.go  # Database initialization and connection management
├── model/           # Data models
│   └── Order.go     # Order entity model
├── repository/       # Data access layer
│   └── order_repository.go  # Order repository interface and implementation
├── service/         # Business logic layer
│   └── order_service.go     # Order service interface and implementation
├── main.go          # Application entry point
├── go.mod           # Go module dependencies
└── README.md        # Project documentation
```

## Architecture

This project follows a **layered architecture** pattern:

1. **Model Layer** (`model/`): Defines the data structures and database schema
2. **Repository Layer** (`repository/`): Handles all database operations (data access)
3. **Service Layer** (`service/`): Contains business logic and validation
4. **Config Layer** (`config/`): Manages application configuration
5. **Database Layer** (`database/`): Handles database connection and migrations

## Features

- ✅ Modern GORM (`gorm.io/gorm`) instead of deprecated `github.com/jinzhu/gorm`
- ✅ Clean separation of concerns (Repository, Service, Model layers)
- ✅ Interface-based design for testability
- ✅ Environment-based configuration
- ✅ Proper error handling
- ✅ Database migrations
- ✅ Soft deletes support

## Prerequisites

- Go 1.21 or higher
- PostgreSQL database (or Docker for running PostgreSQL)
- Go modules enabled
- Docker and Docker Compose (optional, for containerized PostgreSQL)

## Configuration

The application uses environment variables for configuration. You can set these or use the defaults:

```bash
DB_HOST=localhost      # Default: localhost
DB_PORT=5432          # Default: 5432
DB_USER=postgres      # Default: postgres
DB_PASSWORD=password   # Default: password
DB_NAME=userdb        # Default: userdb
DB_SSLMODE=disable    # Default: disable
```

## Installation



1. Install dependencies:
```bash
go mod download
```

2. Set up your PostgreSQL database and update the configuration if needed.

3. Run the application:
```bash
go run main.go
```

Or build and run:
```bash
go build -o postgres-crud.exe .
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

The application demonstrates CRUD operations:

1. **CREATE**: Creates new orders
2. **READ**: Retrieves orders (single, all, filtered)
3. **UPDATE**: Updates order descriptions
4. **DELETE**: Deletes orders

## Code Examples

### Creating an Order
```go
orderService := service.NewOrderService(repository.NewOrderRepository())
order, err := orderService.CreateOrder("Laptop")
```

### Getting an Order by ID
```go
order, err := orderService.GetOrderByID(1)
```

### Updating an Order
```go
order, err := orderService.UpdateOrder(1, "Updated Description")
```

### Deleting an Order
```go
err := orderService.DeleteOrder(1)
```

## Testing

To test the application, ensure your PostgreSQL database is running and accessible with the configured credentials. The application will automatically create the `orders` table on first run.

## License

This project is for educational purposes.

