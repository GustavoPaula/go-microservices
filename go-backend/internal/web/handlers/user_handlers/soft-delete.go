package user_handlers

import (
	"net/http"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/domain"
	"github.com/GustavoPaula/go-microservices/go-backend/internal/web/response"
	"github.com/go-chi/chi/v5"
)

func (h *Handler_impl) SoftDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "id é obrigatório")
		return
	}

	output, err := h.service.SoftDelete(r.Context(), id)

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

	response.WriteSuccess(w, http.StatusOK, "usuário desativado!", output)
}
