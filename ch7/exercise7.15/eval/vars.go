// 返回计算表达式所需的变量
package eval

func (v Var) Vars() []Var {
	return []Var{v}
}

func (literal) Vars() []Var {
	return []Var{}
}

func (u unary) Vars() []Var {
	return u.x.Vars()
}

func (b binary) Vars() []Var {
	return append(b.x.Vars(), b.y.Vars()...)
}

func (c call) Vars() []Var {
	var vars []Var
	for _, e := range c.args {
		vars = append(vars, e.Vars()...)
	}
	return vars
}
