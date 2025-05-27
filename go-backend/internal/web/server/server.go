package server

import (
	"net/http"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/service"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router      *chi.Mux
	server      *http.Server
	userService *service.UserService
	port        string
}

func NewServer(userService *service.UserService, port string) *Server {
	return &Server{
		router:      chi.NewRouter(),
		userService: userService,
		port:        port,
	}
}

func (s *Server) ConfigureRoutes() {
	userHandler := handlers.NewUserHandler(s.userService)

	s.router.Post("/users", userHandler.Create)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
