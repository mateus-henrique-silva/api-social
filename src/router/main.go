package main

import (
	"fmt"
	"log"
	"net/http"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting")
	r := chi.NewRouter()
	m := chiprometheus.NewMiddleware("router")
	r.Use(m)
	r.Mount("/user", userRouter())

	log.Fatal(http.ListenAndServe(":5000", r))
}