package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func userRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", postUser)
	return r
}

func postUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

}
