package main

import (
	"flag"
	"fmt"
	"github.com/vutran/go-hn/lib/hn"
)

const MaxCount = 5

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
		ids = hn.GetBestStories()
		fmt.Print("Best Stories:\n\n")
	case "new":
		ids = hn.GetNewStories()
		fmt.Print("New Stories:\n\n")
	case "top":
		ids = hn.GetTopStories()
		fmt.Print("Top Stories:\n\n")
	default:
		fmt.Print(Usage())
	}

	if len(ids) > 0 {
		for _, id := range ids[:*limit] {
			item := hn.GetItem(id)
			fmt.Printf("\t%s\n\t%s\n\n", item.Title, item.Url)
		}
	}
}
