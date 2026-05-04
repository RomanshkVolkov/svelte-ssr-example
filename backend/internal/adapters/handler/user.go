package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/domain"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/service"
)

// UserHandler define la interfaz para los handlers HTTP de usuarios
type UserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

// userHandler implementa UserHandler
type userHandler struct {
	userService service.UserService
}

// NewUserHandler crea una nueva instancia del handler con inyección de dependencias
func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

// GetUsers maneja la petición GET /users
func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Obtener usuarios del servicio
	users, err := h.userService.GetUsers(r.Context())
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve users", err.Error())
		return
	}

	// Respuesta exitosa
	SendSuccessResponse(w, http.StatusOK, "Users retrieved successfully", users)
}

// GetUserByID maneja la petición GET /users/{id}
func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extraer ID del path
	id := r.PathValue("id")

	fieldsParam := r.URL.Query().Get("fields")
	if id == "" {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", "User ID is required")
		return
	}

	// Obtener usuario del servicio
	user, err := h.userService.GetUserByID(r.Context(), id, fieldsParam)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			SendErrorResponse(w, http.StatusNotFound, "User not found", err.Error())
		} else {
			SendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve user", err.Error())
		}
		return
	}

	// Respuesta exitosa
	SendSuccessResponse(w, http.StatusOK, "User retrieved successfully", user)
}

// CreateUser maneja la petición POST /users
func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decodificar el request body
	req, err := ValidateRequest[domain.CreateUser](r)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Crear usuario a través del servicio
	user, err := h.userService.CreateUser(r.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			SendErrorResponse(w, http.StatusConflict, "User already exists", err.Error())
		} else {
			SendErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err.Error())
		}
		return
	}

	// Respuesta exitosa
	SendSuccessResponse(w, http.StatusCreated, "User created successfully", user)
}

// UpdateUser maneja la petición PUT /users/{id}
func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Extraer ID del path
	id := r.PathValue("id")
	if id == "" {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", "User ID is required")
		return
	}

	req, err := ValidateRequest[domain.UpdateUserRequest](r)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	// Actualizar usuario a través del servicio
	user, err := h.userService.UpdateUser(r.Context(), id, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			SendErrorResponse(w, http.StatusNotFound, "User not found", err.Error())
		} else if strings.Contains(err.Error(), "already in use") {
			SendErrorResponse(w, http.StatusConflict, "Email already in use", err.Error())
		} else {
			SendErrorResponse(w, http.StatusInternalServerError, "Failed to update user", err.Error())
		}
		return
	}

	// Respuesta exitosa
	SendSuccessResponse(w, http.StatusOK, "User updated successfully", user)
}

// DeleteUser maneja la petición DELETE /users/{id}
func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extraer ID del path
	id := r.PathValue("id")
	if id == "" {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", "User ID is required")
		return
	}

	// Eliminar usuario a través del servicio
	response := h.userService.DeleteUser(r.Context(), id)

	// respuesta procesada
	SendSuccessResponse(w, http.StatusOK, "delete user action processed", response)
}

// extractIDFromPath extrae el ID del path de la URL
// Ejemplo: /users/123 -> 123
func (h *userHandler) extractIDFromPath(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) >= 2 && parts[0] == "users" {
		return parts[1]
	}

	return ""
}

// validateCreateUserRequest valida los datos del request de creación
func (h *userHandler) validateCreateUserRequest(req *domain.CreateUser) error {
	if req.Email == "" {
		return errors.New("email is required")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	if req.Name == "" {
		return errors.New("name is required")
	}

	if len(req.Name) < 2 {
		return errors.New("name must be at least 2 characters long")
	}

	// Validación básica de email
	if !strings.Contains(req.Email, "@") {
		return errors.New("invalid email format")
	}

	return nil
}
