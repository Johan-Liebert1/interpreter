package interpreter

import (
	"fmt"
	"log"
	"reflect"

	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/interpreter/callstack"
	"programminglang/interpreter/symbols"
)

type Interpreter struct {
	TextParser         Parser
	CallStack          callstack.CallStack
	ScopedSymbolsTable *symbols.ScopedSymbolsTable
	CurrentScope       *symbols.ScopedSymbolsTable
}

func (i *Interpreter) Init(text string) {
	i.TextParser = Parser{}
	i.TextParser.Init(text)

	i.CallStack = callstack.CallStack{}
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

		fmt.Println("Enter program")

		ar := callstack.ActivationRecord{
			Name:         constants.AR_PROGRAM,
			Type:         constants.AR_PROGRAM,
			NestingLevel: 1,
		}

		i.CallStack.Push(ar)

		if c, ok := node.(Program); ok {
			for _, child := range c.Declarations {
				i.Visit(child, depth+1)
			}

			result = i.Visit(c.CompoundStatement, depth+1)
		}

		fmt.Println("Leave program")

		i.CallStack.Pop()

	} else if reflect.TypeOf(node) == reflect.TypeOf(FunctionCall{}) {

		if f, ok := node.(FunctionCall); ok {

			functionName := f.FunctionName

			topAr := i.CallStack.Peek()

			ar := callstack.ActivationRecord{
				Name:         functionName,
				Type:         constants.AR_FUNCTION,
				NestingLevel: topAr.NestingLevel + 1,
			}

			/*
				1. Get a list of the function's formal parameters
				2. Get a list of the function's actual parameters (arguments)
				3. For each formal parameter, get the corresponding actual parameter and save the pair in the function's activation record by using the formal parameter’s name as a key and the actual parameter (argument), after having evaluated it, as the value
			*/

			funcSymbol, _ := i.CurrentScope.LookupSymbol(functionName, false)

			formalParams := funcSymbol.ParamSymbols
			actualParams := f.ActualParameters

			for index := range formalParams {
				fp := formalParams[index]
				ap := actualParams[index]

				ar.SetItem(fp.Name, i.Visit(ap, depth+1))
			}

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

		variableValue := i.Visit(node.RightOperand(), depth+1)

		activationRecord := i.CallStack.Peek()

		activationRecord.SetItem(variableName, variableValue)

	} else if reflect.TypeOf(node) == reflect.TypeOf(Variable{}) {

		// if we encounter a variable, look for it in the GlobalScope and respond accordingly
		variableName := node.Op().Value

		activationRecord := i.CallStack.Peek()

		varValue, exists := activationRecord.GetItem(variableName)

		floatValue, isFloat := helpers.GetFloat(varValue)

		if exists && isFloat {
			result = floatValue
		} else {
			log.Fatal("Variable ", varValue, " not defined.")
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
