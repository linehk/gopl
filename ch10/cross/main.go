// The cross command prints the values of GOOS and GOARCH for this target.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
