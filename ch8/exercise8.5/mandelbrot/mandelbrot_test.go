package mandelbrot

import (
	"runtime"
	"testing"
)

func BenchmarkSerialRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SerialRender()
	}
}

func benchConcurrentRender(b *testing.B, workers int) {
	for i := 0; i < b.N; i++ {
		ConcurrentRender(workers)
	}
}

func Benchmark1(b *testing.B) {
	benchConcurrentRender(b, 1)
}

func BenchmarkMaxProces(b *testing.B) {
	benchConcurrentRender(b, runtime.GOMAXPROCS(-1))
}

func Benchmark8(b *testing.B) {
	benchConcurrentRender(b, 8)
}

func Benchmark16(b *testing.B) {
	benchConcurrentRender(b, 16)
}

func Benchmark32(b *testing.B) {
	benchConcurrentRender(b, 32)
}

func Benchmark64(b *testing.B) {
	benchConcurrentRender(b, 64)
}

func Benchmark128(b *testing.B) {
	benchConcurrentRender(b, 128)
}

// 设置跟 CPU 线程数一样多的 goroutine 最优
/*
goos: darwin
goarch: amd64
pkg: github.com/linehk/gopl/ch8/exercise8.5/mandelbrot
BenchmarkSerialRender-8   	       5	 212544422 ns/op
Benchmark1-8              	       5	 211501226 ns/op
BenchmarkMaxProces-8      	      30	  51970553 ns/op
Benchmark8-8              	      30	  50856941 ns/op
Benchmark16-8             	      30	  52236223 ns/op
Benchmark32-8             	      30	  53125410 ns/op
Benchmark64-8             	      30	  54051345 ns/op
Benchmark128-8            	      30	  55062813 ns/op
*/
