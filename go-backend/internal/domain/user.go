package domain

import (
	"errors"
	"time"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain/commons"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound      = errors.New("usuário não encontrado")
	ErrUserAlreadyExists = errors.New("e-mail já cadastrado")
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewUser(name, email, password string, isActive bool) (*User, error) {
	nameToUpper, err := commons.IsValidName(name)
	if err != nil {
		return nil, err
	}

	if err := commons.IsValidEmail(email); err != nil {
		return nil, err
	}

	if err := commons.IsValidPassword(password); err != nil {
		return nil, err
	}

	hashedPassword, err := commons.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:        uuid.New().String(),
		Name:      nameToUpper,
		Email:     email,
		Password:  hashedPassword,
		IsActive:  isActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	return user, nil
}
