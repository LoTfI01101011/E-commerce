package rest

import (
	"net/http"

	"github.com/LoTfI01101011/E-commerce/Api_gateway/api/gRPC"
)

func AuthMiddelware(user *gRPC.User) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//get the token from the request header
			token := r.Header.Get("Authorization")
			if token == "" {
				w.WriteHeader(http.StatusUnauthorized)
			}
			res, err := user.CheckUserToken(token)
			if err != nil || !res {
				w.WriteHeader(http.StatusUnauthorized)
			}
			next.ServeHTTP(w, r)
		})

	}
}
