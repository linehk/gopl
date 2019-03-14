package treesort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestValues(t *testing.T) {
	tree := &tree{0, nil, nil}
	add(tree, 1)
	add(tree, 2)

	want := "012"
	got := tree.String()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
