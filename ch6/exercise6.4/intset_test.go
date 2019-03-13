package intset

import (
	"testing"
)

func initIntSet(begin, end int) *IntSet {
	s := &IntSet{}
	for x := begin; x <= end; x++ {
		s.Add(x)
	}
	return s
}

func initElems(begin, end int) []int {
	var elems []int
	for x := begin; x <= end; x++ {
		elems = append(elems, x)
	}
	return elems
}

func TestIntSet_Elems(t *testing.T) {
	s := initIntSet(0, 100)

	got := s.Elems()
	want := initElems(0, 100)

	if !equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
