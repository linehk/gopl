package issue

import (
	"encoding/json"
	"net/http"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []Issue
}

type Issue struct {
	Title     string
	User      User
	Milestone Milestone
}

type User struct {
	Login string
}

type Milestone struct {
	Title       string
	Description string
}

func SearchIssues(params string) (*IssuesSearchResult, error) {
	resp, err := http.Get(IssuesURL + "?q=" + params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
