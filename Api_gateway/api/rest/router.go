package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/login", LoginHundler)
	r.Post("/register", RegisterHundler)
	// r.Post("/logout")
	return r
}
