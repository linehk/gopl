package main

import (
	"fmt"
)

func main() {
	fmt.Println(returnNoZero())
}

// 原理：recover 后不会继续执行，而是直接调用 return
func returnNoZero() (result int) {
	defer func() {
		result = 3
		_ = recover()
	}()
	panic("panic!")
}
