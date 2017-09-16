package hn

import (
	"encoding/json"
	"github.com/vutran/go-hn/lib/utils"
	"log"
	"strconv"
)

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

func GetStories(list string) []int {
	url := Hostname + "/" + list + ".json"
	b := utils.Get(url)

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

func GetAskStories() []int {
	return GetStories("askstories")
}

func GetShowStories() []int {
	return GetStories("showstories")
}

func GetJobStories() []int {
	return GetStories("jobstories")
}

func GetItem(id int) Item {
	url := Hostname + "/item/" + strconv.Itoa(id) + ".json"
	b := utils.Get(url)

	item := Item{}

	if parseErr := json.Unmarshal(b.Bytes(), &item); parseErr != nil {
		log.Fatal(parseErr)
	}

	return item

}
