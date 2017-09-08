package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const MaxCount = 5
const Hostname = "https://hacker-news.firebaseio.com/v0"

type Item struct {
	Kids        []int  `json:kids`
	Descendents int    `json:descendents`
	Id          int    `json:id`
	Score       int    `json:score`
	Time        int    `json:time`
	By          string `json:by`
	Title       string `json:title`
	Url         string `json:url`
}

func Get(url string) *bytes.Buffer {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		log.Fatal(err)
	}

	return &b
}

func GetTopStories() []int {
	url := Hostname + "/topstories.json"
	b := Get(url)

	var ids []int
	if parseErr := json.Unmarshal(b.Bytes(), &ids); parseErr != nil {
		log.Fatal(parseErr)
	}

	return ids
}

func GetItem(id int) Item {
	url := Hostname + "/item/" + strconv.Itoa(id) + ".json"
	b := Get(url)

	item := Item{}

	if parseErr := json.Unmarshal(b.Bytes(), &item); parseErr != nil {
		log.Fatal(parseErr)
	}

	return item

}

func main() {
	ids := GetTopStories()

	fmt.Print("Top Stories:\n\n")

	for _, id := range ids[:MaxCount] {
		item := GetItem(id)
		fmt.Printf("\t%s\n\t%s\n\n", item.Title, item.Url)
	}
}
