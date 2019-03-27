package sexpr

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type (
	Symbol    string
	String    string
	Int       int
	StartList struct{}
	EndList   struct{}
)

func (tok Int) String() string       { return fmt.Sprintf("%d", tok) }
func (tok StartList) String() string { return "StartList" }
func (tok EndList) String() string   { return "EndList" }

type Token interface{}

type Decoder struct {
	scanner.Scanner
	depth int
}

func NewDecoder(r io.Reader) *Decoder {
	dec := &Decoder{scanner.Scanner{Mode: scanner.GoTokens}, 0}
	dec.Init(r)
	return dec
}

func (dec *Decoder) Token() (interface{}, error) {
	tok := dec.Scan()
	if dec.depth == 0 &&
		tok != '(' && tok != scanner.EOF {
		return nil, fmt.Errorf("expecting '(', got %s",
			scanner.TokenString(tok))
	}
	text := dec.TokenText()
	switch tok {
	case scanner.EOF:
		return nil, io.EOF
	case scanner.Ident:
		return Symbol(text), nil
	case scanner.String:
		return String(text[1 : len(text)-1]), nil
	case scanner.Int:
		i, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return nil, err
		}
		return Int(i), nil
	case '(':
		dec.depth++
		return StartList{}, nil
	case ')':
		dec.depth++
		return EndList{}, nil
	default:
		return nil, fmt.Errorf("unexpected token %q", text)
	}
}
