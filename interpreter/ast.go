package interpreter

import (
	"programminglang/types"
)

type AbstractSyntaxTree interface {
	Op() types.Token
	LeftOperand() AbstractSyntaxTree
	RightOperand() AbstractSyntaxTree
	Scope(i *Interpreter)
	// EvaluateNode() float32
}

type CompoundStatementNode interface {
	GetChildren() []AbstractSyntaxTree
}
