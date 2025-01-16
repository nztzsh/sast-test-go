package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nztzsh/sast-test-go/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/search", handlers.Proxy).Methods("GET")
	http.ListenAndServe(":8080", nil)
}
