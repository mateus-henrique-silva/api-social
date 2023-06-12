package rotas

import "net/http"

var usuariosRotas = []Rota{
	{
		Uri:    "usarios",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutentication: false,
	},
	{
		Uri:    "usarios",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutentication: false,
	},
	{
		Uri:    "usarios",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutentication: false,
	},
	{
		Uri:    "usarios",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutentication: false,
	},
}
