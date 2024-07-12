package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranty97/book-managment-system-mysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9091", r))
}
