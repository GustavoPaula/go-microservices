package user_service

import (
	"context"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) Create(ctx context.Context, input dto.CreateUserInput) (*dto.UserOutput, error) {
	user, err := dto.ToUser(input)

	if err != nil {
		return nil, err
	}

	existingUser, _ := s.repository.GetByEmail(ctx, user.Email)

	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	err = s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	newUser, err := s.repository.GetById(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(newUser)
	return &output, nil
}
