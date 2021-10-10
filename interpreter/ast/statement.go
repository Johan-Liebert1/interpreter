package ast

import (
	"log"
	"programminglang/constants"
	"programminglang/interpreter/symbols"
	"programminglang/types"
)

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
func (p Program) Visit(s *symbols.SymbolsTable) {
	for _, decl := range p.Declarations {
		decl.Visit(s)
	}

	p.CompoundStatement.Visit(s)
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
func (cs CompoundStatement) Visit(s *symbols.SymbolsTable) {
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
func (as AssignmentStatement) Visit(s *symbols.SymbolsTable) {
	variableName := as.Left.Op().Value
	val, exists := s.LookupSymbol(variableName)

	if !exists {
		constants.SpewPrinter.Dump(s)
		log.Fatal("AssignmentStatement, ", variableName, " is not defined. Value = ", val)
	}

	as.Right.Visit(s)
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
func (v Variable) Visit(s *symbols.SymbolsTable) {
	varName := v.Value
	_, exists := s.LookupSymbol(varName)

	if !exists {
		log.Fatal("Variable, ", varName, " is not defined")
	}

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
func (bs BlankStatement) Visit(_ *symbols.SymbolsTable) {}
