package main

import (
	"bytes"
	"encoding/json"
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
	http.HandleFunc("/post2ding", handlePost2Ding)

	log.Printf("serving...\n")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handleDing(w http.ResponseWriter, r *http.Request) {
	log.Printf("recv: %s\n", r.URL)
	fmt.Fprintf(w, "recv: %s\n", r.URL)
}

func handlePost2Ding(w http.ResponseWriter, r *http.Request) {
	url := "https://oapi.dingtalk.com/robot/send?access_token=a6b3039dc83a8911870274109dea4671286db8091f646aa44c7dd266bc0c579d"
	contentType := "Content-Type: application/json"

	text := make(map[string]string)
	text["content"] = "test"
	msg := make(map[string]interface{})
	msg["msgtype"] = "text"
	msg["text"] = text

	out, err := json.Marshal(msg)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err)
	} else {
		http.Post(url, contentType, bytes.NewReader(out))
		fmt.Fprintf(w, "ok: %s", string(out))
	}
}
