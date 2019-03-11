package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/linehk/gopl/ch4/exercise4.11/issue"
)

var (
	create = flag.Bool("c", false, "")
	list   = flag.Bool("l", false, "")
	read   = flag.Bool("r", false, "")
	edit   = flag.Bool("e", false, "")

	owner  = flag.String("owner", "", "")
	repo   = flag.String("repo", "", "")
	number = flag.String("number", "", "")
	token  = flag.String("token", "", "")

	title = flag.String("title", "", "")
	body  = flag.String("body", "", "")
)

func main() {
	flag.Parse()
	switch {
	case *list:
		p := issue.Params{Owner: *owner, Repo: *repo}
		issues, err := p.GetIssues()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		for _, i := range issues {
			fmt.Printf("%s\t%s\n", i.Title, i.Body)
		}
	case *read:
		p := issue.Params{Owner: *owner, Repo: *repo,
			Number: *number}
		i, err := p.GetIssue()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Printf("%s\t%s\n", i.Title, i.Body)
	case *create:
		p := issue.Params{Owner: *owner, Repo: *repo,
			Token: *token,
			Issue: issue.Issue{Title: *title, Body: *body}}
		if !p.CreateIssue() {
			fmt.Fprintf(os.Stderr, "create issue fail")
		}
	case *edit:
		p := issue.Params{Owner: *owner, Repo: *repo,
			Token: *token, Number: *number,
			Issue: issue.Issue{Title: *title, Body: *body}}
		if !p.EditIssue() {
			fmt.Fprintf(os.Stderr, "edit issue fail")
		}
	}
}
