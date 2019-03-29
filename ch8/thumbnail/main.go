// The thumbnail command produces thumbnails of JPEG files
// whose names are provided on each line of the standard input.
// go run gopl/ch8/thumbnail/main.go
// foo.jpeg
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/linehk/gopl/ch8/thumbnail/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
