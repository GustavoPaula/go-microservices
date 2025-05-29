package domain

import (
	"fmt"
	"time"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain/commons"
	"github.com/google/uuid"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, email, password string) (*User, error) {
	if err := commons.IsValidEmail(email); err != nil {
		return nil, err
	}

	if err := commons.IsValidPassword(password); err != nil {
		return nil, err
	}

	hashedPassword, err := commons.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar hash da senha: %v", err)
	}

	user := &User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}
