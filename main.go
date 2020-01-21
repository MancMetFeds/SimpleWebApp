package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

// Page generic struct
type Page struct {
	Title  string
	Link   string
	NavBar string
	Body   []byte
}

func loadContents(t string) (*Page, error) {

	return &Page{
		Title:  strings.Title(t),
		Link:   "<link rel=\"stylesheet\" type=\"text/css\" href=\"assets/css/stylesheet.css\" media=\"screen\"/>",
		NavBar: "<Nav>\n\t<ul style=\"list-style-type:none; \"><li style=\"display: inline; padding: 5px; padding: 5px;\"><a href=\"/\">Home</a></li><li style=\"display: inline; padding: 5px;\"><a href=\"/about\">About</a></li><li style=\"display: inline; padding: 5px;\"><a href=\"/contact\">Contact</a></li></ul></Nav>",
	}, nil
}

func getHTML(w http.ResponseWriter, file string, p *Page) {
	t, err := template.ParseFiles("views/" + file + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/"):]
	log.Print(file)
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

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/"):]
	p, err := loadContents(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	getHTML(w, file, p)

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/"):]
	p, err := loadContents(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	getHTML(w, file, p)

}

func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/"):]
	t, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	err = t.Execute(w, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ImgHandler(w http.ResponseWriter, r *http.Request) {
	imgPath := r.URL.Path[len("/"):]
	t, err := template.ParseFiles(imgPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	err = t.Execute(w, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", PageHandler)
	router.HandleFunc("/about", PageHandler)
	router.HandleFunc("/contact", PageHandler)
	router.HandleFunc("/assets/{filetype}/{filename}", AssetsHandler)
	router.HandleFunc("/assets/images/{filename}", ImgHandler)
	http.Handle("/", router)
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
