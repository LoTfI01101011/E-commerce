package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/api/login", LoginHundler)
	r.Post("/api/register", RegisterHundler)
	r.Post("/api/logout", Logout)
	return r
}
