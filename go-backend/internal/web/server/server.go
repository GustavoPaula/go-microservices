package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/service/user_service"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/web/handlers/user_handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router      *chi.Mux
	server      *http.Server
	userService *user_service.Service_impl
	port        string
}

func NewServer(userService *user_service.Service_impl, port string) *Server {
	return &Server{
		router:      chi.NewRouter(),
		userService: userService,
		port:        port,
	}
}

func (s *Server) ConfigureRoutes() {
	userHandler := user_handlers.New(s.userService)

	s.router.Group(func(r chi.Router) {
		s.router.Post("/users", userHandler.Create)
		s.router.Get("/users", userHandler.List)
		s.router.Get("/users/{id}", userHandler.GetById)
		s.router.Get("/users/email", userHandler.GetByEmail)
		s.router.Put("/users/{id}", userHandler.Put)
		s.router.Delete("/users/{id}", userHandler.SoftDelete)
	})
}

func (s *Server) Start(ctx context.Context) error {
	s.server = &http.Server{
		Addr:         ":" + s.port,
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  time.Minute,
	}

	defer func() {
		const timeout = 30 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			slog.Error("Falha ao desligar o servidor HTTP", "error", err)
		}
	}()

	errChannel := make(chan error, 1)
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			errChannel <- err
		}
	}()
	slog.Info("Servidor HTTP em execução!")

	select {
	case <-ctx.Done():
		return nil
	case err := <-errChannel:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}
