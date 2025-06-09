package user_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain/commons"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/web/response"
)

func (h *Handler_impl) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.service.Create(r.Context(), input)

	if err != nil {
		switch err {
		case commons.ErrInvalidEmail:
			response.WriteError(w, http.StatusBadRequest, err.Error())
			return
		case commons.ErrInvalidPassword:
			response.WriteError(w, http.StatusBadRequest, err.Error())
			return
		case domain.ErrUserAlreadyExists:
			response.WriteError(w, http.StatusBadRequest, err.Error())
			return
		default:
			response.WriteError(w, http.StatusInternalServerError, "algo deu errado!")
			return
		}
	}

	response.WriteSuccess(w, http.StatusCreated, "usuário criado com sucesso!", output)
}
