package interpreter

import (
	"fmt"
	"math"

	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/interpreter/callstack"
	"programminglang/interpreter/errors"
)

func (i *Interpreter) EvaluateInteger(node IntegerNumber) interface{} {
	return node.Token.IntegerValue
}

func (i *Interpreter) EvaluateUnaryOperator(node UnaryOperationNode) interface{} {
	var result interface{}

	result1, _ := helpers.GetFloat(i.Visit(node.Operand))

	if node.Operation.Type == constants.PLUS {
		result = +result1
	} else if node.Operation.Type == constants.MINUS {
		result = -result1
	}

	return result
}

func (i *Interpreter) EvaluateProgram(p Program) interface{} {
	var result interface{}

	// fmt.Println("Enter program")

	_, exists := i.CallStack.Peek()

	/* function block is also a "Program", but it's activation record will be created
	when scoping out the functional declaration so no need to do it twice */
	if !exists {
		nl := 1

		ar := callstack.ActivationRecord{
			Name:         constants.AR_PROGRAM,
			Type:         constants.AR_PROGRAM,
			NestingLevel: nl,
		}
		ar.Init()

		i.CallStack.Push(ar)
	}

	for _, child := range p.Declarations {
		i.Visit(child)
	}

	result = i.Visit(p.CompoundStatement)

	// fmt.Println("Leave program")
	if topAr, ok := i.CallStack.Peek(); ok {
		if topAr.Name == constants.AR_PROGRAM {
			i.CallStack.Pop()
		}
	}

	return result
}

func (i *Interpreter) EvaluateFunctionCall(f FunctionCall) interface{} {
	var result interface{}

	// helpers.ColorPrint(constants.LightGreen, 1, 1, constants.SpewPrinter.Sdump(i.CallStack))

	functionName := f.FunctionName

	topAr, _ := i.CallStack.Peek()

	actualParams := f.ActualParameters

	/*
		1. Get a list of the function's formal parameters
		2. Get a list of the function's actual parameters (arguments)
		3. For each formal parameter, get the corresponding actual parameter and save the pair in the function's activation record by using the formal parameter’s name as a key and the actual parameter (argument), after having evaluated it, as the value
	*/

	funcSymbol, _ := i.CurrentScope.LookupSymbol(functionName, false)
	formalParams := funcSymbol.ParamSymbols

	// helpers.ColorPrint(constants.LightCyan, 1, 1, constants.SpewPrinter.Sdump(funcSymbol))

	// print to stdout if it's a print function
	if functionName == constants.PRINT_OUTPUT {
		for index := range actualParams {
			param := actualParams[index]

			color := constants.LightYellow

			if _, ok := param.(ComparisonNode); ok {
				color = constants.LightCyan
			}

			if _, ok := param.(String); ok {
				color = constants.LightGreen
			}

			helpers.ColorPrint(color, 0, 0, i.Visit(param))
		}
		fmt.Println()
		return result
	}

	ar := callstack.ActivationRecord{
		Name:         functionName,
		Type:         constants.AR_FUNCTION,
		NestingLevel: topAr.NestingLevel + 1,
		AboveNode:    &topAr,
	}
	ar.Init()

	// helpers.ColorPrint(constants.LightCyan, 1, 1, "funcsymbol = ", constants.SpewPrinter.Sdump(funcSymbol))
	// helpers.ColorPrint(constants.Magenta, 1, 1, "CurrentScopename = ", i.CurrentScope.CurrentScopeName)
	// helpers.ColorPrint(constants.Magenta, 1, 1, "Formal Params = ", formalParams)
	// helpers.ColorPrint(constants.Cyan, 1, 1, "Actual Params = ", actualParams)

	for index := range formalParams {
		fp := formalParams[index]
		ap := actualParams[index]

		value := map[string]interface{}{
			constants.AR_KEY_TYPE:  "varType",
			constants.AR_KEY_VALUE: i.Visit(ap),
		}

		ar.SetItem(fp.Name, value, false)
	}

	// helpers.ColorPrint(
	// 	constants.LightMagenta, 1, 0,
	// 	"activationRecord = ",
	// 	constants.SpewPrinter.Sdump(ar),
	// )

	i.CallStack.Push(ar)

	i.Visit(funcSymbol.FunctionBlock)

	if funcSymbol.ReturningValue != nil {
		result = i.Visit(funcSymbol.ReturningValue)
	}

	// helpers.ColorPrint(constants.LightGreen, 1, 1, constants.SpewPrinter.Sdump(i.CallStack))
	// helpers.ColorPrint(constants.Green, 1, 1, "returning from function ", result)

	// pop the ActivationRecord at the top of the call stack after function execution is done
	i.CallStack.Pop()

	return result
}

func (i *Interpreter) EvaluateVariableDeclaration(vd VariableDeclaration) interface{} {
	var result interface{}

	variableName := vd.VariableNode.GetToken().Value

	varType := vd.TypeNode.GetToken().Value

	// helpers.ColorPrint(constants.Blue, 1, 1, varType, " ", constants.SpewPrinter.Sdump(vd))

	activationRecord, _ := i.CallStack.Peek()

	arValue := map[string]interface{}{
		constants.AR_KEY_TYPE:  varType,
		constants.AR_KEY_VALUE: nil,
	}

	activationRecord.SetItem(variableName, arValue, true)

	// helpers.ColorPrint(constants.Blue, 1, 1, varType, " ", constants.SpewPrinter.Sdump(activationRecord))

	return result
}

func (i *Interpreter) EvaluateCompoundStatement(cs CompoundStatement) interface{} {
	var result interface{}

	// fmt.Println("found CompoundStatement")
	// i.spewPrinter.Dump(node)

	for _, child := range cs.Children {
		if child.GetToken().Type == constants.BLANK {
			continue
		}

		result = i.Visit(child)
	}

	return result
}

func (i *Interpreter) EvaluateAssignmentStatement(as AssignmentStatement) interface{} {
	var result interface{}

	variableName := as.Left.GetToken().Value

	variableValue := i.Visit(as.Right)

	activationRecord, _ := i.CallStack.Peek()

	// helpers.ColorPrint(constants.Blue, 1, 1, constants.SpewPrinter.Sdump(as))

	arValue := map[string]interface{}{
		constants.AR_KEY_TYPE:  "varType",
		constants.AR_KEY_VALUE: variableValue,
	}

	activationRecord.SetItem(variableName, arValue, false)

	return result
}

func (i *Interpreter) EvaluateVariable(v Variable) interface{} {
	var result interface{}

	// if we encounter a variable, look for it in the GlobalScope and respond accordingly
	variableName := v.Token.Value

	activationRecord, _ := i.CallStack.Peek()

	varValue, exists := activationRecord.GetItem(variableName)

	if exists {
		result = varValue[constants.AR_KEY_VALUE]
	} else {
		errors.ShowError(
			constants.SEMANTIC_ERROR,
			constants.ERROR_VARAIBLE_NOT_DEFINED,
			fmt.Sprintf("Variable '%s' is not defined.", variableName),
			v.Token,
		)
	}

	return result
}

func (i *Interpreter) EvaluateLogicalStatement(l LogicalNode) interface{} {

	var result interface{}

	leftResult, lok := i.Visit(l.Left).(bool)
	rightResult, rok := i.Visit(l.Right).(bool)

	// helpers.ColorPrint(constants.Green, 2, "leftResult ", leftResult, " right result ", rightResult)

	if lok && rok {

		switch l.LogicalOperator.Type {
		case constants.AND:
			result = leftResult && rightResult

		case constants.OR:
			result = leftResult || rightResult

		}

	}

	return result
}

func (i *Interpreter) EvaluateConditionalStatement(c ConditionalStatement) interface{} {
	var result interface{}
	var (
		elseBlock ConditionalStatement
		visitElse bool = false
	)

	enterBlock, _ := i.Visit(c.Conditionals).(bool)

	if enterBlock {
		result = i.Visit(c.ConditionalBlock)

	} else {
		// didn't enter if block, so start traversing the else if ladder if there is any

		for _, statement := range c.Ladder {

			if statement.Type == constants.ELSE {
				// since we reached the else block, every other block failed
				elseBlock = statement
				visitElse = true
			}

			enterInnerBlock, _ := i.Visit(statement.Conditionals).(bool)

			if enterInnerBlock {
				result = i.Visit(statement.ConditionalBlock)
				// else if ladder, one statement was true, execute it and break the loop
				break
			}

		}
	}

	if visitElse {
		result = i.Visit(elseBlock.ConditionalBlock)
	}

	return result
}

func (i *Interpreter) EvaluateRangeLoop(l RangeLoop) interface{} {
	// helpers.ColorPrint(constants.LightYellow, 1, 1, "loop = ", constants.SpewPrinter.Sdump(l))

	var (
		low  int
		high int
	)

	visitLow := i.Visit(l.Low)
	visitHigh := i.Visit(l.High)

	if val, ok := visitLow.(float32); ok {
		low = int(val)
	} else if val, ok := visitLow.(float64); ok {
		low = int(val)
	} else {
		low = visitLow.(int)
	}

	if val, ok := visitHigh.(float32); ok {
		high = int(val)
	} else if val, ok := visitHigh.(float64); ok {
		high = int(val)
	} else {
		high = visitHigh.(int)
	}

	iteratorName := l.IdentifierToken.Value

	topAr, _ := i.CallStack.Peek()

	ar := &callstack.ActivationRecord{
		Name:         constants.AR_LOOP,
		Type:         constants.AR_LOOP,
		NestingLevel: topAr.NestingLevel + 1,
		AboveNode:    &topAr,
	}
	ar.Init()

	i.CallStack.Push(*ar)

	// helpers.ColorPrint(constants.LightGreen, 1, 1, constants.SpewPrinter.Sdump(i.CallStack))

	var result interface{}

	for counter := int(low); counter <= int(high); counter++ {
		arValue := map[string]interface{}{
			constants.AR_KEY_TYPE:  constants.INTEGER,
			constants.AR_KEY_VALUE: counter,
		}

		ar.SetItem(iteratorName, arValue, true)
		i.Visit(l.Block)
	}

	i.CallStack.Pop()

	return result
}

func (i *Interpreter) EvaluateBinaryOperationNode(b BinaryOperationNode) interface{} {
	i.TypeCheckBinaryOperationNode(b)

	var result interface{}

	leftVisit := i.Visit(b.Left)
	rightVisit := i.Visit(b.Right)

	var (
		leftResult     float32
		leftIntResult  int
		rightResult    float32
		rightIntResult int
		isLeftFloat    bool
		isLeftInt      bool
		isRightInt     bool
	)

	// 6 ^ 2 - 16

	leftResult, isLeftFloat = leftVisit.(float32)
	rightResult, _ = rightVisit.(float32)

	leftIntResult, isLeftInt = leftVisit.(int)
	rightIntResult, isRightInt = rightVisit.(int)

	if isLeftInt {
		leftResult = float32(leftIntResult)
	}

	if isRightInt {
		rightResult = float32(rightIntResult)
	}

	divideByZero := func() {

		fmt.Print(b.Right.GetToken())
		errors.ShowError(
			constants.RUNTIME_ERROR,
			constants.LOGICAL_ERROR,
			fmt.Sprintf("Cannot divide by zero. %s", b.Right.GetToken().PrintLineCol()),
			b.Right.GetToken(),
		)
	}

	switch b.Operation.Type {
	case constants.PLUS:
		{
			if isLeftFloat || isLeftInt {
				result = leftResult + rightResult
			} else {
				// TODO: left and right are string
				s := ""

				if ls, ok := leftVisit.(string); ok {
					for i := 0; i < len(ls); i++ {
						s += string(ls[i])
					}
				}

				if rs, ok := leftVisit.(string); ok {
					for i := 0; i < len(rs); i++ {
						s += string(rs[i])
					}
				}

				result = s
			}
		}

	case constants.MINUS:
		result = leftResult - rightResult

	case constants.MUL:
		{
			if isLeftFloat || isLeftInt {
				result = leftResult * rightResult
			} else {
				// TODO: left and right are string
				temp := ""

				if ls, ok := leftVisit.(string); ok {
					for i := 0; i < len(ls); i++ {
						temp += string(ls[i])
					}
				}

				s := ""

				if rightInt, ok := rightVisit.(int); ok {
					for i := 0; i < rightInt; i++ {
						s += temp
					}
				}

				result = s
			}
		}

	case constants.EXPONENT:
		result = float32(math.Pow(float64(leftResult), float64(rightResult)))

	case constants.FLOAT_DIV:
		if rightResult == 0.0 {
			divideByZero()
		}
		result = leftResult / rightResult

	case constants.INTEGER_DIV:
		if rightResult == 0.0 {
			divideByZero()
		}
		result = int(leftResult / rightResult)

	case constants.MODULO:
		result = int(leftResult) % int(rightResult)

	}

	return result
}

func (i *Interpreter) EvaluateComparisonNode(c ComparisonNode) interface{} {
	i.TypeCheckComparisonOperationNode(c)

	var result interface{}

	leftVisit := i.Visit(c.Left)
	rightVisit := i.Visit(c.Right)

	var (
		leftResult      float32
		leftIntResult   int
		rightResult     float32
		rightIntResult  int
		isLeftInt       bool
		isRightInt      bool
		isLeftStr       bool
		leftStrResult   string
		rightStrResult  string
		areStringsEqual bool = true
	)

	leftResult, _ = leftVisit.(float32)
	rightResult, _ = rightVisit.(float32)

	leftIntResult, isLeftInt = leftVisit.(int)
	rightIntResult, isRightInt = rightVisit.(int)

	leftStrResult, isLeftStr = leftVisit.(string)
	rightStrResult, _ = rightVisit.(string)

	if isLeftInt {
		leftResult = float32(leftIntResult)
	}

	if isRightInt {
		rightResult = float32(rightIntResult)
	}

	if isLeftStr {
		leftResult = float32(len(leftStrResult))
		rightResult = float32(len(rightStrResult))

		// check for string equality right here

		if leftResult != rightResult {
			areStringsEqual = false
		} else {
			for i := 0; i < len(leftStrResult); i++ {
				if leftStrResult[i] != rightStrResult[i] {
					areStringsEqual = false
					break
				}
			}

		}

	}

	switch c.Comparator.Type {
	case constants.GREATER_THAN:
		result = leftResult > rightResult

	case constants.LESS_THAN:
		result = leftResult < rightResult

	case constants.GREATER_THAN_EQUAL_TO:
		result = leftResult >= rightResult

	case constants.LESS_THAN_EQUAL_TO:
		result = leftResult <= rightResult

	case constants.EQUALITY:
		if isLeftStr {
			result = areStringsEqual
		} else {
			result = leftResult == rightResult
		}

	case constants.NOT_EQUAL_TO:
		if isLeftStr {
			result = !areStringsEqual
		} else {
			result = leftResult != rightResult
		}

	}

	return result
}
