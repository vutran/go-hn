package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/vutran/go-hn/lib/hn"
	"github.com/vutran/srgnt"
	"io"
	"strconv"
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

func Top(flags *flag.FlagSet) io.Reader {
	limit, _ := strconv.Atoi(flags.Lookup("limit").Value.String())
	var b bytes.Buffer

	_, _ = b.WriteString("Top Stories\n\n")

	ids := hn.GetTopStories()

	items := GetItems(ids, limit)
	b.Write(items.Bytes())

	return &b
}

func New(flags *flag.FlagSet) io.Reader {
	limit, _ := strconv.Atoi(flags.Lookup("limit").Value.String())
	var b bytes.Buffer

	_, _ = b.WriteString("New Stories\n\n")

	ids := hn.GetNewStories()

	items := GetItems(ids, limit)
	b.Write(items.Bytes())

	return &b
}

func Best(flags *flag.FlagSet) io.Reader {
	limit, _ := strconv.Atoi(flags.Lookup("limit").Value.String())
	var b bytes.Buffer

	_, _ = b.WriteString("Best Stories\n\n")

	ids := hn.GetBestStories()

	items := GetItems(ids, limit)
	b.Write(items.Bytes())

	return &b
}

func main() {
	done := make(chan bool, 1)

	cli := srgnt.CreateProgram("hn")

	cli.AddCommand("top", Top, "Show top stories")
	cli.AddCommand("new", New, "Show new stories")
	cli.AddCommand("best", Best, "Show best stories")
	cli.AddIntFlag("limit", MaxCount, "Specify max results")

	cli.Run(done)

	<-done
}
