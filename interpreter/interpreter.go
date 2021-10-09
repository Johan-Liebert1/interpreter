package interpreter

import (
	"fmt"
	"log"
	"reflect"

	"programminglang/constants"
	"programminglang/interpreter/ast"

	"github.com/davecgh/go-spew/spew"
)

type Interpreter struct {
	TextParser  Parser
	GlobalScope map[string]float32
	spewPrinter spew.ConfigState
}

func (i *Interpreter) Init(text string) {
	i.TextParser = Parser{}

	i.GlobalScope = map[string]float32{}
	i.spewPrinter = spew.ConfigState{Indent: "\t"}
	i.TextParser.Init(text)
}

func (i *Interpreter) Visit(node ast.AbstractSyntaxTree, depth int) float32 {
	// fmt.Println("Visiting node ", node, "Left : ", node.LeftOperand(), "Right : ", node.RightOperand())

	// fmt.Print("\n\n")
	// i.spewPrinter.Dump("Node = ", node)
	// fmt.Print("\n\n")

	// if depth >= 5 {
	// 	return 0.0
	// }

	if reflect.TypeOf(node) == reflect.TypeOf(ast.IntegerNumber{}) {
		// node is a Number struct, which is the base case
		fmt.Println("found number")
		return float32(node.Op().IntegerValue)
	}

	if reflect.TypeOf(node) == reflect.TypeOf(ast.FloatNumber{}) {
		// node is a Number struct, which is the base case
		fmt.Println("found float")
		// meed to return a floating point here
		return node.Op().FloatValue
	}

	if reflect.TypeOf(node) == reflect.TypeOf(ast.UnaryOperationNode{}) {
		fmt.Println("found UnaryOperationNode")

		if node.Op().Type == constants.PLUS {
			return +i.Visit(node.LeftOperand(), depth+1)

		} else if node.Op().Type == constants.MINUS {
			return -i.Visit(node.LeftOperand(), depth+1)

		}
	}

	if reflect.TypeOf(node) == reflect.TypeOf(ast.Program{}) {
		fmt.Println("found program")
		// i.spewPrinter.Dump(node)

		if c, ok := node.(ast.Program); ok {
			for _, child := range c.Declarations {
				i.Visit(child, depth+1)
			}

			i.Visit(c.CompoundStatement, depth+1)
		}
	}

	if reflect.TypeOf(node) == reflect.TypeOf(ast.CompoundStatement{}) {
		fmt.Println("found CompoundStatement")
		// i.spewPrinter.Dump(node)

		if c, ok := node.(ast.CompoundStatementNode); ok {
			for _, child := range c.GetChildren() {
				i.Visit(child, depth+1)
			}
		}
	}

	if reflect.TypeOf(node) == reflect.TypeOf(ast.AssignmentStatement{}) {
		variableName := node.LeftOperand().Op().Value

		// fmt.Println(
		// 	"Found an assignment_statement, variableName = ", variableName,
		// 	"node.RightOperand = ", node.RightOperand(),
		// )

		i.GlobalScope[variableName] = i.Visit(node.RightOperand(), depth+1)
	}

	// if we encounter a variable, look for it in the GlobalScope and respond accordingly
	if reflect.TypeOf(node) == reflect.TypeOf(ast.Variable{}) {
		variableName := node.Op().Value

		if value, ok := i.GlobalScope[variableName]; ok {
			return value
		} else {
			log.Fatal("Variable ", value, " not defined.")
		}

	}

	if node.Op().Type == constants.PLUS {

		return i.Visit(node.LeftOperand(), depth+1) + i.Visit(node.RightOperand(), depth+1)

	} else if node.Op().Type == constants.MINUS {

		return i.Visit(node.LeftOperand(), depth+1) - i.Visit(node.RightOperand(), depth+1)

	} else if node.Op().Type == constants.MUL {

		return i.Visit(node.LeftOperand(), depth+1) * i.Visit(node.RightOperand(), depth+1)

	} else if node.Op().Type == constants.FLOAT_DIV {

		return i.Visit(node.LeftOperand(), depth+1) / i.Visit(node.RightOperand(), depth+1)
	} else {
		// integer division
		return i.Visit(node.LeftOperand(), depth+1) / i.Visit(node.RightOperand(), depth+1)
	}
}

func (i *Interpreter) Interpret() float32 {
	tree := i.TextParser.Parse()

	// fmt.Print(tree)
	// fmt.Printf(" type = %t", tree)

	i.spewPrinter.Dump(tree)

	return i.Visit(tree, 1)
	// return 12.34
}
