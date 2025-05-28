package main

import (
	"context"
	"log"
	"os"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/database/config"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/repository"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/service"
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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error ao carregar o arquivo .env")
	}

	db, err := config.Connection(context.Background())
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	port := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(userService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
