package controllers

import "net/http"

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando usuarios"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscar usuarios"))
}

func BuscandoUmUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando um usuario"))
}

func AtualizartUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizar usuarios"))
}
func ApagarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("apagar usuarios"))
}
