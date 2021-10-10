package ast

import (
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

/*
   def visit_VarDecl(self, node):
       type_name = node.type_node.value
       type_symbol = self.symtab.lookup(type_name)
       var_name = node.var_node.value
       var_symbol = VarSymbol(var_name, type_symbol)
       self.symtab.define(var_symbol)
*/

func (v VariableDeclaration) Op() types.Token {
	return types.Token{}
}
func (v VariableDeclaration) LeftOperand() AbstractSyntaxTree {
	return v.VariableNode
}
func (v VariableDeclaration) RightOperand() AbstractSyntaxTree {
	return v.TypeNode
}
func (v VariableDeclaration) Visit(s symbols.SymbolsTable) {
	typeName := v.TypeNode.Op().Value

	typeSymbol, _ := s.LookupSymbol(typeName)

	variableName := v.VariableNode.Op().Value

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
func (v VariableType) Visit(s symbols.SymbolsTable) {}
