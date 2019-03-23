package count

import (
	"io"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		r    io.Reader
		want count
	}{
		{strings.NewReader("abcd"), count{map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1},
			[utf8.UTFMax + 1]int{0, 4, 0, 0, 0}, 0, nil}},
	}
	for i, tt := range tests {
		got := charCount(tt.r)
		if !equal(got, tt.want) {
			t.Errorf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}

func equal(c1, c2 count) bool {
	if !(c1.counts == nil && c2.counts == nil ||
		c1.counts != nil && c2.counts != nil) {
		return false
	}
	for c, n := range c1.counts {
		if n != c2.counts[c] {
			return false
		}
	}
	if c1.utflen != c2.utflen {
		return false
	}
	if c1.invalid != c2.invalid {
		return false
	}
	if c1.err != c2.err {
		return false
	}
	return true
}
