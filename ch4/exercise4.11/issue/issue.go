package issue

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Params struct {
	Owner  string
	Repo   string
	Number string
	Token  string
	Issue
}

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

const baseURL = "https://api.github.com/repos/"

func (p Params) GetIssues() ([]Issue, error) {
	u := baseURL + p.Owner + "/" + p.Repo + "/issues"
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func (p Params) GetIssue() (Issue, error) {
	u := baseURL + p.Owner + "/" + p.Repo + "/issues" +
		"/" + p.Number
	resp, err := http.Get(u)
	if err != nil {
		return Issue{}, err
	}
	defer resp.Body.Close()

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return Issue{}, err
	}
	return issue, nil
}

func (p Params) CreateIssue() bool {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(p.Issue); err != nil {
		return false
	}
	u := baseURL + p.Owner + "/" + p.Repo + "/issues" +
		"?access_token=" + p.Token
	_, err := http.Post(u, "application/json", &buf)
	if err != nil {
		return false
	}
	return true
}

func (p Params) EditIssue() bool {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(p.Issue); err != nil {
		return false
	}
	u := baseURL + p.Owner + "/" + p.Repo + "/issues" +
		"/" + p.Number + "?access_token=" + p.Token
	request, err := http.NewRequest(http.MethodPatch, u, &buf)
	if err != nil {
		return false
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return false
	}
	return true
}
