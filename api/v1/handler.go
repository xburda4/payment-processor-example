package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r.Use(middleware.RequestID)

	// TODO finish sign up
	r.Post("/signup", nil)

	r.Route("/payments", func(r chi.Router) {
		r.Get("/", h.ListPayments)
		r.Post("/", h.CreatePayment)
		r.Get("/{paymentID}", h.GetPayment)
	})

	h.Mux = r
}
