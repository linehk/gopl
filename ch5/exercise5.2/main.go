package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse: %v\n", err)
		os.Exit(1)
	}

	visit(doc)

	for t, c := range count {
		fmt.Println(t, c)
	}
}

var count = make(map[string]int)

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
