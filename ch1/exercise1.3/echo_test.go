package echo

import (
	"testing"
	"time"
)

const n = 5000000

// 5000000 4291 ns/op 21.458202917s
func BenchmarkEcho1(b *testing.B) {
	b.N = n
	now := time.Now()
	for i := 0; i < b.N; i++ {
		echo1()
	}
	b.Log(time.Since(now))
}

// 5000000 4630 ns/op 23.153483183s
func BenchmarkEcho2(b *testing.B) {
	b.N = n
	now := time.Now()
	for i := 0; i < b.N; i++ {
		echo2()
	}
	b.Log(time.Since(now))
}

// 5000000 4742 ns/op 23.710918847s
func BenchmarkEcho3(b *testing.B) {
	b.N = n
	now := time.Now()
	for i := 0; i < b.N; i++ {
		echo3()
	}
	b.Log(time.Since(now))
}
