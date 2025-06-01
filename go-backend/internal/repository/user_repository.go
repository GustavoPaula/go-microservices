package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository_i interface {
	Save(ctx context.Context, user *domain.User) error
	FindById(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User, id string) error
}

type UserRepository_impl struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository_i {
	return &UserRepository_impl{db: db}
}

func (r *UserRepository_impl) Save(ctx context.Context, user *domain.User) error {
	query := `
	INSERT INTO users (name, email, password, is_active, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id
	`

	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.Password, user.IsActive, user.CreatedAt, user.UpdatedAt).
		Scan(&user.ID)
	if err != nil {
		slog.Error("Erro ao gravar dados na tabela users", "error", err)
		return err
	}

	return nil
}

func (r *UserRepository_impl) FindById(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	query := `
	SELECT id, name, email, password, is_active, created_at, updated_at
	FROM users
	WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows || err != nil {
		return nil, domain.ErrUserNotFound
	}

	return &user, nil
}

func (r *UserRepository_impl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	query := `
	SELECT id, name, email, password, is_active, created_at, updated_at
	FROM users
	WHERE email = $1
	`

	err := r.db.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows || err != nil {
		return nil, domain.ErrUserNotFound
	}

	return &user, nil
}

func (r *UserRepository_impl) Update(ctx context.Context, user *domain.User, id string) error {
	query := `
	UPDATE users
	SET name = $1, email = $2, password = $3, is_active = $4, updated_at = $5
	WHERE id = $6
	`

	result, err := r.db.Exec(ctx, query, user.Name, user.Email, user.Password, user.IsActive, user.UpdatedAt, id)

	if err != nil {
		slog.Error("Erro ao atualizar os dados do usuário", "error", err)
		return err
	}

	if result.RowsAffected() == 0 {
		slog.Error("Nenhum usuário encontrado com o id", "error", err)
		return err
	}

	return nil
}
