package main

import (
	"bufio"
	"fmt"
	"strings"
)

type wlCounter int

const exampleString = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"

func main() {
	var test wlCounter
	wc, err := test.wordCounter(exampleString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wc)

	lc, err := test.lineCounter(exampleString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lc)
}

func (wlc *wlCounter) wordCounter(s string) (int, error) {
	c := 0
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		c++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return c, nil
}

func (wlc *wlCounter) lineCounter(s string) (int, error) {
	c := 0
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		c++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return c, nil
}
