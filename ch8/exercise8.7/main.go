package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/linehk/gopl/ch5/links"
)

var (
	base = flag.String("base", "https://www.sulinehk.com", "")
)

func main() {
	flag.Parse()
	for _, url := range crawl(*base) {
		wg.Add(1)
		url := url
		go func() {
			defer wg.Done()
			download(*base, url)
		}()
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	<-done
}

var wg sync.WaitGroup

func download(base, url string) {
	if !strings.HasPrefix(url, base) {
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	dir := strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalln(err)
	}

	filename := dir + "index.html"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
