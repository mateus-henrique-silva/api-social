package router

import (
	"github.com/gorilla/mux"
	"go.mod/src/router/rotas"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
