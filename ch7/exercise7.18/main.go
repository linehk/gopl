package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func main() {
	node, err := parse(xml.NewDecoder(os.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", node)
}

type Node interface {
	String() string
}

type CharData string

func (c CharData) String() string { return string(c) }

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	var attrs, children string
	for _, attr := range e.Attr {
		attrs += fmt.Sprintf(" %s=%q", attr.Name.Local, attr.Value)
	}
	for _, child := range e.Children {
		children += child.String()
	}
	return fmt.Sprintf("<%s%s>%s</%s>",
		e.Type.Local, attrs, children, e.Type.Local)
}

func parse(dec *xml.Decoder) (Node, error) {
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			e := &Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				p.Children = append(p.Children, e)
			}
			stack = append(stack, e)
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected tag closing")
			} else if len(stack) == 1 {
				return stack[0], nil
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				p.Children = append(p.Children, CharData(tok))
			}
		}
	}
}
