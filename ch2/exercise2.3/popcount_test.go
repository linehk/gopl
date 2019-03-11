package popcount

import (
	"testing"
)

var x uint64 = 0x1234567890ABCDEF

// BenchmarkPopCount-8 2000000000 0.65 ns/op
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

// BenchmarkPopCountLoop-8 100000000 17.6 ns/op
func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(x)
	}
}
