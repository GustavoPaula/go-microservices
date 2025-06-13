package user_repository

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

func (r *Repository_impl) Put(ctx context.Context, id, name, email, password string, isActive bool, updatedAt time.Time) (*domain.User, error) {
	var user domain.User

	query := `
	UPDATE users
	SET name = $1, email = $2, password = $3, is_active = $4, updated_at = $5
	WHERE id = $6
		and deleted_at is null
	RETURNING id, name, email, password, is_active, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query, name, email, password, isActive, updatedAt, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		slog.Error("Usuário não encontrado", "error", err)
		return nil, domain.ErrUserNotFound
	}

	if err == sql.ErrNoRows {
		slog.Error("Erro ao atualizar os dados do usuário", "error", err)
		return nil, err
	}

	return &user, nil
}
