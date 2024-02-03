package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MyLibrary struct {
	Author string `json:"author"`
	Name   string `json:"name"`
	Genre  string `json:"genre"`
	Pages  int    `json:"pages"`
}

var books = []MyLibrary{
	{"Arthur Conan Doyle", "The Adventures of Sherlock Holmes", "Detective", 654},
	{"Agatha Christie", "4:50 from Paddington", "Detective", 286},
	{"Agatha Christie", "10 Little Negroes", "Detective", 313},
	{"Agatha Christie", "ABC Murders", "Detective", 243},
	{"Agatha Christie", "Peril at End House", "Detective", 294},
	{"Oscar Wilde", "The Picture of Dorian Gray", "Philosophical fiction", 389},
	{"Stephen King", "The Green Mile", "Dark fantasy", 527},
	{"Dan Brown", "The Da Vinci Code", "Mystery", 598},
	{"Dan Brown", "Digital Fortress", "Mystery", 486},
	{"Daniel Keyes", "The Minds of Billy Milligan", "Biography", 542},
	{"Jane Austen", "Pride and Prejudice", "Romance", 498},
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	name := params["name"]

	var book MyLibrary

	for _, b := range books {
		if b.Name == name {
			book = b
			break
		}
	}

	if book.Name != "" {
		json.NewEncoder(w).Encode(book)
	} else {
		http.NotFound(w, r)
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "My Library Web App\nAuthor: Bexeit Alua")
}