package user_repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

func (r *Repository_impl) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	query := `
	SELECT id, name, email, password, is_active, created_at, updated_at
	FROM users
	WHERE email = $1
		and deleted_at is null
	`

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	if err == sql.ErrNoRows {
		slog.Error("Erro interno do servidor", "error", err)
		return nil, err
	}

	return &user, nil
}
