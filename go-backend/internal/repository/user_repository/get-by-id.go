package user_repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

func (r *Repository_impl) GetById(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	query := `
	SELECT id, name, email, password, is_active, created_at, updated_at
	FROM users
	WHERE id = $1
		and deleted_at is null
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		slog.Error("Erro ao buscar usuário por id", "error", err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		slog.Error("Nenhum usuário encontrado", "error", err)
		return nil, domain.ErrUserNotFound
	}

	return &user, nil
}
