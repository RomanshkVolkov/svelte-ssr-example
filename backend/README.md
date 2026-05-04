# Backend - Clean Architecture con Inyección de Dependencias

Este proyecto implementa una API REST siguiendo los principios de **Clean Architecture** y **Dependency Injection** usando Go. El ejemplo principal muestra el flujo completo del método `GetUsers` desde la ruta HTTP hasta la capa de datos.

## 🏗️ Arquitectura del Proyecto

### Estructura de Directorios

```
backend/
├── cmd/
│   └── main.go                 # Punto de entrada de la aplicación
├── internal/
│   ├── adapters/              # Capa de adaptadores (externa)
│   │   ├── handler/           # Controladores HTTP
│   │   │   └── user.go
│   │   ├── http/              # Configuración de rutas
│   │   │   ├── routes.go
│   │   │   └── user.go
│   │   └── middleware/        # Middlewares HTTP
│   │       └── middleware.go
│   └── core/                  # Núcleo de la aplicación (interna)
│       ├── domain/            # Entidades y modelos de dominio
│       │   └── user.go
│       ├── service/           # Lógica de negocio
│       │   └── user.go
│       └── repository/        # Interfaces y contratos de datos
│           └── user.go
├── bin/                       # Binarios compilados
├── tmp/                       # Archivos temporales (hot reload)
├── go.mod
├── go.sum
├── .air.toml                  # Configuración para hot reload
├── test_api.sh               # Script de pruebas
├── API_EXAMPLES.md           # Ejemplos de uso de la API
└── README.md
```

## 🔄 Flujo de Inyección de Dependencias

### Patrón Implementado

El proyecto sigue el patrón de **Dependency Injection** donde cada capa depende de abstracciones (interfaces) y no de implementaciones concretas.

```go
// 1. Repository Layer (Capa más externa - Infraestructura)
userRepository := repository.NewUserRepository()

// 2. Service Layer (Lógica de negocio) - Inyecta repository
userService := service.NewUserService(userRepository)

// 3. Handler Layer (Controladores HTTP) - Inyecta service  
userHandler := handler.NewUserHandler(userService)

// 4. Routes Layer - Inyecta handlers
routes := InitUserRoutes(userHandler)
```

### Flujo Completo: GET /api/v1/users

```
1. Cliente HTTP
   ↓
2. Middleware Chain (Logger, CORS, Recovery)
   ↓
3. HTTP Router (/api/v1/users)
   ↓
4. UserHandler.GetUsers() [adapters/handler/user.go]
   ├── Validaciones HTTP
   ├── Extracción de parámetros
   └── Llamada al Service
   ↓
5. UserService.GetUsers() [core/service/user.go]
   ├── Validaciones de negocio
   ├── Lógica de negocio
   └── Llamada al Repository
   ↓
6. UserRepository.GetUsers() [core/repository/user.go]
   ├── Acceso a datos (DB/Cache)
   └── Retorno de datos
   ↓
7. Respuesta JSON al cliente
```

## 🏛️ Capas de la Arquitectura

### 1. Domain Layer (Núcleo)

**Ubicación**: `internal/core/domain/`

Define las entidades y modelos de negocio. No depende de ninguna otra capa.

```go
type User struct {
    ID        string    `json:"id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### 2. Repository Layer (Contratos de Datos)

**Ubicación**: `internal/core/repository/`

Define las interfaces para el acceso a datos. Parte del núcleo de la aplicación.

```go
type UserRepository interface {
    GetUsers(ctx context.Context) ([]domain.User, error)
    GetUserByID(ctx context.Context, id string) (*domain.User, error)
    CreateUser(ctx context.Context, user *domain.User) error
    // ... otros métodos
}
```

**Principios aplicados**:
- **Dependency Inversion**: Service depende de la interfaz, no de la implementación
- **Interface Segregation**: Interfaz específica para operaciones de usuarios

### 3. Service Layer (Lógica de Negocio)

**Ubicación**: `internal/core/service/`

Contiene toda la lógica de negocio. Depende solo de interfaces del repository.

```go
type userService struct {
    userRepository repository.UserRepository // ← Inyección de dependencia
}

func NewUserService(userRepository repository.UserRepository) UserService {
    return &userService{
        userRepository: userRepository,
    }
}

func (s *userService) GetUsers(ctx context.Context) ([]domain.User, error) {
    // Validaciones de negocio
    // Lógica específica del dominio
    users, err := s.userRepository.GetUsers(ctx)
    // Procesamiento adicional si es necesario
    return users, err
}
```

**Responsabilidades**:
- Validaciones de negocio
- Orquestación de operaciones
- Transformaciones de datos
- Aplicación de reglas de dominio

### 4. Handler Layer (Controladores HTTP)

**Ubicación**: `internal/adapters/handler/`

Maneja las peticiones HTTP y las convierte en llamadas al dominio.

```go
type userHandler struct {
    userService service.UserService // ← Inyección de dependencia
}

func NewUserHandler(userService service.UserService) UserHandler {
    return &userHandler{
        userService: userService,
    }
}

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    // 1. Extraer parámetros del request
    // 2. Llamar al service
    users, err := h.userService.GetUsers(r.Context())
    // 3. Manejar errores
    // 4. Formatear respuesta JSON
    h.sendSuccessResponse(w, http.StatusOK, "Users retrieved successfully", users)
}
```

**Responsabilidades**:
- Manejo de HTTP requests/responses
- Validación de entrada HTTP
- Serialización/Deserialización JSON
- Manejo de códigos de estado HTTP

### 5. HTTP Layer (Enrutamiento)

**Ubicación**: `internal/adapters/http/`

Configura las rutas HTTP y conecta con los handlers.

```go
func InitUserRoutes(userHandler handler.UserHandler) *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /users", userHandler.GetUsers)
    mux.HandleFunc("GET /users/", userHandler.GetUserByID)
    mux.HandleFunc("POST /users", userHandler.CreateUser)
    mux.HandleFunc("PUT /users/", userHandler.UpdateUser)
    mux.HandleFunc("DELETE /users/", userHandler.DeleteUser)
    return mux
}
```

## 🔧 Principios SOLID Aplicados

### 1. Single Responsibility Principle (SRP)
- **Handler**: Solo maneja HTTP
- **Service**: Solo lógica de negocio  
- **Repository**: Solo acceso a datos
- **Domain**: Solo definiciones de entidades

### 2. Open/Closed Principle (OCP)
- Nuevas implementaciones de Repository sin modificar Service
- Nuevos handlers sin modificar Service

### 3. Liskov Substitution Principle (LSP)
- Cualquier implementación de `UserRepository` puede substituir a otra
- Cualquier implementación de `UserService` puede substituir a otra

### 4. Interface Segregation Principle (ISP)
- Interfaces específicas por dominio (UserRepository, UserService)
- No interfaces monolíticas

### 5. Dependency Inversion Principle (DIP)
- Capas altas no dependen de capas bajas
- Service depende de interfaz Repository, no de implementación
- Handler depende de interfaz Service, no de implementación

## 🚀 Cómo Ejecutar

### Requisitos
- Go 1.24.5 o superior
- curl (para testing)
- jq (opcional, para formatear JSON)

### Compilar y Ejecutar

```bash
# Compilar
cd backend
go build -o bin/server ./cmd/main.go

# Ejecutar
./bin/server

# O ejecutar directamente
go run ./cmd/main.go

# Ejecutar en puerto personalizado
PORT=8081 go run ./cmd/main.go
```

### Hot Reload (Desarrollo)

```bash
# Instalar air
go install github.com/air-verse/air@latest

# Ejecutar con hot reload
air
```

## 🧪 Testing

### Script de Pruebas Automatizadas

```bash
# Ejecutar todas las pruebas
./test_api.sh
```

### Pruebas Manuales con curl

```bash
# Health check
curl http://localhost:8081/health

# Obtener todos los usuarios
curl http://localhost:8081/api/v1/users

# Obtener usuario por ID
curl http://localhost:8081/api/v1/users/1

# Crear nuevo usuario
curl -X POST http://localhost:8081/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123", 
    "name": "Test User"
  }'
```

## 📊 Middlewares Implementados

### 1. Logger Middleware
Registra todas las peticiones HTTP con información detallada:

```
[GET] /api/v1/users 127.0.0.1:54321 - 200 - 1.234567ms - curl/7.68.0
```

### 2. CORS Middleware
Configura headers para Cross-Origin Resource Sharing:
- `Access-Control-Allow-Origin: *`
- `Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS`
- `Access-Control-Allow-Headers: Content-Type, Authorization, X-Requested-With`

### 3. Recovery Middleware
Captura panics y los convierte en errores HTTP 500:
```json
{
  "success": false,
  "message": "Internal server error",
  "error": "An unexpected error occurred"
}
```

## 🔄 Extensibilidad

### Agregar Nuevas Entidades

1. **Domain**: Crear entidad en `internal/core/domain/`
2. **Repository**: Crear interfaz en `internal/core/repository/`
3. **Service**: Implementar lógica en `internal/core/service/`
4. **Handler**: Crear controladores en `internal/adapters/handler/`
5. **Routes**: Configurar rutas en `internal/adapters/http/`
6. **Main**: Registrar en cadena de inyección

### Cambiar Base de Datos

Solo necesitas cambiar la implementación del Repository:

```go
// PostgreSQL Repository
func NewUserRepository(db *sql.DB) UserRepository {
    return &postgresUserRepository{db: db}
}

// MongoDB Repository  
func NewUserRepository(client *mongo.Client) UserRepository {
    return &mongoUserRepository{client: client}
}
```

El resto de capas permanece sin cambios.

## 📝 Estructura de Respuestas

### Respuesta Exitosa
```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": { ... }
}
```

### Respuesta de Error
```json
{
  "success": false,
  "message": "Operation failed", 
  "error": "Detailed error message"
}
```

## 🔒 Seguridad

### Headers de Seguridad
- `X-Content-Type-Options: nosniff`
- `X-Frame-Options: DENY`
- `X-XSS-Protection: 1; mode=block`
- `Strict-Transport-Security: max-age=31536000`

### Validaciones
- Validación de entrada en handlers
- Sanitización de datos
- Validación de formato de email
- Longitud mínima de contraseñas

## 📈 Próximos Pasos

1. **Base de Datos Real**: PostgreSQL/MySQL con GORM
2. **Autenticación**: JWT tokens y middleware de auth
3. **Validación Avanzada**: Librería `validator` 
4. **Testing**: Unit tests para cada capa
5. **Documentación**: Swagger/OpenAPI
6. **Rate Limiting**: Redis para límites de peticiones
7. **Caching**: Cache en memoria o Redis
8. **Logging**: Structured logging con slog
9. **Monitoring**: Métricas con Prometheus
10. **Docker**: Containerización completa

## 💡 Beneficios de esta Arquitectura

### ✅ Mantenibilidad
- Código organizado por responsabilidades
- Fácil localización de bugs
- Cambios aislados por capas

### ✅ Testabilidad  
- Inyección de dependencias facilita mocking
- Tests unitarios independientes por capa
- Tests de integración controlados

### ✅ Escalabilidad
- Fácil agregar nuevas funcionalidades
- Cambio de implementaciones sin afectar otras capas
- Separación clara de concerns

### ✅ Reutilización
- Services pueden ser usados por diferentes adapters
- Repositories pueden ser compartidos entre services
- Domain models reutilizables

Esta arquitectura proporciona una base sólida para aplicaciones Go empresariales, siguiendo las mejores prácticas de la industria y los principios de Clean Architecture.