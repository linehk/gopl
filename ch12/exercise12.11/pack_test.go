package pack

import (
	"testing"
)

func TestPack(t *testing.T) {
	type data struct {
		Labels    []string `http:"l"`
		MaxResult int      `http:"max"`
		Exact     bool     `http:"x"`
	}

	tests := []struct {
		data data
		want string
	}{
		{data{MaxResult: 10}, "?l=&max=10&x=false"},
	}
	for i, tt := range tests {
		if got := Pack(tt.data); got != tt.want {
			t.Errorf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
