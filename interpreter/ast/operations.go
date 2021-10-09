package ast

import (
	"programminglang/types"
)

type IntegerNumber struct {
	Token types.Token
	Value int
}

type FloatNumber struct {
	Token types.Token
	Value float32
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

func (n IntegerNumber) Op() types.Token {
	return n.Token
}
func (n IntegerNumber) LeftOperand() AbstractSyntaxTree {
	return n
}
func (n IntegerNumber) RightOperand() AbstractSyntaxTree {
	return n
}

func (n FloatNumber) Op() types.Token {
	return n.Token
}
func (n FloatNumber) LeftOperand() AbstractSyntaxTree {
	return n
}
func (n FloatNumber) RightOperand() AbstractSyntaxTree {
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
