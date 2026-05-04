package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/domain"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/domain/schema"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/lg"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/repository"
	"github.com/nrednav/cuid2"
)

// UserService define la interfaz para la lógica de negocio de usuarios
type UserService interface {
	GetUsers(ctx context.Context) (domain.APIResponse[[]domain.User], error)
	GetUserByID(ctx context.Context, id string, fieldsParam string) (domain.APIResponse[map[string]any], error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, req *domain.CreateUser) (domain.APIResponse[*domain.User], error)
	UpdateUser(ctx context.Context, id string, req *domain.UpdateUserRequest) (domain.APIResponse[*domain.User], error)
	DeleteUser(ctx context.Context, id string) domain.APIResponse[*domain.User]
}

// userService implementa UserService
type userService struct {
	userRepository repository.UserRepository
}

// NewUserService crea una nueva instancia del servicio con inyección de dependencias
func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// GetUsers obtiene todos los usuarios aplicando lógica de negocio
func (s *userService) GetUsers(ctx context.Context) (domain.APIResponse[[]domain.User], error) {
	// Validaciones de negocio pueden ir aquí
	// Por ejemplo: verificar permisos, filtros, etc.

	users, err := s.userRepository.GetUsers(ctx)
	if err != nil {
		return domain.APIResponse[[]domain.User]{
			Message: domain.Message{
				Es: "Error al obtener usuarios",
				En: "Error retrieving users",
			},
			Data: []domain.User{},
		}, fmt.Errorf("error retrieving users: %w", err)
	}

	// Aquí puedes aplicar lógica de negocio adicional
	// Por ejemplo: filtrar usuarios inactivos, aplicar transformaciones, etc.

	return domain.APIResponse[[]domain.User]{
		Success: true,
		Message: domain.Message{
			Es: "Lista de usuarios",
			En: "Users list",
		},
		Data: users,
	}, nil
}

// GetUserByID obtiene un usuario por ID con validaciones de negocio
func (s *userService) GetUserByID(ctx context.Context, id string, fieldsParam string) (domain.APIResponse[map[string]any], error) {
	notFoundError := domain.APIResponse[map[string]any]{
		Success: false,
		Message: domain.Message{
			En: "user not found",
			Es: "usuario no encontrado",
		},
	}

	if id == "" {
		return notFoundError, errors.New("user ID is required")
	}

	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return domain.APIResponse[map[string]any]{
			Success: false,
			Message: domain.Message{
				Es: "Error obteniendo usuario",
				En: "Error retrieving user",
			},
		}, fmt.Errorf("error retrieving user with ID %s: %w", id, err)
	}

	if user.ID == "" {
		return notFoundError, fmt.Errorf("user with ID %s not found", id)
	}

	userWithFieldDefinition := map[string]any{}

	selectedFields := strings.Split(fieldsParam, ",")
	lg.Info("Debuggin" + strings.Join(selectedFields, " - "))
	if fieldsParam == "" {
		selectedFields = []string{"ID", "Email", "Name", "Role"}
	}

	for _, field := range selectedFields {

		switch field {
		case "ID":
			userWithFieldDefinition[strings.ToLower(field)] = user.ID
		case "CreatedAt":
			userWithFieldDefinition[strings.ToLower(field)] = user.CreatedAt
		case "UpdatedAt":
			userWithFieldDefinition[strings.ToLower(field)] = user.UpdatedAt
		case "Email":
			userWithFieldDefinition[strings.ToLower(field)] = user.Email
		case "Name":
			userWithFieldDefinition[strings.ToLower(field)] = user.Name
		case "Role":
			userWithFieldDefinition[strings.ToLower(field)] = user.Role
		}

	}

	return domain.APIResponse[map[string]any]{
		Success: true,
		Message: domain.Message{
			Es: "Usuario obtenido",
			En: "User retrieved",
		},
		Data: userWithFieldDefinition,
	}, nil
}

// GetUserByEmail obtiene un usuario por email con validaciones
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user with email %s: %w", email, err)
	}

	return user, nil
}

// CreateUser crea un nuevo usuario aplicando reglas de negocio
func (s *userService) CreateUser(ctx context.Context, req *domain.CreateUser) (domain.APIResponse[*domain.User], error) {
	fields := schema.GenericForm[domain.CreateUser]{Data: *req}
	failValidatedFields := schema.FormValidator(fields)

	if len(failValidatedFields) > 0 {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				En: "Check the fields",
				Es: "Verifica los campos",
			},
			Data:        nil,
			SchemaError: failValidatedFields,
		}, nil
	}

	// Validar que el email no esté ya registrado
	existingUser, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				En: "Error checking existing user",
				Es: "Error verificando usuario existente",
			},
			Data: nil,
		}, nil
	}

	if existingUser != nil {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				En: "User with this email already exists",
				Es: "Usuario con este email ya existe",
			},
			Data: nil,
			SchemaError: map[string][]string{
				"email": {"Email already exists"},
			},
		}, nil
	}

	// Crear el usuario
	password, err := repository.HashPassword(fields.Data.Password)
	if err != nil {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				En: "Error hashing password",
				Es: "Error hashando contraseña",
			},
			Data: nil,
		}, nil
	}
	user := &domain.User{
		BaseModel: domain.BaseModel{
			ID: cuid2.Generate(),
		},
		Email:    fields.Data.Email,
		Name:     fields.Data.Name,
		Password: password,
		RoleID:   fields.Data.RoleID,
	}

	createdUser, createdIDerr := s.userRepository.CreateUser(ctx, user)
	if createdIDerr != nil {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				En: "Error creating user",
				Es: "Error creando usuario",
			},
			Data: createdUser,
		}, nil
	}

	return domain.APIResponse[*domain.User]{
		Success: true,
		Message: domain.Message{
			En: "User created successfully",
			Es: "Usuario creado exitosamente",
		},
		Data: user,
	}, nil
}

// actualiza un usuario existente
func (s *userService) UpdateUser(ctx context.Context, id string, req *domain.UpdateUserRequest) (domain.APIResponse[*domain.User], error) {
	if id == "" {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				En: "User ID is required",
				Es: "ID de usuario es requerido",
			},
			Data: nil,
		}, errors.New("user ID is required")
	}

	if req == nil {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				En: "Update user request is required",
				Es: "Solicitud de actualización de usuario es requerida",
			},
			Data: nil,
		}, errors.New("update user request is required")
	}

	// Verificar que el usuario existe
	existingUser, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				En: "User not found",
				Es: "Usuario no encontrado",
			},
			Data: nil,
		}, err
	}

	// Si se está actualizando el email, verificar que no exista otro usuario con ese email
	if req.Email != "" && req.Email != existingUser.Email {
		userWithEmail, err := s.userRepository.GetUserByEmail(ctx, req.Email)
		if err != nil {
			return domain.APIResponse[*domain.User]{
				Success: false,
				Message: domain.Message{
					En: "Error checking email availability",
					Es: "Error al verificar disponibilidad del email",
				},
				Data: nil,
			}, fmt.Errorf("error checking email availability: %w", err)
		}
		if userWithEmail != nil {
			return domain.APIResponse[*domain.User]{
				Success: false,
				Message: domain.Message{
					En: "User with this email already exists",
					Es: "User with this email already exists",
				},
				Data: nil,
			}, errors.New("email is already in use by another user")
		}
	}

	// Actualizar campos
	if req.Email != "" {
		existingUser.Email = req.Email
	}
	if req.Name != "" {
		existingUser.Name = req.Name
	}
	existingUser.UpdatedAt = time.Now()

	user, err := s.userRepository.UpdateUser(ctx, id, existingUser)
	if err != nil {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				En: "Error updating user",
				Es: "Error al actualizar usuario",
			},
			Data: nil,
		}, fmt.Errorf("error updating user: %w", err)
	}

	return domain.APIResponse[*domain.User]{
		Success: true,
		Message: domain.Message{
			Es: "Usuario actualizado correctamente",
			En: "User updated successfully",
		},
		Data: user,
	}, nil
}

// DeleteUser elimina un usuario
func (s *userService) DeleteUser(ctx context.Context, id string) domain.APIResponse[*domain.User] {
	if id == "" {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				En: "please provide user",
				Es: "por favor especifique un usuario",
			},
			Data: nil,
		}
	}

	// Verificar que el usuario existe antes de eliminarlo
	_, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return domain.APIResponse[*domain.User]{
			Success: false,
			Message: domain.Message{
				Es: "user not found",
				En: "usuario no encontrado",
			},
			Data: nil,
		}
	}

	err = s.userRepository.DeleteUser(ctx, id)
	if err == nil {
		return domain.APIResponse[*domain.User]{
			Success: true,
			Message: domain.Message{
				En: "user deleted successfully",
				Es: "usuario eliminado correctamente",
			},
			Data: nil,
		}
	}

	return domain.APIResponse[*domain.User]{
		Success: false,
		Message: domain.Message{
			En: "internal server error",
			Es: "error interno de servidor",
		},
		Data: nil,
	}
}

// Funciones helper (en una implementación real estas estarían en un paquete utils)
func generateID() string {
	// Implementar generación de ID (UUID, etc.)
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
