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

// BenchmarkPopCountShift-8 50000000 24.4 ns/op
func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(x)
	}
}
