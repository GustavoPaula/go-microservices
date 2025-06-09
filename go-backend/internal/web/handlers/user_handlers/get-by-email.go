package user_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/web/response"
)

func (h *Handler_impl) GetByEmail(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.service.GetByEmail(r.Context(), input.Email)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			response.WriteError(w, http.StatusBadRequest, err.Error())
			return
		default:
			response.WriteError(w, http.StatusInternalServerError, "algo deu errado!")
			return
		}
	}

	response.WriteSuccess(w, http.StatusOK, "usuário encontrado!", output)
}
