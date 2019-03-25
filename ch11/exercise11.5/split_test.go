package split_test

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
	}
	for i, tt := range tests {
		if got := strings.Split(tt.s, tt.sep); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d. Split(%q, %q) got %v, want %v", i, tt.s, tt.sep, got, tt.want)
			t.Errorf("%d. Split(%q, %q) returned %d words, want %d",
				i, tt.s, tt.sep, len(got), len(tt.want))
		}
	}
}
