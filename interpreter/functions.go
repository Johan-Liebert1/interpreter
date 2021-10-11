package interpreter

import (
	"fmt"
	"programminglang/constants"
	"programminglang/interpreter/symbols"
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

func (fn FunctionDeclaration) Visit(i *Interpreter) {
	funcName := fn.FunctionName

	funcSymbol := symbols.Symbol{
		Name: funcName,
		Type: constants.FUNCTION_TYPE,
	}

	fmt.Println("Entering Scope, ", funcName)

	funcScope := symbols.ScopedSymbolsTable{
		CurrentScopeName:  funcName,
		CurrentScopeLevel: i.CurrentScope.CurrentScopeLevel + 1,
		EnclosingScope:    i.CurrentScope,
	}

	funcScope.Init()
	i.CurrentScope = &funcScope

	defer i.ReleaseScope()

	for _, param := range fn.FormalParameters {
		paramName := param.VariableNode.Op().Value

		// this is going to be a built in type so it will definitely exist
		paramType := param.TypeNode.Op().Value

		paramSymbol := symbols.Symbol{
			Name: paramName,
			Type: paramType,
		}

		i.CurrentScope.DefineSymbol(paramSymbol)

		// add all the parameters symbols
		funcSymbol.ParamSymbols = append(funcSymbol.ParamSymbols, paramSymbol)
	}

	fn.FunctionBlock.Visit(i)

	fmt.Println("Exit Scope, ", funcName)

}
