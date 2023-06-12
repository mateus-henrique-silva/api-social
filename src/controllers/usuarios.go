package controllers

import (
	"net/http"

	"github.com/uticket/rest"
)

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando usuarios"))
	rest.Send(w, r)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscar usuarios"))
	rest.Send(w, "buscar usuarios")
}

func BuscandoUmUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando um usuario"))
	rest.Send(w, "buscando um usuario")
}

func AtualizartUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizar usuarios"))
	rest.Send(w, r)
}
func ApagarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("apagar usuarios"))
	rest.Send(w, r)
}
