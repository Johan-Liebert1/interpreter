package interpreter

import (
	"programminglang/types"
)

type AbstractSyntaxTree interface {
	GetToken() types.Token
	Scope(i *Interpreter)
	// EvaluateNode() float32
}

type CompoundStatementNode interface {
	GetChildren() []AbstractSyntaxTree
}
