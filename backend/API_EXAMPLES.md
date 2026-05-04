# API Examples - User Management

Este archivo contiene ejemplos de cómo probar todos los endpoints de la API de usuarios siguiendo el patrón de Clean Architecture con Dependency Injection.

## Flujo Completo del Request GetUsers

### 1. Cliente → HTTP Handler → Service → Repository

```
GET /api/v1/users
↓
HTTP Layer (handler/user.go) → GetUsers()
↓
Service Layer (service/user.go) → GetUsers()
↓
Repository Layer (repository/user.go) → GetUsers()
↓
Return mock data → Service → Handler → JSON Response
```

## Endpoints Disponibles

### Base URL
```
http://localhost:8081/api/v1
```

### Health Check
```bash
curl -X GET http://localhost:8081/health
```

**Response:**
```json
{
  "status": "healthy", 
  "message": "API is running"
}
```

## User Endpoints

### 1. GET /api/v1/users - Obtener todos los usuarios

```bash
curl -X GET http://localhost:8081/api/v1/users
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": "1",
      "email": "john@example.com",
      "name": "John Doe",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
    },
    {
      "id": "2",
      "email": "jane@example.com",
      "name": "Jane Smith",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
    },
    {
      "id": "3",
      "email": "bob@example.com",
      "name": "Bob Johnson",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
    }
  ]
}
```

### 2. GET /api/v1/users/{id} - Obtener usuario por ID

```bash
curl -X GET http://localhost:8081/api/v1/users/1
```

**Expected Response:**
```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "id": "1",
    "email": "john@example.com",
    "name": "John Doe",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  }
}
```

**Error Response (User not found):**
```bash
curl -X GET http://localhost:8081/api/v1/users/999
```

```json
{
  "success": false,
  "message": "User not found",
  "error": "user with ID 999 not found"
}
```

### 3. POST /api/v1/users - Crear nuevo usuario

```bash
curl -X POST http://localhost:8081/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newuser@example.com",
    "password": "password123",
    "name": "New User"
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": "1642781234567890123",
    "email": "newuser@example.com",
    "name": "New User",
    "created_at": "2024-01-21T10:30:45Z",
    "updated_at": "2024-01-21T10:30:45Z"
  }
}
```

**Error Response (Validation failed):**
```bash
curl -X POST http://localhost:8081/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "invalid-email",
    "password": "123",
    "name": ""
  }'
```

```json
{
  "success": false,
  "message": "Validation failed",
  "error": "invalid email format"
}
```

### 4. PUT /api/v1/users/{id} - Actualizar usuario

```bash
curl -X PUT http://localhost:8081/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "email": "updated@example.com",
    "name": "Updated Name"
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "message": "User updated successfully",
  "data": {
    "id": "1",
    "email": "updated@example.com",
    "name": "Updated Name",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "2024-01-21T10:35:22Z"
  }
}
```

### 5. DELETE /api/v1/users/{id} - Eliminar usuario

```bash
curl -X DELETE http://localhost:8081/api/v1/users/1
```

**Expected Response:**
```json
{
  "success": true,
  "message": "User deleted successfully"
}
```

## Flujo Detallado GetUsers - Traza del Request

### 1. Inicio del Request
```
Client → GET /api/v1/users
```

### 2. HTTP Layer (internal/adapters/handler/user.go)
```go
func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    // 1. Recibe el HTTP request
    // 2. Extrae contexto del request
    users, err := h.userService.GetUsers(r.Context())
    // 3. Maneja errores si existen
    // 4. Envía respuesta JSON formateada
}
```

### 3. Service Layer (internal/core/service/user.go)
```go
func (s *userService) GetUsers(ctx context.Context) ([]domain.User, error) {
    // 1. Aplica validaciones de negocio
    // 2. Llama al repository
    users, err := s.userRepository.GetUsers(ctx)
    // 3. Aplica lógica de negocio adicional si es necesaria
    // 4. Retorna los usuarios procesados
}
```

### 4. Repository Layer (internal/core/repository/user.go)
```go
func (r *userRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
    // 1. Acceso a datos (DB, cache, etc.)
    // 2. En este ejemplo retorna datos mock
    // 3. En producción haría query a base de datos
    users := []domain.User{...}
    return users, nil
}
```

### 5. Retorno de datos
```
Repository → Service → Handler → HTTP Response → Client
```

## Patrones Implementados

### 1. Dependency Injection
```go
// En main.go
userRepository := repository.NewUserRepository()
userService := service.NewUserService(userRepository)  // Inyecta repository
userHandler := handler.NewUserHandler(userService)     // Inyecta service
```

### 2. Interface Segregation
```go
// Cada capa expone interfaces claras
type UserRepository interface {
    GetUsers(ctx context.Context) ([]domain.User, error)
    // ... otros métodos
}

type UserService interface {
    GetUsers(ctx context.Context) ([]domain.User, error)
    // ... otros métodos
}
```

### 3. Single Responsibility
- **Handler**: Maneja HTTP requests/responses
- **Service**: Contiene lógica de negocio
- **Repository**: Maneja persistencia de datos
- **Domain**: Define entidades y contratos

### 4. Hexagonal Architecture
```
Adapters (HTTP, DB) → Core (Service, Domain) ← Ports (Interfaces)
```

## Testing con diferentes herramientas

### Con httpie
```bash
# Instalar: pip install httpie

# GET users
http GET localhost:8081/api/v1/users

# POST new user
http POST localhost:8081/api/v1/users email=test@example.com password=password123 name="Test User"

# PUT update user
http PUT localhost:8081/api/v1/users/1 email=updated@example.com name="Updated Name"

# DELETE user
http DELETE localhost:8081/api/v1/users/1
```

### Con Postman Collection
```json
{
  "info": {
    "name": "User API",
    "description": "API endpoints for user management"
  },
  "item": [
    {
      "name": "Get All Users",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{base_url}}/api/v1/users",
          "host": ["{{base_url}}"],
          "path": ["api", "v1", "users"]
        }
      }
    }
  ],
  "variable": [
    {
      "key": "base_url",
      "value": "http://localhost:8081",
      "type": "string"
    }
  ]
}
```

## Middleware en Acción

Cada request pasa por los siguientes middlewares:
1. **Logger**: Registra información del request
2. **CORS**: Maneja cross-origin requests
3. **Recovery**: Captura panics y los convierte en errores HTTP 500

## Logs del Servidor
```
[GET] /api/v1/users 127.0.0.1:54321 - 200 - 1.234567ms - curl/7.68.0
```

## Próximos Pasos

1. **Implementar Base de Datos**: Reemplazar datos mock con PostgreSQL/MySQL
2. **Agregar Autenticación**: JWT tokens, middleware de auth
3. **Implementar Validación**: Usar bibliotecas como `validator`
4. **Agregar Tests**: Unit tests para cada capa
5. **Documentación API**: Swagger/OpenAPI
6. **Rate Limiting**: Implementar límites de requests
7. **Caching**: Redis para mejorar performance
8. **Metrics**: Prometheus para monitoreo

Este ejemplo muestra un flujo completo siguiendo los principios SOLID y Clean Architecture con correcta inyección de dependencias.