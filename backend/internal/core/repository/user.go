package repository

import (
	"context"
	"time"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/domain"
	"github.com/nrednav/cuid2"
	"gorm.io/gorm"
)

// UserRepository define la interfaz para las operaciones de persistencia de usuarios
type UserRepository interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id string) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, id string, user domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
}

// userRepository implementa UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository crea una nueva instancia del repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// GetUsers obtiene todos los usuarios
func (r *userRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	users := []domain.User{}

	err := r.db.Preload("Role").Model(&domain.User{}).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID obtiene un usuario por ID
func (r *userRepository) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	user := domain.User{}
	err := r.db.Preload("Role").Where("id = ?", id).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}

	return user, nil // Usuario no encontrado
}

// GetUserByEmail obtiene un usuario por email
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	// Simulamos búsqueda por email
	users, err := r.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, nil // Usuario no encontrado
}

// CreateUser crea un nuevo usuario
func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUser actualiza un usuario existente
func (r *userRepository) UpdateUser(ctx context.Context, id string, user domain.User) (*domain.User, error) {

	if err := r.db.Model(&domain.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser elimina un usuario
func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	user := domain.User{}
	err := r.db.Delete(&user, "id = ?", id).Error
	return err
}

func SeedUsers(db *gorm.DB) error {
	var currentRows int64
	db.Model(&domain.User{}).Count(&currentRows)

	if currentRows > 0 {
		return nil
	}

	now := time.Now()
	users := []domain.User{
		domain.User{
			BaseModel: domain.BaseModel{
				ID:        cuid2.Generate(),
				CreatedAt: now,
				UpdatedAt: now,
			},
			Email:    "jose@guz-studio.dev",
			Name:     "Romanshk Volkov",
			Password: cuid2.Generate(),
			RoleID:   "qqovic77d7bz181zv3l00kq6",
		},
	}

	for _, user := range users {
		password, err := HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = password
		db.Create(&user)
	}

	return nil
}
