package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// I'd like to use my own Router (^^)
func NewRouter() http.Handler {
	router := mux.NewRouter()
	AttachRoutes(router)
	return router
}

func AttachRoutes(router *mux.Router) {
	router.HandleFunc("/hello", HelloHandler)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Stage 1: hello world
	fmt.Fprintf(w, "Hello World\n")
	// Stage 2: get some data from redis

	// Stage 3: get some data with filters
}

func main() {
	// I want to use router!
	// http.HandleFunc("/", HelloHandler)
	handler := NewRouter()
	http.Handle("/", handler)
	http.ListenAndServe(":8000", handler)
}
