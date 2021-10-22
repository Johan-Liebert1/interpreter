package interpreter

import (
	"fmt"
	"log"
	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/interpreter/callstack"
)

func (i *Interpreter) EvaluateInteger(node IntegerNumber) interface{} {
	return float32(node.Token.IntegerValue)
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

	fmt.Println("Enter program")

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

	fmt.Println("Leave program")

	i.CallStack.Pop()

	return result
}

func (i *Interpreter) EvaluateFunctionCall(f FunctionCall) interface{} {
	var result interface{}

	// helpers.ColorPrint(constants.LightGreen, 1, constants.SpewPrinter.Sdump(f))

	functionName := f.FunctionName

	topAr, _ := i.CallStack.Peek()

	actualParams := f.ActualParameters

	// print to stdout if it's a print function
	if functionName == constants.PRINT_OUTPUT {
		for index := range actualParams {
			return i.Visit(actualParams[index])
		}
	}

	ar := callstack.ActivationRecord{
		Name:         functionName,
		Type:         constants.AR_FUNCTION,
		NestingLevel: topAr.NestingLevel + 1,
	}
	ar.Init()

	/*
		1. Get a list of the function's formal parameters
		2. Get a list of the function's actual parameters (arguments)
		3. For each formal parameter, get the corresponding actual parameter and save the pair in the function's activation record by using the formal parameter’s name as a key and the actual parameter (argument), after having evaluated it, as the value
	*/

	funcSymbol, _ := i.CurrentScope.LookupSymbol(functionName, false)

	formalParams := funcSymbol.ParamSymbols

	// helpers.ColorPrint(constants.White, 1, "funcsymbol = ", constants.SpewPrinter.Sdump(funcSymbol))
	// helpers.ColorPrint(constants.Magenta, 1, "Formal Params = ", formalParams)
	// helpers.ColorPrint(constants.Cyan, 1, "Actual Params = ", actualParams)

	for index := range formalParams {
		fp := formalParams[index]
		ap := actualParams[index]

		ar.SetItem(fp.Name, i.Visit(ap))
	}

	// helpers.ColorPrint(
	// 	constants.LightMagenta, 1,
	// 	"activationRecord = ",
	// 	constants.SpewPrinter.Sdump(ar),
	// )

	i.CallStack.Push(ar)

	i.Visit(funcSymbol.FunctionBlock)

	// pop the ActivationRecord at the top of the call stack after function execution is done
	i.CallStack.Pop()

	return result
}

func (i *Interpreter) EvaluateCompoundStatement(cs CompoundStatement) interface{} {
	var result interface{}

	// fmt.Println("found CompoundStatement")
	// i.spewPrinter.Dump(node)

	for _, child := range cs.Children {
		// use Token() here
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

	if varValue == nil {
		helpers.ColorPrint(constants.Red, 1, varValue, " ", variableName, constants.SpewPrinter.Sdump(i.CallStack))
	}

	floatValue, isFloat := helpers.GetFloat(varValue)

	if exists && isFloat {
		result = floatValue
	} else {
		log.Fatal("Variable ", varValue, " not defined.")
	}

	return result
}

func (i *Interpreter) EvaluateBinaryOperationNode(b BinaryOperationNode) interface{} {
	var result interface{}

	leftResult, _ := helpers.GetFloat(i.Visit(b.Left))
	rightResult, _ := helpers.GetFloat(i.Visit(b.Right))

	if b.Operation.Type == constants.PLUS {
		// fmt.Print("adding \n")
		result = leftResult + rightResult
		// fmt.Println("addition result = ", result)

	} else if b.Operation.Type == constants.MINUS {

		result = leftResult - rightResult

	} else if b.Operation.Type == constants.MUL {

		result = leftResult * rightResult

	} else if b.Operation.Type == constants.FLOAT_DIV {

		result = leftResult / rightResult
	} else {
		// integer division
		result = int(leftResult / rightResult)
	}

	return result
}
