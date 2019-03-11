package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type class string

const (
	letter  class = "letter"
	number  class = "number"
	graphic class = "graphic"
	space   class = "space"
	symbol  class = "symbol"
)

func main() {
	classCount := make(map[class]int, 5)
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune() // return rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		switch {
		case unicode.IsLetter(r):
			classCount[letter]++
		case unicode.IsNumber(r):
			classCount[number]++
		case unicode.IsGraphic(r):
			classCount[graphic]++
		case unicode.IsSpace(r):
			classCount[space]++
		case unicode.IsSymbol(r):
			classCount[symbol]++
		}
	}
	for class, count := range classCount {
		fmt.Printf("class: %s, count = %d\n", class, count)
	}
}
