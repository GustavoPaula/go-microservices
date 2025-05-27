package domain

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"

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
	if err := isValidEmail(email); err != nil {
		return nil, err
	}

	if err := isValidPassword(password); err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(password)
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

func isValidEmail(email string) error {
	regex := `(?i)^(?:[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~.-]+)@(?:[a-z0-9-]+\.)+[a-z]{2,}$`
	matched, err := regexp.MatchString(regex, email)
	if err != nil {
		return fmt.Errorf("erro ao validar e-mail: %v", err)
	}

	if !matched {
		return errors.New("e-mail inválido")
	}

	return nil
}

func isValidPassword(password string) error {
	regex := `^[^\s]{6,}$`
	matched, err := regexp.MatchString(regex, password)
	if err != nil {
		return fmt.Errorf("erro ao validar senha: %v", err)
	}

	if !matched {
		return errors.New("senha inválida")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
