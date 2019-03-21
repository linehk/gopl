package main

import (
	"fmt"
	"log"
	"os"

	"github.com/linehk/gopl/ch8/exercise8.10/links"
)

func crawl(url string, cancelled chan struct{}) []string {
	fmt.Println(url)
	list, err := links.Extract(url, cancelled)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	cancelled := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancelled)
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, cancelled)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
