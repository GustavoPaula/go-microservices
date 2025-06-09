package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/database/dbconfig"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/repository/user_repository"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/service/user_service"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/web/server"
	"github.com/joho/godotenv"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error ao carregar o arquivo .env")
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	db, err := dbconfig.Connection(ctx)
	if err != nil {
		panic(err)
	}

	userRepository := user_repository.New(db)
	userService := user_service.New(userRepository)

	port := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(userService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(ctx); err != nil {
		slog.Error("Erro ao iniciar o servidor HTTP", "error", err)
		return
	}

	slog.Info("Servidor offline!")
}
