package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var s []string
var count = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	outline(nil, doc)
	countSameElement(s)
	for k, v := range count {
		fmt.Printf("%s = %d\n", k, v)
	}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		for _, k := range stack {
			s = append(s, k)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func countSameElement(element []string) {
	for _, v := range element {
		count[v]++
	}
}
