package ast

import (
	"programminglang/interpreter/symbols"
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
func (in IntegerNumber) Visit(_ *symbols.SymbolsTable) {}

func (in IntegerNumber) EvaluateNode() float32 {
	return float32(in.Token.IntegerValue)
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
func (fn FloatNumber) Visit(_ *symbols.SymbolsTable) {}

func (fn FloatNumber) EvaluateNode() float32 {
	return fn.Token.FloatValue
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
func (b BinaryOperationNode) Visit(s *symbols.SymbolsTable) {
	b.Left.Visit(s)
	b.Right.Visit(s)
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
func (u UnaryOperationNode) Visit(s *symbols.SymbolsTable) {
	u.Operand.Visit(s)
}

// func (u *UnaryOperationNode) EvaluateNode() float32 {
// 	if u.Operation.Type == constants.PLUS {
// 		return +u.EvaluateNode()

// 	} else if u.Operation.Type == constants.MINUS {
// 		return -u.EvaluateNode()
// 	}
// }
