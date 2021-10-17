package interpreter

import (
	"programminglang/constants"
	"programminglang/interpreter/symbols"
	"programminglang/types"
)

type VariableDeclaration struct {
	VariableNode AbstractSyntaxTree // a Variable struct
	TypeNode     AbstractSyntaxTree // a VariableType struct
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
func (v VariableDeclaration) Scope(i *Interpreter) {
	typeName := v.TypeNode.Op().Value

	typeSymbol, _ := i.CurrentScope.LookupSymbol(typeName, false)

	variableName := v.VariableNode.Op().Value

	// helpers.ColorPrint(constants.Green, 1, v.VariableNode)
	// helpers.ColorPrint(constants.Green, 1, typeSymbol)

	if _, exists := i.CurrentScope.LookupSymbol(variableName, true); exists {
		// variable alreadyDeclaredVarName has already been declared
		i.CurrentScope.Error(
			constants.ERROR_DUPLICATE_ID,
			v.VariableNode.Op(),
		)
	}

	symbol := symbols.Symbol{
		Name: variableName,
		Type: typeSymbol.Name,
	}

	// helpers.ColorPrint(
	// 	constants.Green, 1,
	// 	"defining symbol", symbol,
	// 	"\nin scope ", i.CurrentScope,
	// 	"\ncurrent scope address", &i.CurrentScope,
	// )

	i.CurrentScope.DefineSymbol(symbol)

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
func (v VariableType) Scope(s *Interpreter) {}

func (v Variable) Op() types.Token {
	return v.Token
}
func (v Variable) LeftOperand() AbstractSyntaxTree {
	return v
}
func (v Variable) RightOperand() AbstractSyntaxTree {
	return v
}
func (v Variable) Scope(i *Interpreter) {
	varName := v.Value
	_, exists := i.CurrentScope.LookupSymbol(varName, false)

	if !exists {
		i.CurrentScope.Error(
			constants.ERROR_ID_NOT_FOUND,
			v.Op(),
		)
	}

}
