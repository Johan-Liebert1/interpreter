package ast

import (
	"programminglang/interpreter/symbols"
	"programminglang/types"
)

type AbstractSyntaxTree interface {
	Op() types.Token
	LeftOperand() AbstractSyntaxTree
	RightOperand() AbstractSyntaxTree
	Visit(s symbols.SymbolsTable)
}

type CompoundStatementNode interface {
	GetChildren() []AbstractSyntaxTree
}
