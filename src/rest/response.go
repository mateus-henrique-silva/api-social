package rest

import (
	"encoding/json"
	"net/http"
)

// Send envia resposta de sucesso.
func Send(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	jsonEnc := json.NewEncoder(w)
	return jsonEnc.Encode(v)
}

// SendError envia resposta de erro.
func SendError(w http.ResponseWriter, err error) error {
	e, ok := err.(*Error)
	if !ok {
		e = &Error{Status: 500, Code: "internal", Message: err.Error()}
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(e.Status)
	jsonEnc := json.NewEncoder(w)
	return jsonEnc.Encode(e)
}
