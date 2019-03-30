package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	_, err := html.Parse(NewReader("<h1>Hello</h1>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse err: %v", err)
		os.Exit(1)
	}
}

type StringReader string

func (s *StringReader) Read(p []byte) (int, error) {
	copy(p, *s)
	return len(*s), io.EOF
}

func NewReader(s string) io.Reader {
	sr := StringReader(s)
	return &sr
}
