package user_repository

import (
	"context"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

func (r *Repository_impl) Create(ctx context.Context, user *domain.User) error {
	query := `
	INSERT INTO users (name, email, password, is_active, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id
	`

	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.Password, user.IsActive, user.CreatedAt, user.UpdatedAt).
		Scan(&user.ID)
	if err != nil {
		slog.Error("Erro ao gravar dados na tabela users", "error", err)
		return err
	}

	return nil
}
