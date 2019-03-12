package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var id = flag.String("id", "", "")

func main() {
	flag.Parse()

	var nodes []*html.Node
	for _, url := range flag.Args() {
		node, err := outline(url, *id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v", err)
			os.Exit(1)
		}
		nodes = append(nodes, node)
	}

	for _, node := range nodes {
		fmt.Println(node.Data)
		for _, a := range node.Attr {
			fmt.Println(a.Key, a.Val)
		}
	}
}

func outline(url, id string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	node := ElementByID(doc, id)
	if node := ElementByID(doc, id); node == nil {
		return nil, errors.New("node not found")
	}

	return node, nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, findElement, findElement)
}

func forEachNode(n *html.Node, id string, pre, post func(*html.Node, string) bool) *html.Node {
	if pre != nil {
		if !pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node := forEachNode(c, id, pre, post); node != nil {
			return node
		}
	}

	if post != nil {
		if !post(n, id) {
			return n
		}
	}

	return nil
}

func findElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
	}
	return true
}
