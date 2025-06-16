package user_service

import (
	"context"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) GetByEmail(ctx context.Context, email string) (*dto.UserOutput, error) {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		slog.Error("Erro service User retornado pelo repository GetByEmail", "error", err)
		return nil, err
	}

	output := dto.FromUser(user)
	return &output, nil
}
