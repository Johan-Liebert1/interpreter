package interpreter

import (
	"programminglang/constants"
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

func (v VariableDeclaration) GetToken() types.Token {
	return types.Token{}
}
func (v VariableDeclaration) Scope(i *Interpreter) {
	typeName := v.TypeNode.GetToken().Value

	typeSymbol, _ := i.CurrentScope.LookupSymbol(typeName, false)

	variableName := v.VariableNode.GetToken().Value

	// helpers.ColorPrint(constants.Green, 1, v.VariableNode)
	// helpers.ColorPrint(constants.Green, 1, typeSymbol)

	if _, exists := i.CurrentScope.LookupSymbol(variableName, true); exists {
		// variable alreadyDeclaredVarName has already been declared
		i.CurrentScope.Error(
			constants.ERROR_DUPLICATE_ID,
			v.VariableNode.GetToken(),
		)
	}

	symbol := Symbol{
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

func (v VariableType) GetToken() types.Token {
	return v.Token
}
func (v VariableType) Scope(s *Interpreter) {}

func (v Variable) GetToken() types.Token {
	return v.Token
}
func (v Variable) Scope(i *Interpreter) {
	varName := v.Value
	_, exists := i.CurrentScope.LookupSymbol(varName, false)

	if !exists {
		i.CurrentScope.Error(
			constants.ERROR_ID_NOT_FOUND,
			v.GetToken(),
		)
	}

}
