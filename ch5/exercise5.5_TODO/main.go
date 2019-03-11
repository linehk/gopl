package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	images, err := CountWordsAndImages("https://www.jianshu.com/")
	if err != nil {
		fmt.Println("error!")
	}
	fmt.Println(images)
}

// CountWordsAndImages does an HTTP GET request for the html
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (images int) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		images = countWordsAndImages(c)
	}
	return
}
