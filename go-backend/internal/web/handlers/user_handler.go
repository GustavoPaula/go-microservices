package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain/commons"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/service"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/web/response"
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
		case commons.ErrInvalidEmail:
			response.WriteError(w, http.StatusBadRequest, err.Error())
			return
		case commons.ErrInvalidPassword:
			response.WriteError(w, http.StatusBadRequest, err.Error())
			return
		default:
			response.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	response.WriteSuccess(w, http.StatusCreated, "usuário criado com sucesso!", output)
}
