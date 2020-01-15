package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

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

func main() {
	p1 := &Page{
		Title: "TestPage",
		Body:  []byte("This is a sample Page"),
	}
	p1.save()

	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
