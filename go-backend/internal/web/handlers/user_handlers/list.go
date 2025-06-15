package user_handlers

import (
	"net/http"
	"strconv"

	"github.com/GustavoPaula/go-microservices/go-backend/internal/web/response"
)

func (h *Handler_impl) List(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	if pageStr == "" || limitStr == "" {
		response.WriteError(w, http.StatusBadRequest, "page e limit são obrigatórios")
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "page inválido")
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "limit inválido")
		return
	}

	output, err := h.service.List(r.Context(), page, limit)
	if err != nil {
		switch err {
		default:
			response.WriteError(w, http.StatusInternalServerError, "algo deu errado!")
			return
		}
	}

	response.WriteSuccess(w, http.StatusOK, "usuários encontrados!", output)
}
