package user_service

import (
	"context"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) GetByEmail(ctx context.Context, email string) (*dto.UserOutput, error) {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(user)
	return &output, nil
}
