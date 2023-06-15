package router

import "github.com/gorilla/mux"

func Gerar() *mux.Router {
	return mux.NewRouter()
}
