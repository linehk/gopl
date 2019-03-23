// Charcount computes counts of Unicode characters.
package count

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

type count struct {
	counts  map[rune]int         // counts of Unicode characters
	utflen  [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid int                  // count of invalid UTF-8 characters
	err     error
}

func charCount(r io.Reader) count {
	var c count
	c.counts = make(map[rune]int)
	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // return rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			c.err = err
		}
		if r == unicode.ReplacementChar && n == 1 {
			c.invalid++
			continue
		}
		c.counts[r]++
		c.utflen[n]++
	}
	return c
}
