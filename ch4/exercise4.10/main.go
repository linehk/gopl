package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/linehk/gopl/ch4/github"
)

type class string

const (
	LTOM class = "less than one month"
	MTOM class = "more than one month"
	LTOY class = "less than one year"
	MTOY class = "more than one year"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	issueClass := make(map[class][]github.Issue)
	for _, item := range result.Items {
		item := *item
		y, m, _ := item.CreatedAt.Date()
		cy, cm, _ := time.Now().Date()
		switch {
		case cm-m <= time.Month(1):
			issueClass[LTOM] = append(issueClass[LTOM], item)
		case cm-m > time.Month(1):
			issueClass[MTOM] = append(issueClass[MTOM], item)
		case cy-y <= 1:
			issueClass[LTOY] = append(issueClass[LTOY], item)
		case cy-y > 1:
			issueClass[MTOY] = append(issueClass[MTOY], item)
		}
	}

	for class, issues := range issueClass {
		fmt.Printf("class: %s, issues: %v\n", class, issues)
	}
}
