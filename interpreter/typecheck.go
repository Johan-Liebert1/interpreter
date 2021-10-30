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

	if leftType == constants.IDENTIFIER {
		varNode, _ := i.CurrentScope.LookupSymbol(b.GetLeftOperandToken().Value, false)
		leftType = constants.VAR_TYPE_TO_TOKEN_TYPE[varNode.Type]
		return
	}

	if rightType == constants.IDENTIFIER {
		varNode, _ := i.CurrentScope.LookupSymbol(b.GetRightOperandToken().Value, false)
		rightType = constants.VAR_TYPE_TO_TOKEN_TYPE[varNode.Type]
		return
	}

	// helpers.ColorPrint(
	// 	constants.LightCyan, 1, 1,
	// 	"left = ", constants.SpewPrinter.Sdump(b.Left),
	// 	" right = ", constants.SpewPrinter.Sdump(b.Right),
	// 	" leftType = ", leftType,
	// 	" rightType = ", rightType,
	// )

	abstractTypeCheck(leftType, operation, rightType, b.Operation)
}

func (i *Interpreter) TypeCheckComparisonOperationNode(c ComparisonNode) {
	// helpers.ColorPrint(constants.LightGreen, 1, 1, constants.SpewPrinter.Sdump(c))

	leftType := c.GetLeftOperandToken().Type
	rightType := c.GetRightOperandToken().Type
	operation := c.Comparator.Type

	if leftType == constants.IDENTIFIER {
		varNode, _ := i.CurrentScope.LookupSymbol(c.GetLeftOperandToken().Value, false)
		leftType = constants.VAR_TYPE_TO_TOKEN_TYPE[varNode.Type]
		return
	}

	if rightType == constants.IDENTIFIER {
		varNode, _ := i.CurrentScope.LookupSymbol(c.GetRightOperandToken().Value, false)
		rightType = constants.VAR_TYPE_TO_TOKEN_TYPE[varNode.Type]
		return
	}

	abstractTypeCheck(leftType, operation, rightType, c.Comparator)
}
