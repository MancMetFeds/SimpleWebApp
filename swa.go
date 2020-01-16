package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

}

// pages are defined as structs
type Page struct {
	Title string
	Body  []byte
}

//expects pointer to page, returns error msg
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
		return nil, err // generic error handling
	}

	return &Page{ // returns pointer to a populated 'Page' struct
		Title: title,
		Body:  body,
	}, nil
}

func main()
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
