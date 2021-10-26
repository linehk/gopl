package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	// line map to filename slice
	nameList := make(map[string][]string)
	if len(files) == 0 {
		countLine(os.Stdin, counts, nameList)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLine(f, counts, nameList)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\t%s\n", n, nameList[line], line)
		}
	}
}

func contains(needle string, strs []string) bool {
	for _, s := range strs {
		if needle == s {
			return true
		}
	}
	return false
}

func countLine(f *os.File, counts map[string]int, nameList map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !contains(f.Name(), nameList[line]) {
			nameList[line] = append(nameList[line], f.Name())
		}
	}
}
