package user_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/dto"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/web/response"
	"github.com/go-chi/chi/v5"
)

func (h *Handler_impl) Put(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "id é obrigatório")
		return
	}

	var input dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.service.Put(r.Context(), input, id)

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

	response.WriteSuccess(w, http.StatusOK, "usuário alterado!", output)
}
