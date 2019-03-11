package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
		for _, a := range n.Attr {
			texts = append(texts, a.Val)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = visit(texts, c)
	}
	return texts
}
