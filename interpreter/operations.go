package interpreter

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

type String struct {
	Token types.Token
	Value string
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

func (n IntegerNumber) GetToken() types.Token {
	return n.Token
}
func (in IntegerNumber) Scope(_ *Interpreter) {}

func (n FloatNumber) GetToken() types.Token {
	return n.Token
}
func (fn FloatNumber) Scope(_ *Interpreter) {}

func (s String) GetToken() types.Token {
	return s.Token
}
func (s String) Scope(_ *Interpreter) {}

func (b BinaryOperationNode) GetToken() types.Token {
	return b.Operation
}
func (b BinaryOperationNode) Scope(s *Interpreter) {
	b.Left.Scope(s)
	b.Right.Scope(s)
}

func (u UnaryOperationNode) GetToken() types.Token {
	return u.Operation
}
func (u UnaryOperationNode) Scope(s *Interpreter) {
	u.Operand.Scope(s)
}
