package user_service

import (
	"context"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
)

func (s *Service_impl) Put(ctx context.Context, input dto.CreateUserInput, id string) (*dto.UserOutput, error) {
	userFound, err := s.repository.GetById(ctx, id)
	if err != nil {
		slog.Error("Erro service User Put retornado pelo repository GetById", "error", err)
		return nil, err
	}

	if input.Name == "" {
		input.Name = userFound.Name
	}

	if input.Email == "" {
		input.Email = userFound.Email
	}

	if input.Password == "" {
		input.Password = userFound.Password
	}

	user, err := dto.ToUser(input)
	if err != nil {
		slog.Error("Erro service User Put retornado pelo dto.ToUser", "error", err)
		return nil, err
	}

	UpdatedUser, err := s.repository.Put(ctx, id, user.Name, user.Email, user.Password, user.IsActive, user.UpdatedAt)
	if err != nil {
		slog.Error("Erro service User Put retornado pelo repository Put", "error", err)
		return nil, err
	}

	output := dto.FromUser(UpdatedUser)
	return &output, nil
}
