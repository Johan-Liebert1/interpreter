package interpreter

import (
	"fmt"
	"log"
	"reflect"

	"interpreter/constants"
	"interpreter/interpreter/ast"
)

type Interpreter struct {
	TextParser  Parser
	GlobalScope map[string]int
}

func (i *Interpreter) Init(text string) {
	i.TextParser = Parser{}

	i.GlobalScope = map[string]int{}
	i.TextParser.Init(text)
}

func (i *Interpreter) Visit(node ast.AbstractSyntaxTree, depth int) int {
	// fmt.Println("Visiting node ", node, "Left : ", node.LeftOperand(), "Right : ", node.RightOperand())

	if reflect.TypeOf(node) == reflect.TypeOf(ast.Number{}) {
		// node is a Number struct, which is the base case
		// fmt.Println("found number")
		return node.Op().IntegerValue
	}

	if reflect.TypeOf(node) == reflect.TypeOf(ast.UnaryOperationNode{}) {
		// fmt.Println("found UnaryOperationNode")

		if node.Op().Type == constants.PLUS {
			return +i.Visit(node.LeftOperand(), depth+1)

		} else if node.Op().Type == constants.MINUS {
			return -i.Visit(node.LeftOperand(), depth+1)

		}
	}

	if reflect.TypeOf(node) == reflect.TypeOf(ast.CompoundStatement{}) {
		if c, ok := node.(ast.CompoundStatementNode); ok {
			for _, child := range c.GetChildren() {
				i.Visit(child, depth+1)
			}
		}
	}

	if reflect.TypeOf(node) == reflect.TypeOf(ast.AssignmentStatement{}) {
		variableName := node.LeftOperand().Op().Value

		fmt.Println(
			"Found an assignment_statement, variableName = ", variableName,
			"node.RightOperand = ", node.RightOperand(),
		)

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

	// fmt.Println("found BinaryOperationNode")

	if node.Op().Type == constants.PLUS {

		return i.Visit(node.LeftOperand(), depth+1) + i.Visit(node.RightOperand(), depth+1)

	} else if node.Op().Type == constants.MINUS {

		return i.Visit(node.LeftOperand(), depth+1) - i.Visit(node.RightOperand(), depth+1)

	} else if node.Op().Type == constants.MUL {

		return i.Visit(node.LeftOperand(), depth+1) * i.Visit(node.RightOperand(), depth+1)

	} else {

		return i.Visit(node.LeftOperand(), depth+1) / i.Visit(node.RightOperand(), depth+1)
	}
}

func (i *Interpreter) Interpret() int {
	tree := i.TextParser.Parse()

	// fmt.Println(tree)

	return i.Visit(tree, 1)
}
