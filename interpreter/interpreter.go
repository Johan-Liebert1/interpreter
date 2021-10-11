package interpreter

import (
	"log"
	"reflect"

	"programminglang/constants"
	"programminglang/interpreter/symbols"
)

type Interpreter struct {
	TextParser         Parser
	GlobalScope        map[string]float32
	ScopedSymbolsTable *symbols.ScopedSymbolsTable
	CurrentScope       *symbols.ScopedSymbolsTable
}

func (i *Interpreter) Init(text string) {
	i.TextParser = Parser{}

	// i.GlobalScope = map[string]float32{}
	i.TextParser.Init(text)

}

func (i *Interpreter) InitConcrete() {
	i.ScopedSymbolsTable = &symbols.ScopedSymbolsTable{}
	i.ScopedSymbolsTable.Init()

	i.CurrentScope = &symbols.ScopedSymbolsTable{}
	i.CurrentScope.Init()
}

func (i *Interpreter) Visit(node AbstractSyntaxTree, depth int) float32 {
	// fmt.Print("\n\n")
	// i.spewPrinter.Dump("Depth = ", depth, "Node = ", node)
	// fmt.Print("\n\n")

	var result float32

	if reflect.TypeOf(node) == reflect.TypeOf(IntegerNumber{}) {
		// node is a Number struct, which is the base case
		// fmt.Println("found number", node.Op().IntegerValue)

		// meed to return an integer here
		result = float32(node.Op().IntegerValue)

	} else if reflect.TypeOf(node) == reflect.TypeOf(FloatNumber{}) {
		// node is a Number struct, which is the base case
		// fmt.Println("found float")
		result = node.Op().FloatValue

	} else if reflect.TypeOf(node) == reflect.TypeOf(UnaryOperationNode{}) {
		// fmt.Println("found UnaryOperationNode")

		if node.Op().Type == constants.PLUS {
			result = +i.Visit(node.LeftOperand(), depth+1)

		} else if node.Op().Type == constants.MINUS {
			result = -i.Visit(node.LeftOperand(), depth+1)
		}

	} else if reflect.TypeOf(node) == reflect.TypeOf(Program{}) {
		// fmt.Println("found program")
		// i.spewPrinter.Dump(node)

		if c, ok := node.(Program); ok {
			for _, child := range c.Declarations {
				i.Visit(child, depth+1)
			}

			result = i.Visit(c.CompoundStatement, depth+1)
		}

	} else if reflect.TypeOf(node) == reflect.TypeOf(CompoundStatement{}) {

		// fmt.Println("found CompoundStatement")
		// i.spewPrinter.Dump(node)

		if c, ok := node.(CompoundStatementNode); ok {
			for _, child := range c.GetChildren() {
				// fmt.Println("iterating over compoundStatement child")

				if child.Op().Type == constants.BLANK {
					continue
				}

				result = i.Visit(child, depth+1)
			}
		}

	} else if reflect.TypeOf(node) == reflect.TypeOf(AssignmentStatement{}) {

		variableName := node.LeftOperand().Op().Value

		// fmt.Println(
		// 	"Found an assignment_statement, variableName = ", variableName,
		// 	"node.RightOperand = ", node.RightOperand(),
		// )

		i.GlobalScope[variableName] = i.Visit(node.RightOperand(), depth+1)

	} else if reflect.TypeOf(node) == reflect.TypeOf(Variable{}) {

		// if we encounter a variable, look for it in the GlobalScope and respond accordingly
		variableName := node.Op().Value

		if value, ok := i.GlobalScope[variableName]; ok {
			result = value
		} else {
			log.Fatal("Variable ", value, " not defined.")
		}

	} else if reflect.TypeOf(node) == reflect.TypeOf(BinaryOperationNode{}) {

		// BinaryOperationNode
		if node.Op().Type == constants.PLUS {
			// fmt.Print("adding \n")
			result = i.Visit(node.LeftOperand(), depth+1) + i.Visit(node.RightOperand(), depth+1)
			// fmt.Println("addition result = ", result)

		} else if node.Op().Type == constants.MINUS {

			result = i.Visit(node.LeftOperand(), depth+1) - i.Visit(node.RightOperand(), depth+1)

		} else if node.Op().Type == constants.MUL {

			result = i.Visit(node.LeftOperand(), depth+1) * i.Visit(node.RightOperand(), depth+1)

		} else if node.Op().Type == constants.FLOAT_DIV {

			result = i.Visit(node.LeftOperand(), depth+1) / i.Visit(node.RightOperand(), depth+1)
		} else {
			// integer division
			result = i.Visit(node.LeftOperand(), depth+1) / i.Visit(node.RightOperand(), depth+1)
		}

	}

	// fmt.Printf("\n\n result at Depth %d = %f \n\n", depth, result)
	return result
}

// changes the interpreter's current enclosing scope to its parent's EnclosingScope
func (i *Interpreter) ReleaseScope() {
	i.CurrentScope = i.CurrentScope.EnclosingScope
}

func (i *Interpreter) Interpret() float32 {
	tree := i.TextParser.Parse()

	// fmt.Print(tree)
	// fmt.Printf(" type = %t", tree)

	// i.spewPrinter.Dump(tree)

	tree.Visit(i)

	// constants.SpewPrinter.Dump(i.SymbolsTable, &i.SymbolsTable)

	return i.Visit(tree, 1)
	// return 12.34
}
