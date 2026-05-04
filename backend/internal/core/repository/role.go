package repository

import (
	"context"
	"fmt"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/domain"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRoles(ctx context.Context) ([]domain.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) GetRoles(ctx context.Context) ([]domain.Role, error) {
	roles := []domain.Role{}

	err := r.db.Find(&roles).Error

	if err != nil {
		return nil, err
	}

	fmt.Println(roles)
	return roles, nil
}
