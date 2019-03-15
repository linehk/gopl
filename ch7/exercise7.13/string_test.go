package eval

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{"-1 + -x", "(-1 + -x)"},
		{"-1 - x", "(-1 - x)"},
		{"sqrt(A / pi)", "sqrt((A / pi))"},
		{"pow(x, 3) + pow(y, 3)", "(pow(x, 3) + pow(y, 3))"},
		{"5 / 9 * (F - 32)", "((5 / 9) * (F - 32))"},
	}
	for i, tt := range tests {
		expr, err := Parse(tt.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := expr.String()
		if err != nil {
			t.Error(err)
			continue
		}

		if got != tt.want {
			t.Fatalf("%d. got %v, expr %v", i, got, tt.want)
		}
	}
}
