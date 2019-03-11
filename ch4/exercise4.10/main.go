package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl/ch4/github"
)

func main() {
	date := os.Args[1]
	result, err := github.SearchIssues(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		y, m, _ := item.CreatedAt.Date()
		ly, lm, _ := time.Now().Date()
		if date == "LTAM" {
			if lm-m <= time.Month(1) {
				fmt.Printf("%v\n", item.CreatedAt)
			}
		}
		if date == "LTAY" {
			if ly-y <= 1 {
				fmt.Printf("%v\n", item.CreatedAt)
			}
		}
		if date == "MTOY" {
			if ly-y >= 1 {
				fmt.Printf("%v\n", item.CreatedAt)
			}
		}
	}
}
