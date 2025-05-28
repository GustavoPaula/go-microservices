package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connection(ctx context.Context) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := pgxpool.New(ctx, connString)

	if err != nil {
		slog.Error("Falha ao conectar no banco de dados")
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		slog.Error("Falha a verificar a conexão com o banco de dados")
		return nil, err
	}

	m, err := migrate.New("file://./internal/database/migrations", connString)
	if err != nil {
		slog.Error("Falha ao criar a instancia do migrate")
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error("Falha ao executar as migrations")
		return nil, err
	}

	return db, nil
}
