package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/linehk/gopl/ch4/exercise4.12_TODO/xkcd"
)

var num = flag.String("num", "", "")

func main() {
	info, err := xkcd.Get(*num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get info fail, err: %v", err)
		os.Exit(1)
	}
	fmt.Println(*info)
}
