package user_service

import (
	"context"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) GetById(ctx context.Context, id string) (*dto.UserOutput, error) {
	user, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(user)
	return &output, nil
}
