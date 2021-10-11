package ast

import (
	"log"
	"programminglang/interpreter/symbols"
	"programminglang/types"
)

type VariableDeclaration struct {
	VariableNode AbstractSyntaxTree
	TypeNode     AbstractSyntaxTree
}

type VariableType struct {
	Token types.Token
}

type Variable struct {
	Token types.Token
	Value string
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
func (v VariableDeclaration) Visit(s *symbols.ScopedSymbolsTable) {
	typeName := v.TypeNode.Op().Value

	typeSymbol, _ := s.LookupSymbol(typeName)

	variableName := v.VariableNode.Op().Value

	alreadyDeclaredVarName, exists := s.LookupSymbol(variableName)

	if exists {
		// variable alreadyDeclaredVarName has already been declared
		log.Fatal("Error: Variable, ", alreadyDeclaredVarName, " has already been declared")
	}

	s.DefineSymbol(symbols.Symbol{
		Name: variableName,
		Type: typeSymbol.Name,
	})

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
func (v VariableType) Visit(s *symbols.ScopedSymbolsTable) {}

func (v Variable) Op() types.Token {
	return v.Token
}
func (v Variable) LeftOperand() AbstractSyntaxTree {
	return v
}
func (v Variable) RightOperand() AbstractSyntaxTree {
	return v
}
func (v Variable) Visit(s *symbols.ScopedSymbolsTable) {
	varName := v.Value
	_, exists := s.LookupSymbol(varName)

	if !exists {
		log.Fatal("Variable, ", varName, " is not defined")
	}

}
