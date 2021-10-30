package interpreter

import (
	"fmt"
	"programminglang/constants"
	"programminglang/interpreter/errors"
	"programminglang/types"
)

type Symbol struct {
	Name     string // name of the identifier / symbol
	Category string // whether the symbol is a built in type, or a variable, or a function name
	Type     string // integer, float, string, etc

	ParamSymbols   []Symbol           // all the parameter symbols for functions
	FunctionBlock  AbstractSyntaxTree // the function's block (executable) code
	ReturningValue AbstractSyntaxTree
}

type ScopedSymbolsTable struct {
	CurrentScopeName  string
	CurrentScopeLevel int
	EnclosingScope    *ScopedSymbolsTable
	SymbolTable       map[string]Symbol
}

/*
	Allocate memory for a SymbolTable and add some predefined symbols
*/
func (s *ScopedSymbolsTable) Init() {
	s.SymbolTable = map[string]Symbol{}

	// initialize some built in types
	s.DefineSymbol(Symbol{
		Name: constants.INTEGER_TYPE,
		Type: constants.BUILT_IN_TYPE,
	})

	s.DefineSymbol(Symbol{
		Name: constants.FLOAT_TYPE,
		Type: constants.BUILT_IN_TYPE,
	})

	s.DefineSymbol(Symbol{
		Name: constants.STRING_TYPE,
		Type: constants.BUILT_IN_TYPE,
	})

	s.DefineSymbol(Symbol{
		Name: constants.BOOLEAN_TYPE,
		Type: constants.BUILT_IN_TYPE,
	})

	s.DefineSymbol(Symbol{
		Name: constants.PRINT_OUTPUT,
		Type: constants.FUNCTION_TYPE,
	})

}

/*
	Receive a symbol struct and add to hash map with key as the symbol's name and value as the symbol
*/
func (s *ScopedSymbolsTable) DefineSymbol(symbol Symbol) {
	s.SymbolTable[symbol.Name] = symbol
}

func (s *ScopedSymbolsTable) LookupSymbol(symbolName string, currentScopeOnly bool) (Symbol, bool) {
	value, ok := s.SymbolTable[symbolName]

	if !ok && s.EnclosingScope != s && !currentScopeOnly {
		// variable not found in current scope, check in the parent scope
		// only check if the parent scope is not itself (case for global scope)
		return s.EnclosingScope.LookupSymbol(symbolName, currentScopeOnly)
	}

	// if only checking in current scope then just return
	// need this check to prevent duplicate declaration in the current scope, but duplicate
	// delclarations in the outer scopes are perfectly fine
	return value, ok
}

func (s *ScopedSymbolsTable) Error(errorCode string, token types.Token) {
	errors.ShowError(
		constants.SEMANTIC_ERROR,
		errorCode,
		fmt.Sprintf("%s -> %s", errorCode, token.Print()),
		token,
	)
}
