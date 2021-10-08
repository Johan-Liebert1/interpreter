package ast

import "programminglang/types"

type AbstractSyntaxTree interface {
	Op() types.Token
	LeftOperand() AbstractSyntaxTree
	RightOperand() AbstractSyntaxTree
}

type CompoundStatementNode interface {
	GetChildren() []AbstractSyntaxTree
}
