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

	newUser, err := s.repository.Create(ctx, user.Name, user.Email, user.Password, user.IsActive, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(newUser)
	return &output, nil
}
