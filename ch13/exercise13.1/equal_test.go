package equal

import (
	"testing"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		x, y interface{}
		want bool
	}{
		{1.0, 1.0000000009, true},
		{1.0, 1.000000001, false},
		{1.0, 1.0000000011, false},
	}
	for i, tt := range tests {
		if got := Equal(tt.x, tt.y); got != tt.want {
			t.Errorf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
