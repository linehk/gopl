package main

import (
	"fmt"
)

func main() {
}

func max(vals ...int) (int, bool) {
	if len(vals) == 0 {
		fmt.Println("no arg!")
		return 0, false
	}
	max := vals[0]
	for _, i := range vals {
		if i > max {
			max = i
		}
	}
	return max, true
}

func min(vals ...int) (int, bool) {
	if len(vals) == 0 {
		fmt.Println("no arg!")
		return 0, false
	}
	min := vals[0]
	for _, i := range vals {
		if i < min {
			min = i
		}
	}
	return min, true
}

func atLeastMax(vals ...int) (int, bool) {
	if len(vals) < 1 {
		fmt.Println("no enough args!")
		return 0, false
	}
	max := vals[0]
	for _, i := range vals {
		if i > max {
			max = i
		}
	}
	return max, true
}

func atLeastMin(vals ...int) (int, bool) {
	if len(vals) < 1 {
		fmt.Println("no enough args!")
		return 0, false
	}
	min := vals[0]
	for _, i := range vals {
		if i < min {
			min = i
		}
	}
	return min, true
}
