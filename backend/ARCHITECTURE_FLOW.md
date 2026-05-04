# Flujo de Inyección de Dependencias - Diagrama Visual

Este documento muestra visualmente cómo fluye una petición HTTP a través de todas las capas de la arquitectura siguiendo los principios de inyección de dependencias.

## 🎯 Flujo Completo: GET /api/v1/users

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           CLIENTE HTTP                                     │
│                     curl GET /api/v1/users                                 │
└─────────────────────────────┬───────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                        MIDDLEWARE CHAIN                                    │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────────────────┐ │
│  │     Logger      │  │      CORS       │  │         Recovery            │ │
│  │ - Log request   │→ │ - Set headers   │→ │ - Catch panics              │ │
│  │ - Track timing  │  │ - Handle OPTIONS│  │ - Convert to HTTP 500       │ │
│  └─────────────────┘  └─────────────────┘  └─────────────────────────────┘ │
└─────────────────────────────┬───────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                          HTTP ROUTER                                       │
│                    /api/v1/users → UserHandler.GetUsers                   │
└─────────────────────────────┬───────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                      HANDLER LAYER (ADAPTERS)                              │
│  📁 internal/adapters/handler/user.go                                      │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  func (h *userHandler) GetUsers(w, r)                              │   │
│  │  {                                                                  │   │
│  │    // 1. Extract HTTP parameters                                   │   │
│  │    // 2. Validate HTTP request                                     │   │
│  │    users, err := h.userService.GetUsers(r.Context()) ←─────────────┤   │
│  │    // 3. Handle business errors                                    │   │
│  │    // 4. Format JSON response                                      │   │
│  │    h.sendSuccessResponse(w, 200, "success", users)                 │   │
│  │  }                                                                  │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  💉 DEPENDENCY INJECTION:                                                  │
│      userHandler{userService: service.UserService} ←─── Inyectado         │
└─────────────────────────────────┬───────────────────────────────────────────┘
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                       SERVICE LAYER (CORE)                                 │
│  📁 internal/core/service/user.go                                          │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  func (s *userService) GetUsers(ctx context.Context)               │   │
│  │  {                                                                  │   │
│  │    // 1. Business validations                                      │   │
│  │    // 2. Apply business rules                                      │   │
│  │    users, err := s.userRepository.GetUsers(ctx) ←──────────────────┤   │
│  │    // 3. Transform data if needed                                  │   │
│  │    // 4. Apply business logic                                      │   │
│  │    return users, err                                               │   │
│  │  }                                                                  │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  💉 DEPENDENCY INJECTION:                                                  │
│      userService{userRepository: repository.UserRepository} ←─── Inyectado │
└─────────────────────────────────┬───────────────────────────────────────────┘
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                     REPOSITORY LAYER (CORE)                                │
│  📁 internal/core/repository/user.go                                       │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  func (r *userRepository) GetUsers(ctx context.Context)            │   │
│  │  {                                                                  │   │
│  │    // 1. Data access logic                                         │   │
│  │    // 2. Database queries (PostgreSQL/MySQL)                      │   │
│  │    // 3. Cache operations (Redis)                                  │   │
│  │    // 4. External API calls                                        │   │
│  │    // 5. File system operations                                    │   │
│  │    users := []domain.User{...} // Mock data for now               │   │
│  │    return users, nil                                               │   │
│  │  }                                                                  │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  💉 NO DEPENDENCIES: Esta capa es la implementación final                 │
└─────────────────────────────────┬───────────────────────────────────────────┘
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                        DOMAIN LAYER (CORE)                                 │
│  📁 internal/core/domain/user.go                                           │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  type User struct {                                                 │   │
│  │    ID        string    `json:"id"`                                  │   │
│  │    Email     string    `json:"email"`                               │   │
│  │    Name      string    `json:"name"`                                │   │
│  │    CreatedAt time.Time `json:"created_at"`                          │   │
│  │    UpdatedAt time.Time `json:"updated_at"`                          │   │
│  │  }                                                                  │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  💎 PURE ENTITIES: Sin dependencias, núcleo del dominio                   │
└─────────────────────────────────┬───────────────────────────────────────────┘
                                  │
                            RETORNO DE DATOS
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                         RESPUESTA HTTP                                     │
│  {                                                                          │
│    "success": true,                                                         │
│    "message": "Users retrieved successfully",                              │
│    "data": [                                                               │
│      {                                                                      │
│        "id": "1",                                                           │
│        "email": "john@example.com",                                         │
│        "name": "John Doe",                                                  │
│        "created_at": "0001-01-01T00:00:00Z",                               │
│        "updated_at": "0001-01-01T00:00:00Z"                                │
│      }                                                                      │
│    ]                                                                        │
│  }                                                                          │
└─────────────────────────────────────────────────────────────────────────────┘
```

## 🏗️ Cadena de Inyección de Dependencias

```
main.go
├── 1. userRepository := repository.NewUserRepository()
├── 2. userService := service.NewUserService(userRepository) ←─── 💉 Inyecta Repository
├── 3. userHandler := handler.NewUserHandler(userService)   ←─── 💉 Inyecta Service  
├── 4. userRoutes := http.InitUserRoutes(userHandler)       ←─── 💉 Inyecta Handler
└── 5. mainRouter := http.InitRoutes() ←──────────────────── ←─── 💉 Configura todas las rutas
```

### Detalles de cada Constructor:

```go
// 1. Repository (Sin dependencias - Capa más externa)
func NewUserRepository() UserRepository {
    return &userRepository{
        // db: database connection
        // cache: redis client  
        // logger: structured logger
    }
}

// 2. Service (Depende de Repository interface)
func NewUserService(repo repository.UserRepository) UserService {
    return &userService{
        userRepository: repo, // ← Inyección por constructor
    }
}

// 3. Handler (Depende de Service interface)
func NewUserHandler(svc service.UserService) UserHandler {
    return &userHandler{
        userService: svc, // ← Inyección por constructor
    }
}
```

## 🔄 Principios SOLID en Acción

### 📋 Single Responsibility Principle (SRP)
```
┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐
│     Handler     │  │     Service     │  │   Repository    │  │     Domain      │
│                 │  │                 │  │                 │  │                 │
│ • HTTP Requests │  │ • Business      │  │ • Data Access   │  │ • Pure Entities │
│ • HTTP Response │  │   Logic         │  │ • Queries       │  │ • Value Objects │
│ • JSON Marshal  │  │ • Validation    │  │ • Persistence   │  │ • Business      │
│ • Error Codes   │  │ • Orchestration │  │ • Caching       │  │   Rules         │
└─────────────────┘  └─────────────────┘  └─────────────────┘  └─────────────────┘
```

### 🔄 Dependency Inversion Principle (DIP)
```
High Level Modules (Service) ──────────→ Abstraction (Repository Interface)
                                                    ↑
                                                    │
Low Level Modules (Database Implementation) ────────┘

Service no conoce PostgreSQL, MySQL, MongoDB, etc.
Service solo conoce la interfaz UserRepository.
```

### 🔧 Open/Closed Principle (OCP)
```
                    UserRepository Interface
                            │
                            │
           ┌────────────────┼────────────────┐
           │                │                │
           ▼                ▼                ▼
   PostgreSQLRepo     MySQLRepo      MongoDBRepo
   
   Puedes agregar nuevas implementaciones sin modificar Service
```

## 🎪 Beneficios de esta Arquitectura

### ✅ Testabilidad
```go
// Test del Service usando mock repository
func TestUserService_GetUsers(t *testing.T) {
    mockRepo := &MockUserRepository{}
    mockRepo.On("GetUsers").Return(mockUsers, nil)
    
    service := NewUserService(mockRepo) // ← Inyección de mock
    
    users, err := service.GetUsers(context.Background())
    
    assert.NoError(t, err)
    assert.Len(t, users, 3)
    mockRepo.AssertExpectations(t)
}
```

### ✅ Mantenibilidad
- Cada capa tiene una responsabilidad clara
- Cambios aislados por capas
- Fácil localización de bugs

### ✅ Escalabilidad  
- Fácil agregar nuevos endpoints
- Nuevas entidades siguen el mismo patrón
- Cambio de implementaciones transparente

## 🚦 Flujo de Error Handling

```
Repository Error
       │
       ▼
Service Layer ──→ Business Logic Error
       │                    │
       ▼                    ▼
Handler Layer ──────────────────→ HTTP Error Response
       │
       ▼
{
  "success": false,
  "message": "Failed to retrieve users",
  "error": "database connection timeout"
}
```

## 🔍 Logging Flow

```
Request:  [INFO] GET /api/v1/users from 127.0.0.1
          │
Handler:  [DEBUG] Calling UserService.GetUsers()
          │
Service:  [DEBUG] Applying business validations
          │
Repository: [DEBUG] Executing query: SELECT * FROM users
          │
Repository: [DEBUG] Found 3 users
          │
Service:  [DEBUG] Business logic applied successfully  
          │
Handler:  [DEBUG] Sending JSON response
          │
Response: [INFO] 200 OK - 2.34ms - 3 users returned
```

## 🏆 Resultado Final

Esta arquitectura proporciona:

1. **🔒 Inversión de Control**: Las dependencias fluyen hacia adentro
2. **🧪 Alta Testabilidad**: Fácil inyección de mocks
3. **🔄 Bajo Acoplamiento**: Capas independientes
4. **🎯 Alta Cohesión**: Responsabilidades bien definidas
5. **📈 Escalabilidad**: Fácil extensión y mantenimiento
6. **🛡️ Robustez**: Manejo de errores en cada capa
7. **📊 Observabilidad**: Logging y métricas en cada nivel

El flujo completo desde HTTP hasta datos y vuelta garantiza que cada capa cumpla su función específica mientras mantiene la flexibilidad para cambios futuros.