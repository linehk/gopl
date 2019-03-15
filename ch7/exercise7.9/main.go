package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/title":
		click("title")
	case "/artist":
		click("artist")
	case "/album":
		click("album")
	case "/year":
		click("year")
	case "/length":
		click("length")
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, &tracks); err != nil {
		log.Println(err)
	}
}

func click(t string) {
	switch t {
	case "title":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Title < y.Title
			},
			func(x, y *Track) {
				x.Title, y.Title = y.Title, x.Title
			}})
	case "artist":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Artist < y.Artist
			},
			func(x, y *Track) {
				x.Artist, y.Artist = y.Artist, x.Artist
			}})
	case "album":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Album < y.Album
			},
			func(x, y *Track) {
				x.Album, y.Album = y.Album, x.Album
			}})
	case "year":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Year < y.Year
			},
			func(x, y *Track) {
				x.Year, y.Year = y.Year, x.Year
			}})
	case "length":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return int64(x.Length) < int64(y.Length)
			},
			func(x, y *Track) {
				x.Length, y.Length = y.Length, x.Length
			}})
	}
}

type custom struct {
	t    []*Track
	less func(x, y *Track) bool
	swap func(x, y *Track)
}

func (x custom) Len() int           { return len(x.t) }
func (x custom) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x custom) Swap(i, j int)      { x.swap(x.t[i], x.t[j]) }

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
