package main


import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	mylibrary "github.com/BexeitAlua/golang/tsis1/pkg/my-library"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "info": mylibrary.Info()})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	books := mylibrary.GetBooks()
	respondWithJSON(w, http.StatusOK, books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	book, err := mylibrary.GetBook(name)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}


const PORT = ":8080"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health-check", healthCheck).Methods("GET")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{name}", getBook).Methods("GET")

	log.Printf("Starting server on %s\n", PORT)
	err := http.ListenAndServe(PORT, r)
	log.Fatal(err)
}


