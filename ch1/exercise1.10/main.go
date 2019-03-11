// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchTwoTimes(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
}

func fetchTwoTimes(url string, ch chan<- string) {
	ch <- "url: " + url + "\nresp1: " + fetch(url) + "\n" + "resp2: " + fetch(url)
}

func fetch(url string) string {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprint(err)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		return fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.2fs %7d", secs, nbytes)
}
