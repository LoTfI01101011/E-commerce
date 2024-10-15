package rest

import (
	"net/http"

	"github.com/LoTfI01101011/E-commerce/Api_gateway/api/gRPC"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(user *gRPC.User) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/api/login", LoginHundler)
	r.Post("/api/register", func(w http.ResponseWriter, r *http.Request) {
		RegisterHundler(w, r, user)
	})
	r.Post("/api/logout", Logout)
	return r
}
