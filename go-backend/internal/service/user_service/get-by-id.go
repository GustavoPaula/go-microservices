package user_service

import (
	"context"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) GetById(ctx context.Context, id string) (*dto.UserOutput, error) {
	user, err := s.repository.GetById(ctx, id)
	if err != nil {
		slog.Error("Erro service User retornado pelo repository GetById", "error", err)
		return nil, err
	}

	output := dto.FromUser(user)
	return &output, nil
}
