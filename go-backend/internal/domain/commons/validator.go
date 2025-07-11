package commons

import (
	"errors"
	"log/slog"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidName     = errors.New("nome inválido")
	ErrInvalidEmail    = errors.New("e-mail inválido")
	ErrInvalidPassword = errors.New("senha inválida")
)

func IsValidName(name string) (string, error) {
	regex := `(?i)^[\p{L} ]{3,}$`
	matched, err := regexp.MatchString(regex, name)
	if err != nil {
		slog.Error("Erro ao validar o nome", "error", err)
		return "", err
	}

	if !matched {
		return "", ErrInvalidName
	}

	nameToUpper := strings.ToUpper(name)
	return nameToUpper, nil
}

func IsValidEmail(email string) error {
	regex := `(?i)^(?:[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~.-]+)@(?:[a-z0-9-]+\.)+[a-z]{2,}$`
	matched, err := regexp.MatchString(regex, email)
	if err != nil {
		slog.Error("Erro ao validar e-mail", "error", err)
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
		slog.Error("Erro ao validar senha", "error", err)
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
		slog.Error("Erro ao criptografar a senha", "error", err)
		return "", err
	}
	return string(bytes), nil
}
