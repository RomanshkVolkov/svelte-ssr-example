package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/domain"
)

func ValidateRequest[T any](r *http.Request) (*T, error) {
	var request T

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return &request, nil
}

// sendSuccessResponse envía una respuesta exitosa en formato JSON
func SendSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := domain.APIResponse[any]{
		Success: true,
		Message: domain.Message{
			En: "Completed action successfully",
			Es: "Se completo la accion correctamente",
		},
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}

// sendErrorResponse envía una respuesta de error en formato JSON
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string, error string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := domain.APIResponse[any]{
		Success: false,
		Message: domain.Message{
			En: message,
			Es: message,
		},
		Error: error,
	}

	json.NewEncoder(w).Encode(response)
}
