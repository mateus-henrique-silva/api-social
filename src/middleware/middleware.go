package middleware

import (
	"fmt"
	"net/http"

	"go.mod/src/autenticacao"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verificar a autenticação aqui
		// Por exemplo, verificar se o token de autenticação está presente no cabeçalho da solicitação
		authenticated, err := autenticacao.ValidationToken(r)
		if err != nil {
			return
		}
		if authenticated {
			// Se estiver autenticado, chamar o próximo manipulador
			next.ServeHTTP(w, r)
		} else {
			// Se não estiver autenticado, retornar um erro de autenticação
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Autenticação necessária.")
		}
	})
}
