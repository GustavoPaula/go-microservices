package user_repository

import (
	"context"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

func (r *Repository_impl) SoftDelete(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	query := `
		UPDATE users
		SET deleted_at = current_timestamp, is_active = false
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)

	if err != nil {
		slog.Error("Erro ao deletar os dados do usuário", "error", err)
		return nil, err
	}

	if result.RowsAffected() == 0 {
		slog.Error("Nenhum usuário encontrado com o id", "error", err)
		return nil, err
	}

	return &user, nil
}
