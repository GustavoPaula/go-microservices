package user_service

import (
	"context"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) List(ctx context.Context, page, limit int) (*[]dto.UserOutput, error) {
	users, err := s.repository.List(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	output := dto.FromUsers(users)
	return &output, nil
}
