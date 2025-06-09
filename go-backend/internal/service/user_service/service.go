package user_service

import "github.com/GustavoPaula/go-microservices/go-backend/internal/repository/user_repository"

type Service_impl struct {
	repository user_repository.Repository_i
}

func New(repository user_repository.Repository_i) *Service_impl {
	return &Service_impl{repository: repository}
}
