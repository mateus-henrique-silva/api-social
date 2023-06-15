package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting")
	r := chi.NewRouter()
	r.Mount("/user", userRouter())

	log.Fatal(http.ListenAndServe(":5000", r))
}
