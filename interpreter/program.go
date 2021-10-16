package interpreter

import (
	"fmt"
	"programminglang/interpreter/symbols"
	"programminglang/types"
)

type Program struct {
	Declarations      []AbstractSyntaxTree
	CompoundStatement AbstractSyntaxTree
}

func (p Program) Op() types.Token {
	return types.Token{}
}
func (p Program) LeftOperand() AbstractSyntaxTree {
	return p
}
func (p Program) RightOperand() AbstractSyntaxTree {
	return p
}
func (p Program) Scope(i *Interpreter) {
	fmt.Println("Entering global scope")

	globalScope := &symbols.ScopedSymbolsTable{
		CurrentScopeName:  "global",
		CurrentScopeLevel: 1,
	}

	globalScope.Init()
	globalScope.EnclosingScope = globalScope // no EnclosingScope so just points to itself

	// release the scope before getting out of the current scope
	defer i.ReleaseScope()

	i.CurrentScope = globalScope

	for _, decl := range p.Declarations {
		decl.Scope(i)
	}

	p.CompoundStatement.Scope(i)

	fmt.Println("Exiting global scope")

}
