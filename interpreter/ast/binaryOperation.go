package ast

import (
	"interpreter/types"
)

type Number struct {
	Token types.Token
	Value int
}

type BinaryOperationNode struct {
	Left      AbstractSyntaxTree
	Operation types.Token
	Right     AbstractSyntaxTree
}

type UnaryOperationNode struct {
	Operation types.Token
	Operand   AbstractSyntaxTree
}

func (n Number) Op() types.Token {
	return n.Token
}
func (n Number) LeftOperand() AbstractSyntaxTree {
	return n
}
func (n Number) RightOperand() AbstractSyntaxTree {
	return n
}

func (b BinaryOperationNode) Op() types.Token {
	return b.Operation
}
func (b BinaryOperationNode) LeftOperand() AbstractSyntaxTree {
	return b.Left
}
func (b BinaryOperationNode) RightOperand() AbstractSyntaxTree {
	return b.Right
}

func (u UnaryOperationNode) Op() types.Token {
	return u.Operation
}
func (u UnaryOperationNode) LeftOperand() AbstractSyntaxTree {
	return u.Operand
}
func (u UnaryOperationNode) RightOperand() AbstractSyntaxTree {
	return u.Operand
}
