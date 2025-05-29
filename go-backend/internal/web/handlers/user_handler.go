package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/service"
)

type UserHandler_impl struct {
	userService *service.UserService_impl
}

func NewUserHandler(userService *service.UserService_impl) *UserHandler_impl {
	return &UserHandler_impl{userService: userService}
}

func (h *UserHandler_impl) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.userService.CreateUser(r.Context(), input)

	if err != nil {
		switch err {
		case domain.ErrInvalidEmail:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case domain.ErrInvalidPassword:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
