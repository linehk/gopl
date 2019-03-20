package mandelbrot

import (
	"runtime"
	"testing"
)

//goos: darwin
//goarch: amd64
//pkg: github.com/linehk/gopl/ch8/exercise8.5/mandelbrot
//BenchmarkSerialRender-8          	       5	 234825068 ns/op
//BenchmarkConcurrentRender1-8     	       5	 238189873 ns/op
//BenchmarkConcurrentRender8-8     	      30	  51880934 ns/op
//BenchmarkConcurrentRender100-8   	      20	  53125191 ns/op

func BenchmarkSerialRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SerialRender()
	}
}

func BenchmarkConcurrentRender1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentRender(1)
	}
}

func BenchmarkConcurrentRender8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentRender(runtime.GOMAXPROCS(-1))
	}
}

func BenchmarkConcurrentRender100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentRender(100)
	}
}
