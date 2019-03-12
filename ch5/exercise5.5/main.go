package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	err := CountWordsAndImages("https://www.sulinehk.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("images: %d\nwords: %d\n", images, words)
}

func CountWordsAndImages(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing HTML: %s", err)
	}

	countWordsAndImages(doc)

	return nil
}

var (
	images int
	words  int
)

func countWordsAndImages(n *html.Node) {
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countWordsAndImages(c)
	}
}
