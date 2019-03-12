package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	visit(doc)

	fmt.Println(links)
	fmt.Println(imgs)
	fmt.Println(scripts)
	fmt.Println(styles)
}

var (
	links   []string
	imgs    []string
	scripts []string
	styles  []string
)

func visit(n *html.Node) {
	if n.Type != html.ElementNode {
		goto LOOP
	}

	switch n.Data {
	case "a":
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	case "img":
		for _, s := range n.Attr {
			if s.Key == "src" {
				imgs = append(imgs, s.Val)
			}
		}
	case "script":
		for _, s := range n.Attr {
			if s.Key == "src" {
				scripts = append(scripts, s.Val)
			}
		}
	case "link":
		for _, m := range n.Attr {
			if m.Key == "media" {
				styles = append(styles, m.Val)
			}
		}
	}

LOOP:
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
