package popcount

import (
	"testing"
)

var x uint64 = 0x1234567890ABCDEF

// BenchmarkPopCount-8 2000000000 0.32 ns/op
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

// BenchmarkPopCountClean-8 100000000 16.9 ns/op
func BenchmarkPopCountClean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClean(x)
	}
}
