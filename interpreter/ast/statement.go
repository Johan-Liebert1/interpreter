package ast

import (
	"log"
	"programminglang/interpreter/symbols"
	"programminglang/types"
)

type CompoundStatement struct {
	Token    types.Token
	Children []AbstractSyntaxTree
}

type AssignmentStatement struct {
	Left  AbstractSyntaxTree
	Token types.Token
	Right AbstractSyntaxTree
}

type BlankStatement struct {
	Token types.Token
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
func (cs CompoundStatement) Visit(s *symbols.ScopedSymbolsTable) {
	for _, child := range cs.Children {
		child.Visit(s)
	}
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
func (as AssignmentStatement) Visit(s *symbols.ScopedSymbolsTable) {
	variableName := as.Left.Op().Value
	_, exists := s.LookupSymbol(variableName)

	if !exists {
		log.Fatal("AssignmentStatement, ", variableName, " is not defined")
	}

	as.Right.Visit(s)
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
func (bs BlankStatement) Visit(_ *symbols.ScopedSymbolsTable) {}
