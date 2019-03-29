package cycle

import (
	"testing"
)

func TestCycle(t *testing.T) {
	type node struct {
		value string
		next  *node
	}

	cycle1 := &node{"cycle1", nil}
	cycle2 := &node{"cycle2", nil}
	cycle3 := &node{"cycle3", nil}
	cycle1.next = cycle2
	cycle2.next = cycle3
	cycle3.next = cycle1

	if got, want := Cycle(cycle1), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	line := &node{"line1", &node{"line2", &node{"line3", nil}}}
	if got, want := Cycle(line), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
