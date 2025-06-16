package user_service

import (
	"context"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) SoftDelete(ctx context.Context, id string) (*dto.UserOutput, error) {
	userDeleted, err := s.repository.SoftDelete(ctx, id)
	if err != nil {
		slog.Error("Erro service User SoftDelete retornado pelo repository SoftDelete", "error", err)
		return nil, err
	}

	output := dto.FromUser(userDeleted)
	return &output, nil
}
