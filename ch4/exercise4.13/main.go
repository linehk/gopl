package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	title  = flag.String("t", "", "")
	apikey = flag.String("apikey", "", "")
)

const api = "https://www.omdbapi.com/"

type result struct {
	PosterURL string `json:"Poster"`
}

func main() {
	flag.Parse()
	url := api + "?t=" + url.QueryEscape(*title) + "&apikey=" + url.QueryEscape(*apikey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get url: %v fail. err: %v", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var res result
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		fmt.Fprintf(os.Stderr, "json decode fail. err: %v", err)
		os.Exit(1)
	}

	poster, err := http.Get(res.PosterURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get poster url: %v fail. err: %v", res.PosterURL, err)
		os.Exit(1)
	}
	defer poster.Body.Close()

	f, err := os.Create(*title + ".jpeg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file fail. err: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = io.Copy(f, poster.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "io copy fail. err: %v", err)
		os.Exit(1)
	}
}
