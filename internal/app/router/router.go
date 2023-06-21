package router

import (
	"github.com/go-chi/chi"

	"github.com/freepaddler/yap-shortener/internal/app/handlers"
)

func NewHttpRouter(h *handlers.HTTPHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", h.Put)
	r.Get("/{id}", h.Get)
	return r
}
