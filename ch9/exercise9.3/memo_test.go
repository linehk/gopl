package memo

import (
	"testing"
	"time"

	"github.com/linehk/gopl/ch9/exercise9.3/memotest"
)

var httpGetBody = memotest.HTTPGetBody

const timeout = 10 * time.Second

func TestSequential(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	done := make(chan struct{})
	go func() {
		time.Sleep(timeout)
		close(done)
	}()
	memotest.Sequential(t, m, done)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	done := make(chan struct{})
	go func() {
		time.Sleep(timeout)
		close(done)
	}()
	memotest.Concurrent(t, m, done)
}
