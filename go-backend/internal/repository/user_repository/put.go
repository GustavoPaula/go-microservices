package user_repository

import (
	"context"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

func (r *Repository_impl) Put(ctx context.Context, user *domain.User, id string) error {
	query := `
	UPDATE users
	SET name = $1, email = $2, password = $3, is_active = $4, updated_at = $5
	WHERE id = $6
		and deleted_at is null
	`

	result, err := r.db.Exec(ctx, query, user.Name, user.Email, user.Password, user.IsActive, user.UpdatedAt, id)

	if err != nil {
		slog.Error("Erro ao atualizar os dados do usuário", "error", err)
		return err
	}

	if result.RowsAffected() == 0 {
		slog.Error("Nenhum usuário encontrado com o id", "error", err)
		return domain.ErrUserNotFound
	}

	return nil
}
