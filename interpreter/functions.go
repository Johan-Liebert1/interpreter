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

type FunctionCall struct {
	FunctionName     string
	ActualParameters []AbstractSyntaxTree
	Token            types.Token
}

// function declaration

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

// function parameters

func (fn FunctionParameters) Op() types.Token {
	return types.Token{}
}

func (fn FunctionParameters) LeftOperand() AbstractSyntaxTree {
	return fn.VariableNode
}

func (fn FunctionParameters) RightOperand() AbstractSyntaxTree {
	return fn.TypeNode
}

func (fn FunctionParameters) Visit(i *Interpreter) {}

// function call

func (fn FunctionCall) Op() types.Token {
	return types.Token{}
}

func (fn FunctionCall) LeftOperand() AbstractSyntaxTree {
	return fn.ActualParameters[0]
}

func (fn FunctionCall) RightOperand() AbstractSyntaxTree {
	return fn.ActualParameters[0]
}

func (fn FunctionCall) Visit(i *Interpreter) {
	for _, paramNode := range fn.ActualParameters {
		paramNode.Visit(i)
	}
}
