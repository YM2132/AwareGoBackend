package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter() // Declare a new router
	r.HandleFunc("/hello", handler).Methods("GET")

	return r
}

func main() {

	r := newRouter()

	// We can then pass our router(after declaring all our routes,
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello boiiiii!")
}
