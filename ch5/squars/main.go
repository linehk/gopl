// The squares program demonstrates a function value with state.
package main

import (
	"fmt"
)

// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
