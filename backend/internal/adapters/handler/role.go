package handler

import (
	"net/http"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/service"
)

type RoleHandler interface {
	GetRoles(w http.ResponseWriter, r *http.Request)
}

type roleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) RoleHandler {
	return &roleHandler{
		roleService: roleService,
	}
}

func (h *roleHandler) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.roleService.GetRoles(r.Context())
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to get roles", err.Error())
		return
	}

	SendSuccessResponse(w, http.StatusOK, "Roles retrieved successfully", roles)
}
