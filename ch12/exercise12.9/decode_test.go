package sexpr

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestToken(t *testing.T) {
	tests := []struct {
		input string
		want  []Token
	}{
		{`(3 "a" (a))`, []Token{
			StartList{},
			Int(3),
			String("a"),
			StartList{},
			Symbol("a"),
			EndList{},
			EndList{}},
		},
	}
	for i, tt := range tests {
		dec := NewDecoder(strings.NewReader(tt.input))
		var got []Token
		for {
			tok, err := dec.Token()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatal(err)
			}
			got = append(got, tok)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
