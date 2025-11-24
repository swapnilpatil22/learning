# Postman Collection Setup Guide

This guide will help you import and use the Postman collection for testing the PostgreSQL CRUD API.

## Files Included

1. **Postman_Collection.json** - Complete API collection with all endpoints
2. **Postman_Environment.json** - Environment variables for easy configuration

## Quick Start

### Step 1: Import Collection

1. Open Postman application
2. Click the **Import** button (top left)
3. Click **Upload Files** or drag and drop `Postman_Collection.json`
4. Click **Import**

### Step 2: Import Environment (Recommended)

1. Click **Import** again
2. Select `Postman_Environment.json`
3. Click **Import**
4. Select the imported environment from the environment dropdown (top right)

### Step 3: Start Your API Server

```bash
# Start PostgreSQL (if using Docker)
docker-compose up -d

# Start the API server
go run cmd/api/main.go
```

The server will start at `http://localhost:8080`

### Step 4: Test the API

1. Select the **PostgreSQL CRUD API** collection
2. Select the **PostgreSQL CRUD API - Local** environment (if imported)
3. Start testing endpoints!

## Collection Structure

### Health Check
- **GET** `/health` - Check server status

### Orders
- **POST** `/api/v1/orders` - Create a new order
- **GET** `/api/v1/orders` - Get all orders
- **GET** `/api/v1/orders/:id` - Get order by ID
- **PUT** `/api/v1/orders/:id` - Update an order
- **DELETE** `/api/v1/orders/:id` - Delete an order

## Features

### Automatic Order ID Capture
The "Create Order" request includes a test script that automatically saves the created order ID to the `order_id` environment variable. This means:
1. Create an order → ID is saved automatically
2. Use "Get Order by ID" → It will use the saved ID
3. Update/Delete → Same ID is used automatically

### Example Responses
Each request includes example responses for:
- Success scenarios (200, 201)
- Error scenarios (400, 404, 500)

### Environment Variables

The collection uses these variables:
- `base_url` - API base URL (default: `http://localhost:8080`)
- `order_id` - Current order ID (auto-set when creating orders)

## Customizing the Environment

### Change Base URL

If your API runs on a different port or host:

1. Select the environment from dropdown
2. Click the eye icon to view variables
3. Edit `base_url` value
4. Click **Save**

### Manual Order ID

To test with a specific order ID:

1. Select the environment
2. Edit `order_id` value
3. Click **Save**

## Testing Workflow

### Recommended Testing Order

1. **Health Check** - Verify server is running
2. **Create Order** - Create a test order (ID is auto-saved)
3. **Get All Orders** - See all orders including the new one
4. **Get Order by ID** - Get the specific order (uses saved ID)
5. **Update Order** - Update the order (uses saved ID)
6. **Delete Order** - Delete the order (uses saved ID)

### Testing Error Scenarios

The collection includes error examples:
- Try creating an order with description less than 3 characters (validation error)
- Try getting/updating/deleting a non-existent order (404 error)

## Tips

1. **Use Collection Runner**: Run all requests in sequence
   - Click on collection name
   - Click "Run" button
   - Select requests to run
   - Click "Run PostgreSQL CRUD API"

2. **Save Responses**: Right-click on response → Save Response → Save as example

3. **Create Tests**: Add test scripts to validate responses automatically

4. **Use Pre-request Scripts**: Set up data before requests

## Troubleshooting

### Collection Not Importing
- Ensure you're using Postman v8.0 or higher
- Check that JSON files are valid (no syntax errors)

### Requests Failing
- Verify API server is running: `curl http://localhost:8080/health`
- Check environment variables are set correctly
- Ensure PostgreSQL is running and accessible

### Order ID Not Updating
- Check that test scripts are enabled in Postman settings
- Manually set `order_id` in environment if needed

## Alternative: Using cURL

If you prefer command line, see `API.md` for cURL examples.

## Support

For API documentation, see:
- `API.md` - Detailed API documentation
- `README.md` - Project overview and setup

