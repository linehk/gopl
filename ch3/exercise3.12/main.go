package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("abcd", "dcba"))
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m := make(map[rune]int, len(s1))

	for _, v := range s1 {
		m[v]++
	}

	for _, v := range s2 {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
