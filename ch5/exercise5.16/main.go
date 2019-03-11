package main

import (
	"fmt"
)

func main() {
	fmt.Println(join("aaa", "bbb", "ccc", "|"))
}

func join(stringVals ...string) string {
	sep := stringVals[len(stringVals)-1]
	var joinString string
	for i := 0; i < len(stringVals)-1; i++ {
		joinString += stringVals[i] + sep
	}
	return joinString
}
