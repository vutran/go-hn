package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/vutran/go-hn/lib/hn"
	"github.com/vutran/srgnt"
	"strconv"
)

const MaxCount = 5

func GetItems(ids []int, limit int) *bytes.Buffer {
	var b bytes.Buffer

	if len(ids) > 0 {
		for _, id := range ids[:limit] {
			item := hn.GetItem(id)
			fmt.Fprintf(&b, "\t%s\n\t%s\n\n", item.Title, item.Url)
		}
	}

	return &b
}

func Top(flags *flag.FlagSet) {
	limit, _ := strconv.Atoi(flags.Lookup("limit").Value.String())

	ids := hn.GetTopStories()
	items := GetItems(ids, limit)

	fmt.Println("Top Stories\n")
	fmt.Print(items.String())
}

func New(flags *flag.FlagSet) {
	limit, _ := strconv.Atoi(flags.Lookup("limit").Value.String())

	ids := hn.GetNewStories()
	items := GetItems(ids, limit)

	fmt.Println("New Stories\n")
	fmt.Print(items.String())
}

func Best(flags *flag.FlagSet) {
	limit, _ := strconv.Atoi(flags.Lookup("limit").Value.String())

	ids := hn.GetBestStories()
	items := GetItems(ids, limit)

	fmt.Println("Best Stories\n")
	fmt.Print(items.String())
}

func main() {
	cli := srgnt.CreateProgram("go-hn")

	cli.AddCommand("top", Top, "Show top stories")
	cli.AddCommand("new", New, "Show new stories")
	cli.AddCommand("best", Best, "Show best stories")
	cli.AddIntFlag("limit", MaxCount, "Specify max results")

	cli.Run()
}
