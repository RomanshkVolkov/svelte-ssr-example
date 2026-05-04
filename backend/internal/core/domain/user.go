package domain

import "time"

type User struct {
	BaseModel
	Email      string    `gorm:"type:varchar(200);not null;unique;" json:"email"`
	Password   string    `gorm:"type:varchar(100);not null;" json:"-"`
	Name       string    `gorm:"type:varchar(200);not null;" json:"name"`
	OTP        string    `gorm:"type:varchar(100);" json:"otp"`
	ExpiresOTP time.Time `gorm:"default:NULL" json:"expiresOTP"`

	RoleID string `gorm:"type:varchar(24);not null;" json:"-"`
	Role   Role   `json:"role"`
}

type Role struct {
	BaseModel
	Name string `gorm:"type:varchar(200);not null;unique;" json:"name"`
}

type CreateUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required,min=2"`
	RoleID   string `json:"roleID" validate:"required"`
}

type UpdateUserRequest struct {
	Email string `json:"email,omitempty" validate:"omitempty,email"`
	Name  string `json:"name,omitempty" validate:"omitempty,min=2"`
}
