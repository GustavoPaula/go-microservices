package repository

import (
	"context"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, user *domain.User) error {
	query := `
	INSERT INTO users (name, email, password, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING id
	`

	err := r.db.QueryRow(ctx, query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
