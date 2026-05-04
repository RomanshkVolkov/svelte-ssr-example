package http

import (
	"net/http"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/adapters/handler"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/domain"
)

func InitRoleRoutes(roleHandler handler.RoleHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", roleHandler.GetRoles)

	return mux
}

func GetRoleRoutes() []domain.RouteInfo {
	return []domain.RouteInfo{
		{
			Method:      "GET",
			Path:        "/roles",
			Description: "Retrieve all roles",
			Handler:     "GetRoles",
		},
	}
}
