package ast

import (
	"programminglang/interpreter/symbols"
	"programminglang/types"
)

type AbstractSyntaxTree interface {
	Op() types.Token
	LeftOperand() AbstractSyntaxTree
	RightOperand() AbstractSyntaxTree
	Visit(s *symbols.ScopedSymbolsTable)
	// EvaluateNode() float32
}

type CompoundStatementNode interface {
	GetChildren() []AbstractSyntaxTree
}
