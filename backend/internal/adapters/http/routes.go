package http

import (
	"net/http"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/adapters/handler"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/repository"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/service"
	"gorm.io/gorm"
)

// InitRoutes configura todas las rutas de la aplicación
func InitRoutes(db *gorm.DB) *http.ServeMux {
	// Crear instancias siguiendo el patrón de inyección de dependencias

	// Repository Layer (capa más externa - infraestructura)
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)

	// Service Layer (lógica de negocio) - inyecta repository
	userService := service.NewUserService(userRepository)
	roleService := service.NewRoleService(roleRepository)

	// Handler Layer (controladores HTTP) - inyecta service
	userHandler := handler.NewUserHandler(userService)
	roleHandler := handler.NewRoleHandler(roleService)

	// Configurar rutas de usuarios
	userRoutes := InitUserRoutes(userHandler)
	roleRoutes := InitRoleRoutes(roleHandler)

	// Crear el multiplexor principal
	mainMux := http.NewServeMux()

	// Montar las rutas de usuarios bajo el path /api/v1
	mainMux.Handle("/api/v1/users/", http.StripPrefix("/api/v1/users", userRoutes))
	mainMux.Handle("/api/v1/roles/", http.StripPrefix("/api/v1/roles", roleRoutes))

	// Ruta de health check
	mainMux.HandleFunc("GET /health", healthCheckHandler)

	// Ruta para documentación de API
	mainMux.HandleFunc("GET /api/routes", apiRoutesHandler)

	return mainMux
}

// healthCheckHandler maneja el endpoint de health check
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy", "message": "API is running"}`))
}

// apiRoutesHandler maneja el endpoint de documentación de rutas
func apiRoutesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// En una implementación real, aquí podrías generar documentación automática
	response := `{
		"message": "Available API routes",
		"base_url": "/api/v1",
		"routes": {
			"users": {
				"GET /users": "Get all users",
				"GET /users/{id}": "Get user by ID",
				"POST /users": "Create new user",
				"PUT /users/{id}": "Update user",
				"DELETE /users/{id}": "Delete user"
			},
			"health": {
				"GET /health": "Health check endpoint"
			}
		}
	}`

	w.Write([]byte(response))
}
