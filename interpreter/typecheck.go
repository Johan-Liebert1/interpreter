package interpreter

import (
	"fmt"
	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/interpreter/errors"
	"programminglang/types"
)

func abstractTypeCheck(leftType, operation, rightType string, opeartionToken types.Token) {
	// check if the left operand even supports the operation.
	// example: "hello" ^ 3 is not supported
	supportedOpOnLeft, ok := constants.ALLOWED_OPERATIONS_ON_TYPES[operation][leftType]

	if !ok {
		errors.ShowError(
			constants.RUNTIME_ERROR,
			constants.TYPE_ERROR,
			fmt.Sprintf("Operand '%s' not defined for type %s", opeartionToken.Value, leftType),
			opeartionToken,
		)
	}

	_, ok = supportedOpOnLeft[rightType]

	if !ok {
		errors.ShowError(
			constants.RUNTIME_ERROR,
			constants.TYPE_ERROR,
			fmt.Sprintf("Unsupported operand types for '%s' : %s and %s", opeartionToken.Value, leftType, rightType),
			opeartionToken,
		)
	}
}

func (i *Interpreter) TypeCheckBinaryOperationNode(b BinaryOperationNode) {

	if _, ok := b.Left.(Variable); ok {
		// no checks for identifies yet
		return
	}

	leftType := b.Left.GetToken().Type
	rightType := b.Right.GetToken().Type
	operation := b.Operation.Type

	abstractTypeCheck(leftType, operation, rightType, b.Operation)
}

func (i *Interpreter) TypeCheckComparisonOperationNode(c ComparisonNode) {
	helpers.ColorPrint(constants.LightGreen, 1, 1, constants.SpewPrinter.Sdump(c))

	if _, ok := c.Left.(Variable); ok {
		// no checks for identifies yet
		return
	}

	leftType := c.Left.GetToken().Type
	rightType := c.Right.GetToken().Type
	operation := c.Comparator.Type

	abstractTypeCheck(leftType, operation, rightType, c.Comparator)
}
