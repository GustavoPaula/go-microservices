package user_service

import (
	"context"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) SoftDelete(ctx context.Context, id string) (*dto.UserOutput, error) {
	userDeleted, err := s.repository.SoftDelete(ctx, id)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(userDeleted)
	return &output, nil
}
