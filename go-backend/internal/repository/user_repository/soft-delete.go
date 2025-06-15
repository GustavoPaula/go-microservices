package user_repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

func (r *Repository_impl) SoftDelete(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	query := `
		UPDATE users
		SET deleted_at = current_timestamp, is_active = false
		WHERE id = $1
		RETURNING id, name, email, password, is_active, created_at, updated_at, deleted_at
	`

	err := r.db.QueryRow(ctx, query, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		slog.Error("Usuário não encontrado", "error", err)
		return nil, domain.ErrUserNotFound
	}

	if err == sql.ErrNoRows {
		slog.Error("Erro ao deletar os dados do usuário", "error", err)
		return nil, err
	}

	return &user, nil
}
