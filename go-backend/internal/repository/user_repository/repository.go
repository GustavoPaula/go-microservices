package user_repository

import (
	"context"
	"time"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository_i interface {
	Create(ctx context.Context, name, email, password string, isActive bool, createdAt, updatedAt time.Time) (*domain.User, error)
	GetById(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Put(ctx context.Context, id, name, email, password string, isActive bool, updatedAt time.Time) (*domain.User, error)
	SoftDelete(ctx context.Context, id string) (*domain.User, error)
}

type Repository_impl struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Repository_i {
	return &Repository_impl{db: db}
}
