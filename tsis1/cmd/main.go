package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/BexeitAlua/golang/tsis1/pkg"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", pkg.GetBooksHandler).Methods("GET")
	r.HandleFunc("/books/{name}", pkg.GetBookHandler).Methods("GET")
	r.HandleFunc("/health", pkg.HealthCheckHandler).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}