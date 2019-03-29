package mandelbrot

import (
	"image/color"
	"testing"
)

func bench(b *testing.B, f func(complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrotComplex64(b *testing.B) {
	bench(b, mandelbrot64)
}

func BenchmarkMandelbrotComplex128(b *testing.B) {
	bench(b, mandelbrot128)
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	bench(b, mandelbrotBigFloat)
}

func BenchmarkMandelbrotBigRat(b *testing.B) {
	bench(b, mandelbrotBigRat)
}

/*
goos: darwin
goarch: amd64
pkg: github.com/linehk/gopl/ch3/exercise3.8
BenchmarkMandelbrotComplex64-8    	20000000	        57.9 ns/op
BenchmarkMandelbrotComplex128-8   	20000000	        56.5 ns/op
BenchmarkMandelbrotBigFloat-8     	 2000000	       590 ns/op
BenchmarkMandelbrotBigRat-8       	 1000000	      1334 ns/op
*/
