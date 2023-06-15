package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Starting")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/user", userRouter())

	log.Fatal(http.ListenAndServe(":5000", r))
}
