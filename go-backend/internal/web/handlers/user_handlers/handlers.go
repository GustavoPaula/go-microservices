package user_handlers

import "github.com/GustavoPaula/go-microservices/go-backend/internal/service/user_service"

type Handler_impl struct {
	service *user_service.Service_impl
}

func New(service *user_service.Service_impl) *Handler_impl {
	return &Handler_impl{service: service}
}
