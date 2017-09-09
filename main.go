package main

import (
	"bytes"
	"encoding/json"
	"flag"
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

func GetStories(list string) []int {
	url := Hostname + "/" + list + ".json"
	b := Get(url)

	var ids []int
	if parseErr := json.Unmarshal(b.Bytes(), &ids); parseErr != nil {
		log.Fatal(parseErr)
	}

	return ids
}

func GetNewStories() []int {
	return GetStories("newstories")
}

func GetBestStories() []int {
	return GetStories("beststories")
}

func GetTopStories() []int {
	return GetStories("topstories")
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

func Usage() string {
	return `
Usage:

	$ hn [flags] <command>

Commands:

	top		Display top posts
	new		Display new posts
	best		Display best posts

Flags:

	--limit		Cap max results

Examples:

	$ hn --limit 3 top
	$ hn best

`
}

func main() {
	limit := flag.Int("limit", MaxCount, "Cap max results")
	flag.Parse()
	cmd := flag.Arg(0)

	var ids []int

	switch cmd {
	case "best":
		ids = GetBestStories()
		fmt.Print("Best Stories:\n\n")
	case "new":
		ids = GetNewStories()
		fmt.Print("New Stories:\n\n")
	case "top":
		ids = GetTopStories()
		fmt.Print("Top Stories:\n\n")
	default:
		fmt.Print(Usage())
	}

	if len(ids) > 0 {
		for _, id := range ids[:*limit] {
			item := GetItem(id)
			fmt.Printf("\t%s\n\t%s\n\n", item.Title, item.Url)
		}
	}
}
