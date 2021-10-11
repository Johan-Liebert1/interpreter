package interpreter

import (
	"programminglang/types"
)

type Program struct {
	Declarations      []AbstractSyntaxTree
	CompoundStatement AbstractSyntaxTree
}

func (p Program) Op() types.Token {
	return types.Token{}
}
func (p Program) LeftOperand() AbstractSyntaxTree {
	return p
}
func (p Program) RightOperand() AbstractSyntaxTree {
	return p
}
func (p Program) Visit(i *Interpreter) {
	for _, decl := range p.Declarations {
		decl.Visit(i)
	}

	p.CompoundStatement.Visit(i)
}
