package http

import (
	"net/http"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/adapters/handler"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/domain"
)

// InitUserRoutes configura todas las rutas relacionadas con usuarios
func InitUserRoutes(userHandler handler.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// GET /users - Obtener todos los usuarios
	mux.HandleFunc("GET /", userHandler.GetUsers)

	// GET /users/{id} - Obtener un usuario específico por ID
	mux.HandleFunc("GET /{id}", userHandler.GetUserByID)

	// POST /users - Crear un nuevo usuario
	mux.HandleFunc("POST /", userHandler.CreateUser)

	// PUT /users/{id} - Actualizar un usuario existente
	mux.HandleFunc("PUT /{id}", userHandler.UpdateUser)

	// DELETE /users/{id} - Eliminar un usuario
	mux.HandleFunc("DELETE /{id}", userHandler.DeleteUser)

	return mux
}

// getUserRoutes retorna un slice con información sobre las rutas disponibles
// Útil para documentación o debugging
func GetUserRoutes() []domain.RouteInfo {
	return []domain.RouteInfo{
		{
			Method:      "GET",
			Path:        "/users",
			Description: "Retrieve all users",
			Handler:     "GetUsers",
		},
		{
			Method:      "GET",
			Path:        "/users/{id}",
			Description: "Retrieve a specific user by ID",
			Handler:     "GetUserByID",
		},
		{
			Method:      "POST",
			Path:        "/users",
			Description: "Create a new user",
			Handler:     "CreateUser",
		},
		{
			Method:      "PUT",
			Path:        "/users/{id}",
			Description: "Update an existing user",
			Handler:     "UpdateUser",
		},
		{
			Method:      "DELETE",
			Path:        "/users/{id}",
			Description: "Delete a user",
			Handler:     "DeleteUser",
		},
	}
}
