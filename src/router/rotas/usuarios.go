package rotas

import "net/http"

var usuariosRotas = []Rota{
	{
		Uri:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutentication: false,
	},
	{
		Uri:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutentication: false,
	},
	{
		Uri:    "/usuarios/{id}",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutentication: false,
	},
	{
		Uri:    "/usuarios/{id}",
		Metodo: http.MethodPut,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutentication: false,
	},
}
