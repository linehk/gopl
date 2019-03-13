package intset

import (
	"testing"
)

const N = 100

func initIntSet() *IntSet {
	s := &IntSet{}
	for x := 0; x < N; x++ {
		s.Add(x)
	}
	return s
}

func TestIntSet_Has(t *testing.T) {
	s := initIntSet()
	for x := 0; x < N; x++ {
		if !s.Has(x) {
			t.Fatalf("%d isn't exist", x)
		}
	}
}

func TestIntSet_Len(t *testing.T) {
	s := initIntSet()
	if got := s.Len(); got != N {
		t.Fatalf("got %d, want %d", got, N)
	}
}

func TestIntSet_Remove(t *testing.T) {
	s := initIntSet()
	x := 0
	want := !s.Has(x)
	s.Remove(x)
	if got := s.Has(x); got != want {
		t.Fatalf("didn't remove %v", x)
	}
}
