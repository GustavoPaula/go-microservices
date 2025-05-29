package commons

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidEmail    = errors.New("e-mail inválido")
	ErrInvalidPassword = errors.New("senha inválida")
)

func IsValidEmail(email string) error {
	regex := `(?i)^(?:[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~.-]+)@(?:[a-z0-9-]+\.)+[a-z]{2,}$`
	matched, err := regexp.MatchString(regex, email)
	if err != nil {
		fmt.Printf("erro ao validar e-mail: %v", err)
		return err
	}

	if !matched {
		return ErrInvalidEmail
	}

	return nil
}

func IsValidPassword(password string) error {
	regex := `^[^\s]{6,}$`
	matched, err := regexp.MatchString(regex, password)
	if err != nil {
		fmt.Printf("erro ao validar senha: %v", err)
		return err
	}

	if !matched {
		return ErrInvalidPassword
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Printf("erro ao criptografar a senha: %v", err)
		return "", err
	}
	return string(bytes), nil
}
