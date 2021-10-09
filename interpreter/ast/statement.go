package ast

import "programminglang/types"

type Program struct {
	Declarations      []AbstractSyntaxTree
	CompoundStatement AbstractSyntaxTree
}

type CompoundStatement struct {
	Token    types.Token
	Children []AbstractSyntaxTree
}

type AssignmentStatement struct {
	Left  AbstractSyntaxTree
	Token types.Token
	Right AbstractSyntaxTree
}

type Variable struct {
	Token types.Token
	Value string
}

type BlankStatement struct {
	Token types.Token
}

func (p Program) Op() types.Token {
	return types.Token{}
}
func (p Program) LeftOperand() AbstractSyntaxTree {
	return p
}
func (p Program) RightOperand() AbstractSyntaxTree {
	return p
}

func (cs CompoundStatement) Op() types.Token {
	return cs.Token
}
func (cs CompoundStatement) LeftOperand() AbstractSyntaxTree {
	return cs.Children[0]
}
func (cs CompoundStatement) RightOperand() AbstractSyntaxTree {
	return cs.Children[0]
}
func (cs CompoundStatement) GetChildren() []AbstractSyntaxTree {
	return cs.Children
}

func (v AssignmentStatement) Op() types.Token {
	return v.Token
}
func (v AssignmentStatement) LeftOperand() AbstractSyntaxTree {
	return v.Left
}
func (v AssignmentStatement) RightOperand() AbstractSyntaxTree {
	return v.Right
}

func (v Variable) Op() types.Token {
	return v.Token
}
func (v Variable) LeftOperand() AbstractSyntaxTree {
	return v
}
func (v Variable) RightOperand() AbstractSyntaxTree {
	return v
}

func (bs BlankStatement) Op() types.Token {
	return bs.Token
}
func (bs BlankStatement) LeftOperand() AbstractSyntaxTree {
	return bs
}
func (bs BlankStatement) RightOperand() AbstractSyntaxTree {
	return bs
}
