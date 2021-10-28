package interpreter

import (
	"fmt"
	"programminglang/constants"
	"programminglang/interpreter/errors"
	"programminglang/types"
)

func abstractTypeCheck(leftType, operation, rightType string, opeartionToken types.Token) {
	// check if the left operand even supports the operation.
	// example: "hello" ^ 3 is not supported
	supportedOpOnLeft, ok := constants.ALLOWED_OPERATIONS_ON_TYPES[operation][leftType]

	if !ok {
		errors.ShowError(
			constants.TYPE_ERROR,
			constants.TYPE_ERROR,
			fmt.Sprintf("Operand '%s' not defined for type %s", opeartionToken.Value, leftType),
			opeartionToken,
		)
	}

	_, ok = supportedOpOnLeft[rightType]

	if !ok {
		errors.ShowError(
			constants.TYPE_ERROR,
			constants.TYPE_ERROR,
			fmt.Sprintf("Unsupported operand types for '%s' : %s and %s", opeartionToken.Value, leftType, rightType),
			opeartionToken,
		)
	}
}

func (i *Interpreter) TypeCheckBinaryOperationNode(b BinaryOperationNode) {
	// helpers.ColorPrint(constants.LightCyan, 1, 1, constants.SpewPrinter.Sdump(b))

	leftType := b.GetLeftOperandToken().Type
	rightType := b.GetRightOperandToken().Type
	operation := b.Operation.Type

	if leftType != constants.INTEGER && leftType != constants.FLOAT && leftType != constants.STRING {
		// no checks for identifies yet
		return
	}

	if rightType != constants.INTEGER && rightType != constants.FLOAT && rightType != constants.STRING {
		// no checks for identifies yet
		return
	}

	abstractTypeCheck(leftType, operation, rightType, b.Operation)
}

func (i *Interpreter) TypeCheckComparisonOperationNode(c ComparisonNode) {
	// helpers.ColorPrint(constants.LightGreen, 1, 1, constants.SpewPrinter.Sdump(c))

	leftType := c.GetLeftOperandToken().Type
	rightType := c.GetRightOperandToken().Type
	operation := c.Comparator.Type

	if leftType != constants.INTEGER && leftType != constants.FLOAT && leftType != constants.STRING {
		// no checks for identifies yet
		return
	}

	if rightType != constants.INTEGER && rightType != constants.FLOAT && rightType != constants.STRING {
		// no checks for identifies yet
		return
	}

	abstractTypeCheck(leftType, operation, rightType, c.Comparator)
}
