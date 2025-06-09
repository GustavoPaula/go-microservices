package user_repository

import (
	"context"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository_i interface {
	Create(ctx context.Context, user *domain.User) error
	GetById(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Put(ctx context.Context, user *domain.User, id string) error
	SoftDelete(ctx context.Context, id string) (*domain.User, error)
}

type Repository_impl struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Repository_i {
	return &Repository_impl{db: db}
}
