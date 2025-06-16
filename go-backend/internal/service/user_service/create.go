package user_service

import (
	"context"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) Create(ctx context.Context, input dto.CreateUserInput) (*dto.UserOutput, error) {
	user, err := dto.ToUser(input)
	if err != nil {
		slog.Error("Erro service User Create retornado pelo dto.ToUser", "error", err)
		return nil, err
	}

	existingUser, _ := s.repository.GetByEmail(ctx, user.Email)
	if existingUser != nil {
		slog.Error("Erro service User Create retornado pelo repository GetByEmail", "error", err)
		return nil, domain.ErrUserAlreadyExists
	}

	newUser, err := s.repository.Create(ctx, user.Name, user.Email, user.Password, user.IsActive, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		slog.Error("Erro no service User Create retornado pelo repository Create", "error", err)
		return nil, err
	}

	output := dto.FromUser(newUser)
	return &output, nil
}
