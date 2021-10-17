package interpreter

import (
	"fmt"
	"programminglang/constants"
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
	Token            types.Token // IDENTIFIER token for the function name
	FunctionSymbol   Symbol
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

func (fn FunctionDeclaration) Scope(i *Interpreter) {
	funcName := fn.FunctionName

	funcSymbol := Symbol{
		Name: funcName,
		Type: constants.FUNCTION_TYPE,
	}

	// used by the interpreter when executing the function
	funcSymbol.FunctionBlock = fn.FunctionBlock

	fmt.Println("Entering Scope, ", funcName)

	// helpers.ColorPrint(
	// 	constants.Blue, 2,
	// 	"current function scope ", funcScope,
	// 	"\nglobal scope ", funcScope.EnclosingScope
	// )

	for _, param := range fn.FormalParameters {
		paramName := param.VariableNode.Op().Value

		// this is going to be a built in type so it will definitely exist
		paramType := param.TypeNode.Op().Value

		paramSymbol := Symbol{
			Name: paramName,
			Type: paramType,
		}

		i.CurrentScope.DefineSymbol(paramSymbol)

		// add all the parameters symbols
		funcSymbol.ParamSymbols = append(funcSymbol.ParamSymbols, paramSymbol)
	}

	i.CurrentScope.DefineSymbol(funcSymbol)

	funcScope := ScopedSymbolsTable{
		CurrentScopeName:  funcName,
		CurrentScopeLevel: i.CurrentScope.CurrentScopeLevel + 1,
		EnclosingScope:    i.CurrentScope,
	}

	funcScope.Init()
	i.CurrentScope = &funcScope
	defer i.ReleaseScope()

	fn.FunctionBlock.Scope(i)

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

func (fn FunctionParameters) Scope(i *Interpreter) {}

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

func (fn FunctionCall) Scope(i *Interpreter) {
	for _, paramNode := range fn.ActualParameters {
		paramNode.Scope(i)
	}

	// funcSymbol, _ := i.CurrentScope.LookupSymbol(fn.FunctionName, false)

	// accessed by the interpreter when executing procedure call
	// fn.FunctionSymbol = funcSymbol
}
