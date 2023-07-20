package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

// ParseBody interpreta o corpo de uma requisição no formato JSON.
func ParseBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	jsonDec := json.NewDecoder(r.Body)
	if err := jsonDec.Decode(v); err != nil {
		log.Println(err)
		if err := SendError(w, JSONSyntaxError()); err != nil {
			return err
		}
		return err
	}
	r.Body.Close()
	return nil
}
