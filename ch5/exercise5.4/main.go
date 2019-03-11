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

	links, imgs, scripts, styles := visit(nil, nil, nil, nil, doc)

	fmt.Println(links)
	fmt.Println(imgs)
	fmt.Println(scripts)
	fmt.Println(styles)
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, imgs []string, scripts []string, styles []string, n *html.Node) ([]string, []string, []string, []string) {
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}

		if n.Data == "img" {
			for _, s := range n.Attr {
				if s.Key == "src" {
					imgs = append(imgs, s.Val)
				}
			}
		}

		if n.Data == "script" {
			for _, s := range n.Attr {
				if s.Key == "src" {
					scripts = append(scripts, s.Val)
				}
			}
		}

		if n.Data == "link" {
			for _, m := range n.Attr {
				if m.Key == "media" {
					styles = append(styles, m.Val)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links, imgs, scripts, styles = visit(links, imgs, scripts, styles, c)
	}
	return links, imgs, scripts, styles
}
