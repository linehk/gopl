package main

import (
	"fmt"
)

func main() {
	foo := "foo"
	fmt.Println(expand(foo, replace))
}

func expand(s string, f func(string) string) string {
	return f(s)
}

func replace(s string) string {
	return s + "-next"
}
