package symbols

import (
	"programminglang/constants"
)

type SymbolType struct {
}

type Symbol struct {
	Name     string // name of the identifier / symbol
	Category string // whether the symbol is a built in type, or a variable, or a function name
	Type     string // integer, float, string, etc
}

type SymbolsTable struct {
	SymbolTable map[string]Symbol
}

type ScopedSymbolsTable struct {
	CurrentScope      string
	CurrentScopeLevel int
	EnclosingScope    string
	SymbolTable       map[string]Symbol
}

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
}

/*
	Receive a symbol struct and add to hash map with key as the symbol's name and value as the symbol
*/
func (s *ScopedSymbolsTable) DefineSymbol(symbol Symbol) {
	s.SymbolTable[symbol.Name] = symbol
}

func (s *ScopedSymbolsTable) LookupSymbol(symbolName string) (Symbol, bool) {
	value, ok := s.SymbolTable[symbolName]

	return value, ok
}
