package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"go.mod/src/connect"
)

func main() {
	fmt.Println("Starting")
	r := chi.NewRouter()
	m := chiprometheus.NewMiddleware("router")
	connect.Load()
	r.Use(m)
	r.Mount("/login", LoginRouter())
	r.Mount("/user", userRouter())
	r.Mount("/category", categoryRouter())
	r.Mount("/post", postRouter())
	r.Mount("/comment", CommentsRouter())
	r.Mount("/youtube", YoutbeRouter())
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
