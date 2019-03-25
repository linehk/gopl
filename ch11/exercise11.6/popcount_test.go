package popcount

import (
	"github.com/linehk/gopl/ch11/exercise11.6/bitcount"
	"github.com/linehk/gopl/ch11/exercise11.6/popcount"
	"testing"
)

const bin = 0x1234567890ABCDEF

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(bin)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitcount.BitCount(bin)
	}
}

func BenchmarkClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitcount.Clearing(bin)
	}
}

func BenchmarkShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitcount.Shifting(bin)
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/linehk/gopl/ch11/exercise11.6
BenchmarkPopCount-8   	2000000000	         0.65 ns/op
BenchmarkBitCount-8   	2000000000	         0.32 ns/op
BenchmarkClearing-8   	50000000	        30.7 ns/op
BenchmarkShifting-8   	30000000	        50.0 ns/op
*/
