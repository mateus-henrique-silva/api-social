package rotas

import "net/http"

type Rota struct {
	Uri                 string
	Metodo              string
	Funcao              func(w http.ResponseWriter, r *http.Request)
	RequerAutentication bool
}
