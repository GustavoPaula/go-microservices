package service

import (
	"context"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/repository"
)

type UserService_impl struct {
	repository repository.UserRepository_i
}

func NewUserService(repository repository.UserRepository_i) *UserService_impl {
	return &UserService_impl{repository: repository}
}

func (s *UserService_impl) Create(ctx context.Context, input dto.CreateUserInput) (*dto.UserOutput, error) {
	user, err := dto.ToUser(input)

	if err != nil {
		return nil, err
	}

	existingUser, _ := s.repository.FindByEmail(ctx, user.Email)

	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	err = s.repository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(user)
	return &output, nil
}

func (s *UserService_impl) GetById(ctx context.Context, id string) (*dto.UserOutput, error) {
	user, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(user)
	return &output, nil
}

func (s *UserService_impl) GetByEmail(ctx context.Context, email string) (*dto.UserOutput, error) {
	user, err := s.repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(user)
	return &output, nil
}

func (s *UserService_impl) Update(ctx context.Context, input dto.CreateUserInput, id string) (*dto.UserOutput, error) {
	user, err := dto.ToUser(input)

	if err != nil {
		return nil, err
	}

	err = s.repository.Update(ctx, user, id)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(user)
	return &output, nil
}
