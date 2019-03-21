package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	cancel := make(chan struct{})
	resps := make(chan string, len(os.Args[1:]))
	for _, url := range os.Args[1:] {
		url := url
		go func() {
			resps <- fetch(url, cancel)
		}()
	}
	resp := <-resps
	close(cancel)
	fmt.Println(resp)
}

func fetch(url string, cancel <-chan struct{}) string {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ""
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s", b)
}
