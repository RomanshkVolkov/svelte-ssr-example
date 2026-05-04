package service

import (
	"context"
	"fmt"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/domain"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/repository"
)

type RoleService interface {
	GetRoles(ctx context.Context) (domain.APIResponse[[]domain.Role], error)
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return &roleService{
		roleRepository: roleRepository,
	}
}

func (s *roleService) GetRoles(ctx context.Context) (domain.APIResponse[[]domain.Role], error) {
	roles, err := s.roleRepository.GetRoles(ctx)
	if err != nil {
		return domain.APIResponse[[]domain.Role]{
			Message: domain.Message{
				Es: "Error al obtener la lista de roles",
				En: "Error retrieving user list",
			},
			Data: []domain.Role{},
		}, fmt.Errorf("error retrieving roles: %w", err)
	}

	return domain.APIResponse[[]domain.Role]{
		Success: true,
		Message: domain.Message{
			Es: "Lista de roles",
			En: "Role list",
		},
		Data: roles,
	}, nil
}
