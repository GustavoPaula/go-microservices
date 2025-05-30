package repository

import (
	"context"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository_i interface {
	Save(ctx context.Context, user *domain.User) error
}

type UserRepository_impl struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository_i {
	return &UserRepository_impl{db: db}
}

func (r *UserRepository_impl) Save(ctx context.Context, user *domain.User) error {
	query := `
	INSERT INTO users (name, email, password, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING id
	`

	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		slog.Error("Erro ao gravar dados na tabela users", "error", err)
		return err
	}

	return nil
}
