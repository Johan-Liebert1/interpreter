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

	leftToken := b.GetLeftOperandToken()
	rightToken := b.GetRightOperandToken()

	leftType := leftToken.Type
	rightType := rightToken.Type
	operation := b.Operation.Type

	activationRecord, _ := i.CallStack.Peek()

	if leftType == constants.IDENTIFIER {
		// this has "int" and not INTEGER
		val, _ := activationRecord.GetItem(leftToken.Value)
		leftType = val[constants.AR_KEY_TYPE].(string)
		leftType = constants.VAR_TYPE_TO_TOKEN_TYPE[leftType]
	}

	if rightType == constants.IDENTIFIER {
		val, _ := activationRecord.GetItem(rightToken.Value)
		rightType = val[constants.AR_KEY_TYPE].(string)
		rightType = constants.VAR_TYPE_TO_TOKEN_TYPE[rightType]
	}

	// helpers.ColorPrint(
	// 	constants.LightCyan, 1, 1,
	// 	" leftType = ", leftType,
	// 	" rightType = ", rightType,
	// )

	abstractTypeCheck(leftType, operation, rightType, b.Operation)
}

func (i *Interpreter) TypeCheckComparisonOperationNode(c ComparisonNode) {
	// helpers.ColorPrint(constants.LightGreen, 1, 1, constants.SpewPrinter.Sdump(c))

	leftToken := c.GetLeftOperandToken()
	rightToken := c.GetRightOperandToken()

	leftType := leftToken.Type
	rightType := rightToken.Type
	operation := c.Comparator.Type

	activationRecord, _ := i.CallStack.Peek()

	if leftType == constants.IDENTIFIER {
		// this has "int" and not INTEGER
		val, _ := activationRecord.GetItem(leftToken.Value)
		leftType = val[constants.AR_KEY_TYPE].(string)
		leftType = constants.VAR_TYPE_TO_TOKEN_TYPE[leftType]
	}

	if rightType == constants.IDENTIFIER {
		val, _ := activationRecord.GetItem(rightToken.Value)
		rightType = val[constants.AR_KEY_TYPE].(string)
		rightType = constants.VAR_TYPE_TO_TOKEN_TYPE[rightType]
	}

	abstractTypeCheck(leftType, operation, rightType, c.Comparator)
}
