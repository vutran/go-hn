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

func Ask(flags *flag.FlagSet) {
	limit, _ := strconv.Atoi(flags.Lookup("limit").Value.String())

	ids := hn.GetAskStories()
	items := GetItems(ids, limit)

	fmt.Println("Ask HN\n")
	fmt.Print(items.String())
}

func Show(flags *flag.FlagSet) {
	limit, _ := strconv.Atoi(flags.Lookup("limit").Value.String())

	ids := hn.GetShowStories()
	items := GetItems(ids, limit)

	fmt.Println("Show HN\n")
	fmt.Print(items.String())
}

func Jobs(flags *flag.FlagSet) {
	limit, _ := strconv.Atoi(flags.Lookup("limit").Value.String())

	ids := hn.GetJobStories()
	items := GetItems(ids, limit)

	fmt.Println("Jobs\n")
	fmt.Print(items.String())
}

func main() {
	cli := srgnt.CreateProgram("go-hn")

	cli.AddCommand("top", Top, "View top stories")
	cli.AddCommand("new", New, "View new stories")
	cli.AddCommand("best", Best, "View best stories")
	cli.AddCommand("ask", Ask, "View Ask HN stories")
	cli.AddCommand("show", Show, "View Show HN stories")
	cli.AddCommand("jobs", Jobs, "View jobs")
	cli.AddIntFlag("limit", MaxCount, "Specify max results")

	cli.Run()
}
