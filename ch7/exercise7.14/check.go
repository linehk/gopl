package eval

import (
	"fmt"
	"strings"
)

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (t ternary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("?:", t.op1) {
		return fmt.Errorf("unexpected ternary op1 %q", t.op1)
	} else if !strings.ContainsRune("?:", t.op2) {
		return fmt.Errorf("unexpected ternary op2 %q", t.op2)
	}

	if err := t.x.Check(vars); err != nil {
		return err
	} else if err := t.y.Check(vars); err != nil {
		return err
	}
	return t.z.Check(vars)
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
