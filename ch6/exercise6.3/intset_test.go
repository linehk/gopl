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

func initIntSetByRange(b1, e1, b2, e2 int) *IntSet {
	s := &IntSet{}
	for x := b1; x <= e1; x++ {
		s.Add(x)
	}
	for x := b2; x <= e2; x++ {
		s.Add(x)
	}
	return s
}

func TestIntSet_IntersectWith(t *testing.T) {
	tests := []struct {
		s    *IntSet
		t    *IntSet
		want *IntSet
	}{
		// 相同
		{initIntSet(0, 100), initIntSet(0, 100), initIntSet(0, 100)},

		{initIntSet(0, 100), initIntSet(0, 50), initIntSet(0, 50)},
		{initIntSet(0, 50), initIntSet(0, 100), initIntSet(0, 50)},

		{initIntSet(0, 75), initIntSet(50, 100), initIntSet(50, 75)},
	}
	for i, tt := range tests {
		tt.s.IntersectWith(tt.t)
		if !equal(tt.s.words, tt.want.words) {
			t.Fatalf("%d. got %v, want %v\n", i, tt.s.words, tt.want.words)
		}
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	tests := []struct {
		s    *IntSet
		t    *IntSet
		want *IntSet
	}{
		// 相同
		{initIntSet(0, 100), initIntSet(0, 100), &IntSet{}},

		{initIntSet(0, 100), initIntSet(0, 50), initIntSet(51, 100)},
		{initIntSet(0, 50), initIntSet(0, 100), &IntSet{}},

		{initIntSet(0, 50), initIntSet(25, 100), initIntSet(0, 24)},
		{initIntSet(25, 100), initIntSet(0, 50), initIntSet(51, 100)},
	}
	for i, tt := range tests {
		tt.s.DifferenceWith(tt.t)
		if !equal(tt.s.words, tt.want.words) {
			t.Fatalf("%d. got %v, want %v\n", i, tt.s.words, tt.want.words)
		}
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	tests := []struct {
		s    *IntSet
		t    *IntSet
		want *IntSet
	}{
		// 相同
		{initIntSet(0, 100), initIntSet(0, 100), &IntSet{}},

		{initIntSet(0, 100), initIntSet(0, 50), initIntSet(51, 100)},
		{initIntSet(0, 50), initIntSet(0, 100), initIntSet(51, 100)},

		{initIntSet(0, 50), initIntSet(25, 100), initIntSetByRange(0, 24, 51, 100)},
		{initIntSet(25, 100), initIntSet(0, 50), initIntSetByRange(0, 24, 51, 100)},
	}
	for i, tt := range tests {
		tt.s.SymmetricDifference(tt.t)
		if !equal(tt.s.words, tt.want.words) {
			t.Fatalf("%d. got %v, want %v", i, tt.s.words, tt.want.words)
		}
	}
}

func equal(s1, s2 []uint64) bool {
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
