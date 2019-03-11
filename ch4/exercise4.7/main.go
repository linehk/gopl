package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("一 二 三")
	reverseUTF8(b)
	fmt.Printf("%s\n", b)
}

func reverseUTF8(b []byte) {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
}

func reverse(b []byte) {
	last := len(b) - 1
	for i := 0; i < len(b)/2; i++ {
		b[i], b[last-i] = b[last-i], b[i]
	}
}
