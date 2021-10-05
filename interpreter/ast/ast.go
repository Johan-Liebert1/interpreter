package ast

import "interpreter/types"

type AbstractSyntaxTree interface {
	TraverseTree(traversalType string)
	Op() types.Token
	LeftOperand() AbstractSyntaxTree
	RightOperand() AbstractSyntaxTree
}
