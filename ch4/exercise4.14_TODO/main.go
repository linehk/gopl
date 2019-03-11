// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if r.Form["i"] != nil {
		result, err := SearchIssues(r.Form["i"])
		if err != nil {
			log.Print(err)
		}
		for _, v := range result.Items {
			fmt.Fprintf(w, "%s\n", v.Milestone.Id)
		}
	}
	/*
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	*/
}
