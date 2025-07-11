package dto

import (
	"time"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
}

type UserOutput struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
	DeletedAt time.Time `json:"deleted_at,omitzero"`
}

func ToUser(input CreateUserInput) (*domain.User, error) {
	if !input.IsActive {
		input.IsActive = true
	}

	return domain.NewUser(input.Name, input.Email, input.Password, input.IsActive)
}

func FromUser(user *domain.User) UserOutput {
	return UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func FromUsers(users []*domain.User) []UserOutput {
	output := make([]UserOutput, 0, len(users))
	for _, user := range users {
		output = append(output, FromUser(user))
	}
	return output
}
