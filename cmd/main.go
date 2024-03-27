package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dubey2k/go_starter/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	fmt.Println("server running at 3000")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
