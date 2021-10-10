package symbols

type SymbolTableBuilder struct {
	SymbolTable SymbolsTable
}

func (stb *SymbolTableBuilder) Init() {
	stb.SymbolTable = SymbolsTable{}
	stb.SymbolTable.Init()
}
