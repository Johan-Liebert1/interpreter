package interpreter

import (
	"fmt"
	"programminglang/types"
)

type Program struct {
	Declarations      []AbstractSyntaxTree
	CompoundStatement AbstractSyntaxTree
}

func (p Program) GetToken() types.Token {
	return types.Token{}
}
func (p Program) Scope(i *Interpreter) {
	var globalScope *ScopedSymbolsTable

	// a function's innner declaration also calls this Scope function so we don't want
	// another global scope being added when calling a function
	if i.CurrentScope.CurrentScopeLevel == 0 {
		fmt.Println("Entering Global Scope")

		globalScope = &ScopedSymbolsTable{
			CurrentScopeName:  "global",
			CurrentScopeLevel: 1,
		}

		globalScope.Init()
		globalScope.EnclosingScope = globalScope // no EnclosingScope so just points to itself

		// release the scope before getting out of the current scope
		defer i.ReleaseScope()

		i.CurrentScope = globalScope
	}

	for _, decl := range p.Declarations {
		decl.Scope(i)
	}

	p.CompoundStatement.Scope(i)

	if i.CurrentScope.CurrentScopeLevel == 1 {
		fmt.Println("Exiting global scope")
	}
}
