package intset

import (
	"testing"
)

type interval struct {
	begin, end int
}

func initIntSet(i interval) *IntSet {
	s := &IntSet{}
	for x := i.begin; x <= i.end; x++ {
		s.Add(x)
	}
	return s
}

func initMapIntSet(i interval) map[int]bool {
	s := make(map[int]bool)
	for x := i.begin; x <= i.end; x++ {
		s[x] = true
	}
	return s
}

func equal(s1 *IntSet, s2 map[int]bool) bool {
	for k := range s2 {
		if !s1.Has(k) {
			return false
		}
	}
	return true
}

func TestIntSet(t *testing.T) {
	tests := []struct {
		i    interval
		mi   interval
		want bool
	}{
		{interval{0, 50}, interval{0, 50}, true},
		{interval{0, 100}, interval{0, 100}, true},
		{interval{0, 49}, interval{0, 50}, false},
	}
	for i, tt := range tests {
		s1 := initIntSet(tt.i)
		s2 := initMapIntSet(tt.mi)
		if got := equal(s1, s2); got != tt.want {
			t.Errorf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestIntSetUnionwith(t *testing.T) {
	tests := []struct {
		i1, i2   interval
		mi1, mi2 interval
		want     bool
	}{
		{interval{0, 50}, interval{50, 100},
			interval{0, 50}, interval{50, 100}, true},
		{interval{0, 50}, interval{50, 99},
			interval{0, 50}, interval{0, 100}, false},
	}
	for i, tt := range tests {
		s1, s2 := initIntSet(tt.i1), initIntSet(tt.i2)
		s1.Unionwith(s2)

		ms1, ms2 := initMapIntSet(tt.mi1), initMapIntSet(tt.mi2)
		for k := range ms2 {
			ms1[k] = true
		}

		if got := equal(s1, ms1); got != tt.want {
			t.Errorf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
