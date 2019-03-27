package sexpr

import (
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	tests := []struct {
		v    bool
		want string
	}{
		{true, "t"},
		{false, "nil"},
	}
	for i, tt := range tests {
		got, err := Marshal(tt.v)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != tt.want {
			t.Errorf("%d. got %s, want %s", i, got, tt.want)
		}
	}
}

func TestFloat(t *testing.T) {
	tests := []struct {
		v    float64
		want string
	}{
		{3.14, fmt.Sprintf("%g", 3.14)},
		{0, fmt.Sprintf("%g", 0.0)},
	}
	for i, tt := range tests {
		got, err := Marshal(tt.v)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != tt.want {
			t.Errorf("%d. got %s, want %s", i, got, tt.want)
		}
	}
}

func TestComplex(t *testing.T) {
	tests := []struct {
		v    complex128
		want string
	}{
		{1 + 2i, "#C(1 2)"},
		{0 + 2i, "#C(0 2)"},
		{1 + 0i, "#C(1 0)"},
	}
	for i, tt := range tests {
		got, err := Marshal(tt.v)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != tt.want {
			t.Errorf("%d. got %s, want %s", i, got, tt.want)
		}
	}
}

func TestInterface(t *testing.T) {
	type w struct {
		s interface{}
	}
	i := w{[]int{1, 2, 3}}
	want := `((s ("[]int" (1 2 3))))`
	got, err := Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
