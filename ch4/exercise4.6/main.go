package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	b := []byte("哈哈  哈 哈哈  a")
	b = replace(b)
	fmt.Printf("%s\n", b)
}

func replace(b []byte) []byte {
	for i := 0; i < len(b); {
		first, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(first) {
			second, _ := utf8.DecodeRune(b[i+size:])
			if unicode.IsSpace(second) {
				copy(b[i:], b[i+size:])
				b = b[:len(b)-size]
			}
		}
		i += size
	}
	return b
}
