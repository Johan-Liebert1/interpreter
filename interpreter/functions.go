package interpreter

import (
	"programminglang/types"
)

type FunctionParameters struct {
	VariableNode AbstractSyntaxTree
	TypeNode     AbstractSyntaxTree
}

type FunctionDeclaration struct {
	FunctionName     string
	FunctionBlock    AbstractSyntaxTree
	FormalParameters []FunctionParameters
}

func (fn FunctionDeclaration) Op() types.Token {
	return types.Token{}
}

func (fn FunctionDeclaration) LeftOperand() AbstractSyntaxTree {
	return fn.FunctionBlock
}

func (fn FunctionDeclaration) RightOperand() AbstractSyntaxTree {
	return fn.FunctionBlock
}

func (fn FunctionDeclaration) Visit(i *Interpreter) {}
