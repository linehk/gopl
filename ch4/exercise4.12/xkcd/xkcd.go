package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	MinNum = 1
	MaxNum = 2125
)

type comic struct {
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

func (c *comic) String() string {
	return fmt.Sprintf("Comic: %d\n", c.Num) +
		fmt.Sprintf("Image: %s\n", c.Img) +
		fmt.Sprintf("Transcript: %s\n", c.Transcript)
}

type Index struct {
	Comics []*comic
}

func New() Index {
	return Index{[]*comic{}}
}

func Get(num int) (*comic, error) {
	url := "https://xkcd.com/" + strconv.Itoa(num) + "/info.0.json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	var comic comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}
	return &comic, nil
}

func Search(index Index, keywords []string) []*comic {
	var result []*comic
	for _, c := range index.Comics {
		isMatch := true
		for _, term := range keywords {
			if !match(c, term) {
				isMatch = false
			}
		}
		if isMatch {
			result = append(result, c)
		}
	}
	return result
}

func match(c *comic, keyword string) bool {
	return strings.Contains(c.Month, keyword) ||
		strings.Contains(strconv.Itoa(c.Num), keyword) ||
		strings.Contains(c.Link, keyword) ||
		strings.Contains(c.Year, keyword) ||
		strings.Contains(c.News, keyword) ||
		strings.Contains(c.SafeTitle, keyword) ||
		strings.Contains(c.Transcript, keyword) ||
		strings.Contains(c.Alt, keyword) ||
		strings.Contains(c.Img, keyword) ||
		strings.Contains(c.Title, keyword) ||
		strings.Contains(c.Day, keyword)
}
