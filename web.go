package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	log.SetPrefix("goweb: ")

	http.HandleFunc("/ding", handleDing)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func handleDing(w http.ResponseWriter, r *http.Request) {
	log.Printf("recv: %s\n", r.URL)
	fmt.Fprintf(w, "recv: %s\n", r.URL)
}
