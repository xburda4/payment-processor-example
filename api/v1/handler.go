package v1

import (
	"github.com/go-chi/chi/v5"
)

// Handler for v1 endpoints.
type Handler struct {
	*chi.Mux
}

func NewHandler() *Handler {
	h := &Handler{}
	h.initRouter()
	return h
}

func (h *Handler) initRouter() {
	r := chi.NewRouter()

	h.Mux = r
}
