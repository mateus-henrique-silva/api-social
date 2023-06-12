package main

import (
	"fmt"
	"log"
	"net/http"

	"go.mod/src/router"
)

func main() {
	fmt.Println("Starting")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
