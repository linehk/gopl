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

func benchConcurrentRender(b *testing.B, procs int) {
	for i := 0; i < b.N; i++ {
		ConcurrentRender(procs)
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

// 最佳的值为 8
// 英特尔® 酷睿™ i7-6700HQ 处理器
// 单个芯片中有 4 核 8 线程
// 超线程技术：每个线程可以当做一个核

/*
goos: darwin
goarch: amd64
pkg: github.com/linehk/gopl/ch9/exercise9.6/mandelbrot
*/

// GOMAXPROCS=1 go test -bench=.
/*
BenchmarkSerialRender 	       5	 211424690 ns/op
Benchmark1            	       5	 209559894 ns/op
BenchmarkMaxProces    	       5	 210228271 ns/op
Benchmark8            	       5	 205241960 ns/op
Benchmark16           	       5	 205425923 ns/op
Benchmark32           	       5	 205986393 ns/op
Benchmark64           	       5	 205711365 ns/op
Benchmark128          	       5	 207079342 ns/op
*/

// GOMAXPROCS=2 go test -bench=.
/*
BenchmarkSerialRender-2   	       5	 210759586 ns/op
Benchmark1-2              	       5	 216224623 ns/op
BenchmarkMaxProces-2      	      10	 120278975 ns/op
Benchmark8-2              	      10	 119348429 ns/op
Benchmark16-2             	      10	 111547236 ns/op
Benchmark32-2             	      10	 110390066 ns/op
Benchmark64-2             	      10	 111594443 ns/op
Benchmark128-2            	      10	 111326037 ns/op
*/

// GOMAXPROCS=8 go test -bench=.
/*
BenchmarkSerialRender-8   	       5	 211322261 ns/op
Benchmark1-8              	       5	 209520325 ns/op
BenchmarkMaxProces-8      	      30	  50867360 ns/op
Benchmark8-8              	      30	  50840032 ns/op
Benchmark16-8             	      30	  51112422 ns/op
Benchmark32-8             	      30	  51369394 ns/op
Benchmark64-8             	      30	  52700661 ns/op
Benchmark128-8            	      30	  52493327 ns/op
*/

// GOMAXPROCS=16 go test -bench=.
/*
BenchmarkSerialRender-16    	       5	 208343022 ns/op
Benchmark1-16               	       5	 211780131 ns/op
BenchmarkMaxProces-16       	      30	  51611771 ns/op
Benchmark8-16               	      30	  51749592 ns/op
Benchmark16-16              	      30	  52042830 ns/op
Benchmark32-16              	      30	  57294224 ns/op
Benchmark64-16              	      20	  56774419 ns/op
Benchmark128-16             	      30	  57656986 ns/op
*/
