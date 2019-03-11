package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4}
	rotate(s, 2)
	fmt.Println(s)
}

func rotate(s []int, n int) {
	n %= len(s)
	tmp := append(s, s[:n]...)
	copy(s, tmp[n:])
}
