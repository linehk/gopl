package memo

import (
	"testing"
	"time"

	"github.com/linehk/gopl/ch9/exercise9.3/memotest"
)

var httpGetBody = memotest.HTTPGetBody

const timeout = 1 * time.Minute

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
