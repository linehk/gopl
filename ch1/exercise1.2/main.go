package main

import (
	"fmt"
	"os"
)

func main() {
	for key, arg := range os.Args[1:] {
		fmt.Println(key, arg)
	}
}
