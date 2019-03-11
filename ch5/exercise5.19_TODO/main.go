package main

import "fmt"

func main() {
	a := returnN()
	fmt.Println(a)
}

func returnN() (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = p.(int)
		}
	}()
	panic(3)
}
