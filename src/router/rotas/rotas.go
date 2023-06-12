package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	Uri                 string
	Metodo              string
	Funcao              func(w http.ResponseWriter, r *http.Request)
	RequerAutentication bool
}

// configurar rotas
func Configurar(r *mux.Router) *mux.Router {
	rotas := usuariosRotas
	for _, rota := range rotas {
		r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
