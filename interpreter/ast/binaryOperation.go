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

func (n Number) TraverseTree(traversalType string) {
}

func (n BinaryOperationNode) TraverseTree(traversalType string) {
}

func (n Number) Op() types.Token {
	return n.Token
}

func (n BinaryOperationNode) Op() types.Token {
	return n.Operation
}

func (n Number) LeftOperand() AbstractSyntaxTree {
	return n
}

func (n BinaryOperationNode) LeftOperand() AbstractSyntaxTree {
	return n.Left
}

func (n Number) RightOperand() AbstractSyntaxTree {
	return n
}

func (n BinaryOperationNode) RightOperand() AbstractSyntaxTree {
	return n.Right
}
