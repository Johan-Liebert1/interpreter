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
	// helpers.ColorPrint(constants.Magenta, 1, "Formal Params = ", formalParams)
	// helpers.ColorPrint(constants.Cyan, 1, "Actual Params = ", actualParams)

	for index := range formalParams {
		fp := formalParams[index]
		ap := actualParams[index]

		ar.SetItem(fp.Name, i.Visit(ap))
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

	activationRecord.SetItem(variableName, variableValue)

	return result
}

func (i *Interpreter) EvaluateVariable(v Variable) interface{} {
	var result interface{}

	// if we encounter a variable, look for it in the GlobalScope and respond accordingly
	variableName := v.Token.Value

	activationRecord, _ := i.CallStack.Peek()

	varValue, exists := activationRecord.GetItem(variableName)

	if exists {
		result = varValue
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
		ar.SetItem(iteratorName, counter)
		i.Visit(l.Block)
	}

	i.CallStack.Pop()

	return result
}

func (i *Interpreter) EvaluateBinaryOperationNode(b BinaryOperationNode) interface{} {
	var result interface{}

	leftResult, _ := helpers.GetFloat(i.Visit(b.Left))
	rightResult, _ := helpers.GetFloat(i.Visit(b.Right))

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
		result = leftResult + rightResult

	case constants.MINUS:
		result = leftResult - rightResult

	case constants.MUL:
		result = leftResult * rightResult

	case constants.EXPONENT:
		result = math.Pow(float64(leftResult), float64(rightResult))

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
	var result interface{}

	left, _ := helpers.GetFloat(i.Visit(c.Left))
	right, _ := helpers.GetFloat(i.Visit(c.Right))

	switch c.Comparator.Type {
	case constants.GREATER_THAN:
		result = left > right

	case constants.LESS_THAN:
		result = left < right

	case constants.GREATER_THAN_EQUAL_TO:
		result = left >= right

	case constants.LESS_THAN_EQUAL_TO:
		result = left <= right

	case constants.EQUALITY:
		result = left == right

	case constants.NOT_EQUAL_TO:
		result = left != right

	}

	return result
}
