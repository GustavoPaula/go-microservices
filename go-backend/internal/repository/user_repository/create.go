package user_repository

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
)

func (r *Repository_impl) Create(ctx context.Context, name, email, password string, isActive bool, createdAt, updatedAt time.Time) (*domain.User, error) {
	var user domain.User
	query := `
	INSERT INTO users (name, email, password, is_active, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id, name, email, password, is_active, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query, name, email, password, isActive, createdAt, updatedAt).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		slog.Error("Erro ao encontrar usuário", "error", err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		slog.Error("Erro ao gravar dados na tabela users", "error", err)
		return nil, err
	}

	return &user, nil
}
