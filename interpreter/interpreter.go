package interpreter

import (
	"reflect"

	"interpreter/constants"
	"interpreter/interpreter/ast"
)

type Interpreter struct {
	TextParser Parser
}

func (i *Interpreter) Init(text string) {
	i.TextParser = Parser{}

	i.TextParser.Init(text)
}

func (i *Interpreter) Visit(node ast.AbstractSyntaxTree, depth int) int {
	// fmt.Println("Visiting node ", node, "Left : ", node.LeftOperand(), "Right : ", node.RightOperand())

	if reflect.TypeOf(node) == reflect.TypeOf(ast.Number{}) {
		// node is a Number struct, which is the base case
		return node.Op().IntegerValue
	}

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
