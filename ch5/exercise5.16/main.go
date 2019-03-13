package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("aaa", "bbb", "ccc", "|"))
}

func join(strs ...string) string {
	if len(strs) < 2 {
		return ""
	}

	sep := strs[len(strs)-1]
	last := strs[len(strs)-2]
	// 为了不改变输入
	tempStrs := strs[:len(strs)-2]

	var b strings.Builder
	for _, s := range tempStrs {
		b.WriteString(s)
		b.WriteString(sep)
	}
	b.WriteString(last)

	return b.String()
}
