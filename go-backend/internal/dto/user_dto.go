package dto

import (
	"time"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active,omitempty"`
}

type UserOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToUser(input CreateUserInput) (*domain.User, error) {
	return domain.NewUser(input.Name, input.Email, input.Password, input.IsActive)
}

func FromUser(user *domain.User) UserOutput {
	return UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
