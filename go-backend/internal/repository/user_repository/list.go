package user_repository

import (
	"context"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

func (r *Repository_impl) List(ctx context.Context, page, limit int) ([]*domain.User, error) {
	var users []*domain.User
	offset := (page - 1) * limit

	query := `
		SELECT id, name, email, is_active, created_at, updated_at
		FROM users
		WHERE deleted_at is null
		ORDER BY name
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)

	if err != nil {
		slog.Error("Erro ao buscar usuários", "error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &domain.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt); err != nil {
			slog.Error("Erro ao ler a linha", "error", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Erro ao iterar resultados", "error", err)
		return nil, err
	}

	return users, nil
}
