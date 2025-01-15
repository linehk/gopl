package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

func main() {
	printTracks(tracks)

	click("title")
	printTracks(tracks)

	click("artist")
	printTracks(tracks)

	click("album")
	printTracks(tracks)

	click("year")
	printTracks(tracks)

	click("length")
	printTracks(tracks)
}

func click(t string) {
	switch t {
	case "title":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Title < y.Title
			}})
	case "artist":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Artist < y.Artist
			}})
	case "album":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Album < y.Album
			}})
	case "year":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Year < y.Year
			}})
	case "length":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return int64(x.Length) < int64(y.Length)
			}})
	}
}

type custom struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x custom) Len() int           { return len(x.t) }
func (x custom) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x custom) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
