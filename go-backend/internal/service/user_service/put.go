package user_service

import (
	"context"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) Put(ctx context.Context, input dto.CreateUserInput, id string) (*dto.UserOutput, error) {
	user, err := dto.ToUser(input)

	if err != nil {
		return nil, err
	}

	err = s.repository.Put(ctx, user, id)
	if err != nil {
		return nil, err
	}

	userUpdated, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(userUpdated)
	return &output, nil
}
