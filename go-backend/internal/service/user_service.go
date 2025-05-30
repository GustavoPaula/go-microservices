package service

import (
	"context"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/repository"
)

type UserService_impl struct {
	repository repository.UserRepository_i
}

func NewUserService(repository repository.UserRepository_i) *UserService_impl {
	return &UserService_impl{repository: repository}
}

func (s *UserService_impl) CreateUser(ctx context.Context, input dto.CreateUserInput) (*dto.UserOutput, error) {
	user, err := dto.ToUser(input)

	if err != nil {
		return nil, err
	}

	err = s.repository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(user)
	return &output, nil
}
