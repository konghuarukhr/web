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
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

	http.HandleFunc("/ding", handleDing)
	http.HandleFunc("/post2ding", handlePost2Ding)

	log.Printf("serving...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleDing(w http.ResponseWriter, r *http.Request) {
	log.Printf("recv: %s\n", r.URL)
	fmt.Fprintf(w, "recv: %s\n", r.URL)
	handlePost2Ding(w, r)
}

func handlePost2Ding(w http.ResponseWriter, r *http.Request) {
	url := "https://oapi.dingtalk.com/robot/send?access_token=a6b3039dc83a8911870274109dea4671286db8091f646aa44c7dd266bc0c579d"
	contentType := "application/json"

	// text := make(map[string]string)
	// text["content"] = "消息: test"
	// msg := make(map[string]interface{})
	// msg["msgtype"] = "text"
	// msg["text"] = text
	msg := makeDingActionCard("test", "![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划")

	out, err := json.Marshal(msg)
	if err != nil {
		fmt.Fprintf(w, "error: %s\n", err)
	} else {
		log.Printf("msg: %s\n", out)
		resp, err := http.Post(url, contentType, bytes.NewReader(out))
		if err != nil {
			fmt.Fprintf(w, "error: %s\n", err)
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(w, "error: %s\n", err)
			} else {
				fmt.Fprintf(w, "ok: %s\n", body)
			}
		}
	}
}

type dingActionCard struct {
	Msgtype    string     `json:"msgtype"`
	ActionCard actionCard `json:"actionCard"`
}

type actionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	Btns           []btn  `json:"btns"`
}

type btn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

func makeDingActionCard(title string, text string) *dingActionCard {
	return &dingActionCard{
		Msgtype: "actionCard",
		ActionCard: actionCard{
			Title:          title,
			Text:           text,
			HideAvatar:     "0",
			BtnOrientation: "0",
			Btns: []btn{
				{Title: "标题1", ActionURL: "https://konghuarukhr.github.io"},
				{Title: "标题1", ActionURL: "https://konghuarukhr.github.io"},
			},
		}}
}
