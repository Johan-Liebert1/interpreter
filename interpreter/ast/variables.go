package ast

import "programminglang/types"

type VariableDeclaration struct {
	VariableNode AbstractSyntaxTree
	TypeNode     AbstractSyntaxTree
}

type VariableType struct {
	Token types.Token
}

func (v VariableDeclaration) Op() types.Token {
	return types.Token{}
}
func (v VariableDeclaration) LeftOperand() AbstractSyntaxTree {
	return v.VariableNode
}
func (v VariableDeclaration) RightOperand() AbstractSyntaxTree {
	return v.TypeNode
}

func (v VariableType) Op() types.Token {
	return v.Token
}
func (v VariableType) LeftOperand() AbstractSyntaxTree {
	return v
}
func (v VariableType) RightOperand() AbstractSyntaxTree {
	return v
}
