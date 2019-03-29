package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/linehk/gopl/ch4/exercise4.12/xkcd"
)

var (
	f = flag.Bool("f", false, "")
	n = flag.Int("n", 100, "")
)

// fetch: go run main.go -f -n=100 > in.json
// search: cat in.json | go run main.go keywords
func main() {
	flag.Parse()
	if *f {
		if *n > xkcd.MaxNum {
			log.Fatalf("%d can't bigger than %d", *n, xkcd.MaxNum)
		}
		fetch(*n)
	} else {
		search(flag.Args())
	}
}

func fetch(n int) {
	index := xkcd.New()
	for num := xkcd.MinNum; num < n; num++ {
		c, err := xkcd.Get(num)
		if err != nil {
			log.Fatal(err)
		}
		index.Comics = append(index.Comics, c)
	}
	out, err := json.MarshalIndent(index, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}

func search(keywords []string) {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	index := xkcd.New()
	if err := json.Unmarshal(in, &index); err != nil {
		log.Fatal(err)
	}
	result := xkcd.Search(index, keywords)
	for _, c := range result {
		fmt.Println(c)
	}
}
