package interpreter

import (
	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/interpreter/callstack"
)

type Interpreter struct {
	TextParser         Parser
	CallStack          callstack.CallStack
	ScopedSymbolsTable *ScopedSymbolsTable
	CurrentScope       *ScopedSymbolsTable
}

func (i *Interpreter) Init(text string, printToken bool) {
	i.TextParser = Parser{}
	i.TextParser.Init(text, printToken)

	i.CallStack = callstack.CallStack{}

	// i.InitConcrete()
}

func (i *Interpreter) InitConcrete() {
	i.ScopedSymbolsTable = &ScopedSymbolsTable{}
	i.ScopedSymbolsTable.Init()

	i.CurrentScope = &ScopedSymbolsTable{}
	i.CurrentScope.Init()
}

func (i *Interpreter) Visit(node AbstractSyntaxTree) interface{} {
	// also need to return an empty interface here as comparisons won't be in float32

	// fmt.Print("\n\n")
	// i.spewPrinter.Dump("Depth = ", depth, "Node = ", node)
	// fmt.Print("\n\n")
	// helpers.ColorPrint(constants.LightGreen, 1, "node ", constants.SpewPrinter.Sdump(node))

	var result interface{}

	if in, ok := node.(IntegerNumber); ok {
		// node is a Number struct, which is the base case
		// fmt.Println("found number", node.Op().IntegerValue)

		// meed to return an integer here
		result = i.EvaluateInteger(in)

	} else if f, ok := node.(FloatNumber); ok {
		// node is a Number struct, which is the base case
		// fmt.Println("found float")
		result = f.Token.FloatValue

	} else if u, ok := node.(UnaryOperationNode); ok {
		// fmt.Println("found UnaryOperationNode")

		result = i.EvaluateUnaryOperator(u)

	} else if p, ok := node.(Program); ok {
		// fmt.Println("found program")
		// i.spewPrinter.Dump(node)

		result = i.EvaluateProgram(p)

	} else if f, ok := node.(FunctionCall); ok {

		result = i.EvaluateFunctionCall(f)

	} else if c, ok := node.(CompoundStatement); ok {

		result = i.EvaluateCompoundStatement(c)

	} else if as, ok := node.(AssignmentStatement); ok {

		result = i.EvaluateAssignmentStatement(as)

	} else if v, ok := node.(Variable); ok {

		result = i.EvaluateVariable(v)

	} else if b, ok := node.(BinaryOperationNode); ok {
		// BinaryOperationNode
		result = i.EvaluateBinaryOperationNode(b)

	} else if c, ok := node.(ComparisonNode); ok {
		result = i.EvaluateComparisonNode(c)

	} else if l, ok := node.(LogicalNode); ok {

		result = i.EvaluateLogicalStatement(l)

	} else if c, ok := node.(ConditionalStatement); ok {
		result = i.EvaluateConditionalStatement(c)
	}

	// fmt.Printf("\n\n result at Depth %d = %f \n\n", depth, result)
	return result
}

// changes the interpreter's current enclosing scope to its parent's EnclosingScope
func (i *Interpreter) ReleaseScope() {
	// helpers.ColorPrint(
	// 	constants.Green, 1,
	// 	"\n Releasing Scope ", i.CurrentScope,
	// 	"\n New Scope ", i.CurrentScope.EnclosingScope,
	// )
	i.CurrentScope = i.CurrentScope.EnclosingScope
}

func (i *Interpreter) Interpret() interface{} {
	tree := i.TextParser.Parse()

	// fmt.Print(tree)
	// fmt.Printf(" type = %t", tree)

	printTree := false

	if printTree {
		helpers.ColorPrint(constants.LightGreen, 1, constants.SpewPrinter.Sdump(tree))
	}

	tree.Scope(i)

	// constants.SpewPrinter.Dump(i.CurrentScope)

	return i.Visit(tree)
}
