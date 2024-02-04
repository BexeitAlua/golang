package pkg

import (
	
	"errors"
)

type MyLibrary struct {
	Author string `json:"author"`
	Name   string `json:"name"`
	Genre  string `json:"genre"`
	Pages  int    `json:"pages"`
}

var Books = []MyLibrary{
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

func Info() string {
	return "This is the MyLibrary package"
}

func GetBooks() []MyLibrary {
	return Books
}

func GetBook(name string) (*MyLibrary, error) {
	for _, b := range Books {
		if b.Name == name {
			return &b, nil
		}
	}
	return nil, errors.New("Book not found")
}

