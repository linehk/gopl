package xkcd

import (
	"encoding/json"
	"net/http"
)

const URL = "https://xkcd.com/"

const suffix = "/info.0.json"

type Info struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func Get(num string) (*Info, error) {
	url := URL + num + suffix
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	var info Info
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}
