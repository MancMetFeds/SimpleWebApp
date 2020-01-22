package main

import (
	"log"
	"net/http"
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
		Link:   "<link rel=\"stylesheet\" type=\"text/css\" href=\"assets/css/stylesheet.css\" /><link rel=\"stylesheet\" type=\"text/css\" href=\"assets/css/bootstrap.css\" /><script type=\"text/javascript\" src=\"assets/js/astroid.js\"></script>",
		NavBar: "<nav class=\"navbar navbar-expand-lg navbar-light bg-light\"><button class=\"navbar-toggler\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarNav\" aria-controls=\"navbarNav\" aria-expanded=\"false\" aria-label=\"Toggle navigation\"><span class=\"navbar-toggler-icon\"></span>        </button>        <div class=\"collapse navbar-collapse\" id=\"navbarNav\">          <ul class=\"navbar-nav\">            <li class=\"nav-item active\">              <a class=\"nav-link\" href=\"/\">Home</a>            </li>           <li class=\"nav-item\">              <a class=\"nav-link\" href=\"/about\">About this site & Me</a>            </li>            <li class=\"nav-item\">              <a class=\"nav-link\" href=\"/contact\">Contact Me</a>            </li>          </ul>        </div>      </nav>",
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
	//http.ListenAndServe(":"+os.Getenv("PORT"), router)			// uncomment before push to master
	http.ListenAndServe(":8080", nil) // comment before push to master
}
