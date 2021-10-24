package interpreter

import (
	"fmt"
	"programminglang/constants"
	"programminglang/interpreter/errors"
	"programminglang/types"
)

type FunctionParameters struct {
	VariableNode AbstractSyntaxTree // a Variable struct
	TypeNode     AbstractSyntaxTree // a VariableType struct
}

type FunctionDeclaration struct {
	FunctionName     string
	FunctionBlock    AbstractSyntaxTree // a Program struct
	FormalParameters []FunctionParameters
	ReturningValue   AbstractSyntaxTree
}

type FunctionCall struct {
	FunctionName     string
	ActualParameters []AbstractSyntaxTree
	Token            types.Token // IDENTIFIER token for the function name
	FunctionSymbol   Symbol
}

// function declaration

func (fn FunctionDeclaration) GetToken() types.Token {
	return types.Token{}
}

func (fn FunctionDeclaration) Scope(i *Interpreter) {
	funcName := fn.FunctionName

	funcSymbol := Symbol{
		Name: funcName,
		Type: constants.FUNCTION_TYPE,
	}

	// used by the interpreter when executing the function
	funcSymbol.FunctionBlock = fn.FunctionBlock

	funcScope := ScopedSymbolsTable{
		CurrentScopeName:  funcName,
		CurrentScopeLevel: i.CurrentScope.CurrentScopeLevel + 1,
		EnclosingScope:    i.CurrentScope,
	}

	funcScope.Init()
	i.CurrentScope = &funcScope
	defer i.ReleaseScope()

	// fmt.Println("Entering Scope, ", funcName)

	// helpers.ColorPrint(
	// 	constants.Blue, 2,
	// 	"current function scope ", funcScope,
	// 	"\nglobal scope ", funcScope.EnclosingScope
	// )

	for _, param := range fn.FormalParameters {
		paramName := param.VariableNode.GetToken().Value

		// this is going to be a built in type so it will definitely exist
		paramType := param.TypeNode.GetToken().Value

		paramSymbol := Symbol{
			Name: paramName,
			Type: paramType,
		}

		i.CurrentScope.DefineSymbol(paramSymbol)

		// add all the parameters symbols
		funcSymbol.ParamSymbols = append(funcSymbol.ParamSymbols, paramSymbol)
	}

	funcSymbol.ReturningValue = fn.ReturningValue

	// we've already created a new scope, so need to add the funcSymbol to the enclosing scope
	i.CurrentScope.EnclosingScope.DefineSymbol(funcSymbol)

	fn.FunctionBlock.Scope(i)

	// fmt.Println("Exit Scope, ", funcName)

}

// function parameters

func (fn FunctionParameters) GetToken() types.Token {
	return types.Token{}
}

func (fn FunctionParameters) Scope(i *Interpreter) {}

// function call

func (fn FunctionCall) GetToken() types.Token {
	return types.Token{}
}

func (fn FunctionCall) Scope(i *Interpreter) {
	_, exists := i.CurrentScope.LookupSymbol(fn.FunctionName, false)

	if !exists {
		errors.ShowError(
			constants.SEMANTIC_ERROR,
			constants.ERROR_VARAIBLE_NOT_DEFINED,
			fmt.Sprintf("Function %s is not defined", fn.FunctionName),
			fn.Token,
		)
	}

	for _, paramNode := range fn.ActualParameters {
		paramNode.Scope(i)
	}
}
