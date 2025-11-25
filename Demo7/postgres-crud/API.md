# REST API Documentation

## Base URL
```
http://localhost:8080/api/v1
```

## Endpoints

### Health Check
- **GET** `/health`
- Returns server status

**Response:**
```json
{
  "status": "ok",
  "message": "Server is running"
}
```

---

### Orders

#### Create Order
- **POST** `/orders`
- Creates a new order

**Request Body:**
```json
{
  "description": "Laptop - Gaming"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "description": "Laptop - Gaming",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

**Validation:**
- `description`: Required, min 3 characters, max 255 characters

---

#### Get All Orders
- **GET** `/orders`
- Retrieves all orders

**Response (200 OK):**
```json
{
  "orders": [
    {
      "id": 1,
      "description": "Laptop - Gaming",
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-01T12:00:00Z"
    }
  ],
  "count": 1
}
```

---

#### Get Order by ID
- **GET** `/orders/:id`
- Retrieves a specific order by ID

**Response (200 OK):**
```json
{
  "id": 1,
  "description": "Laptop - Gaming",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

**Error (404 Not Found):**
```json
{
  "error": "Order not found",
  "code": 404
}
```

---

#### Update Order
- **PUT** `/orders/:id`
- Updates an existing order

**Request Body:**
```json
{
  "description": "Updated Description"
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "description": "Updated Description",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:30:00Z"
}
```

---

#### Delete Order
- **DELETE** `/orders/:id`
- Deletes an order by ID

**Response (200 OK):**
```json
{
  "message": "Order deleted successfully"
}
```

---

## Error Responses

All errors follow this format:

```json
{
  "error": "Error message",
  "details": "Additional error details (optional)",
  "code": 400
}
```

### Common Status Codes
- `200` - Success
- `201` - Created
- `400` - Bad Request (validation errors)
- `404` - Not Found
- `500` - Internal Server Error

---

## Example Requests

### Using cURL

**Create Order:**
```bash
curl -X POST http://localhost:8080/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{"description": "New Order"}'
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
  -d '{"description": "Updated Order"}'
```

**Delete Order:**
```bash
curl -X DELETE http://localhost:8080/api/v1/orders/1
```

### Using HTTPie

**Create Order:**
```bash
http POST localhost:8080/api/v1/orders description="New Order"
```

**Get All Orders:**
```bash
http GET localhost:8080/api/v1/orders
```

**Update Order:**
```bash
http PUT localhost:8080/api/v1/orders/1 description="Updated Order"
```

**Delete Order:**
```bash
http DELETE localhost:8080/api/v1/orders/1
```

---

## CORS

The API supports CORS and allows requests from any origin. The following headers are set:
- `Access-Control-Allow-Origin: *`
- `Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS, PATCH`
- `Access-Control-Allow-Headers: Content-Type, Authorization, X-Requested-With`

