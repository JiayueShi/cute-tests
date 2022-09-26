package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Stage 1: hello world
	fmt.Fprintf(w, "Hello World")
	// Stage 2: get some data from redis
	// Stage 3: get some data with filters
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}
