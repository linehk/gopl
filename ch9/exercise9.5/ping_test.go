package ping

import (
	"testing"
)

func BenchmarkPing(b *testing.B) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	done := make(chan struct{})
	go func() {
		for i := 0; i < b.N; i++ {
			ch1 <- "ping"
			<-ch2
		}
		done <- struct{}{}
	}()

	go func() {
		for i := 0; i < b.N; i++ {
			<-ch1
			ch2 <- "pong"
		}
		done <- struct{}{}
	}()

	<-done
	<-done

	close(ch1)
	close(ch2)
	close(done)
}

/*
goos: darwin
goarch: amd64
pkg: github.com/linehk/gopl/ch9/exercise9.5
BenchmarkPing-8   	 3000000	       506 ns/op
*/
