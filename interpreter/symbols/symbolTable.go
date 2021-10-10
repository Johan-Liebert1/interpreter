package symbols

import (
	"programminglang/constants"
)

type Symbol struct {
	Name string
	Type string
}

type SymbolsTable struct {
	SymbolTable map[string]Symbol
}

func (s *SymbolsTable) Init() {
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
func (s *SymbolsTable) DefineSymbol(symbol Symbol) {
	s.SymbolTable[symbol.Name] = symbol
}

func (s *SymbolsTable) LookupSymbol(symbolName string) (Symbol, bool) {
	value, ok := s.SymbolTable[symbolName]

	return value, ok
}
