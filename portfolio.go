package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

// Page generic struct
type Page struct {
	Title string
	Body  []byte
}

func loadContents(t string) (*Page, error) {
	return &Page{
		Title: strings.Title(t),
	}, nil
}

func getHTML(w http.ResponseWriter, file string, p *Page) {
	t, err := template.ParseFiles(file + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func viewController(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/"):]
	p, err := loadContents(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	if file == "" {
		file = "index"
		p.Title = "Home"
	}
	getHTML(w, file, p)

}

func main() {
	http.HandleFunc("/", viewController)
	http.HandleFunc("/about", viewController)
	http.HandleFunc("/contact", viewController)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
