package middleware

import (
	"fmt"
	"net/http"

	"go.mod/src/autenticacao"
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authenticated, err := autenticacao.ValidateToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Autenticação necessária.")
			return
		}
		if authenticated {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Autenticação necessária.")
		}
	}
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := autenticacao.ValidateToken(r); err != nil {
			return
		}
		next(w, r)
	}
}
