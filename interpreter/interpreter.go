package interpreter

import (
	"fmt"
	"log"

	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/interpreter/callstack"
)

type Interpreter struct {
	TextParser         Parser
	CallStack          callstack.CallStack
	ScopedSymbolsTable *ScopedSymbolsTable
	CurrentScope       *ScopedSymbolsTable
}

func (i *Interpreter) Init(text string) {
	i.TextParser = Parser{}
	i.TextParser.Init(text)

	i.CallStack = callstack.CallStack{}

	i.InitConcrete()
}

func (i *Interpreter) InitConcrete() {
	i.ScopedSymbolsTable = &ScopedSymbolsTable{}
	i.ScopedSymbolsTable.Init()

	i.CurrentScope = &ScopedSymbolsTable{}
	i.CurrentScope.Init()
}

func (i *Interpreter) Visit(node AbstractSyntaxTree, depth int) float32 {
	// fmt.Print("\n\n")
	// i.spewPrinter.Dump("Depth = ", depth, "Node = ", node)
	// fmt.Print("\n\n")

	var result float32

	if in, ok := node.(IntegerNumber); ok {
		// node is a Number struct, which is the base case
		// fmt.Println("found number", node.Op().IntegerValue)

		// meed to return an integer here
		result = float32(in.Token.IntegerValue)

	} else if f, ok := node.(FloatNumber); ok {
		// node is a Number struct, which is the base case
		// fmt.Println("found float")
		result = f.Token.FloatValue

	} else if u, ok := node.(UnaryOperationNode); ok {
		// fmt.Println("found UnaryOperationNode")

		if u.Operation.Type == constants.PLUS {
			result = +i.Visit(node.LeftOperand(), depth+1)

		} else if u.Operation.Type == constants.MINUS {
			result = -i.Visit(node.LeftOperand(), depth+1)
		}

	} else if p, ok := node.(Program); ok {
		// fmt.Println("found program")
		// i.spewPrinter.Dump(node)

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
			i.Visit(child, depth+1)
		}

		result = i.Visit(p.CompoundStatement, depth+1)

		fmt.Println("Leave program")

		i.CallStack.Pop()

	} else if f, ok := node.(FunctionCall); ok {

		functionName := f.FunctionName

		topAr, _ := i.CallStack.Peek()

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
		actualParams := f.ActualParameters

		// helpers.ColorPrint(constants.White, 1, "funcsymbol = ", constants.SpewPrinter.Sdump(funcSymbol))
		// helpers.ColorPrint(constants.Magenta, 1, "Formal Params = ", formalParams)
		// helpers.ColorPrint(constants.Cyan, 1, "Actual Params = ", actualParams)
		helpers.ColorPrint(
			constants.White, 1,
			"current scope = ", i.CurrentScope.CurrentScopeName, "  ",
			constants.SpewPrinter.Sdump(i.CurrentScope),
		)

		for index := range formalParams {
			fp := formalParams[index]
			ap := actualParams[index]

			ar.SetItem(fp.Name, i.Visit(ap, depth+1))
		}

		i.CallStack.Push(ar)

		i.Visit(funcSymbol.FunctionBlock, depth+1)

		// pop the ActivationRecord at the top of the call stack after function execution is done
		i.CallStack.Pop()

	} else if c, ok := node.(CompoundStatementNode); ok {

		// fmt.Println("found CompoundStatement")
		// i.spewPrinter.Dump(node)

		for _, child := range c.GetChildren() {
			// fmt.Println("iterating over compoundStatement child")

			// use Token() here
			if child.Op().Type == constants.BLANK {
				continue
			}

			result = i.Visit(child, depth+1)
		}

	} else if as, ok := node.(AssignmentStatement); ok {

		variableName := as.Left.Op().Value

		// fmt.Println(
		// 	"Found an assignment_statement, variableName = ", variableName,
		// 	"node.RightOperand = ", node.RightOperand(),
		// )

		variableValue := i.Visit(as.Right, depth+1)

		activationRecord, _ := i.CallStack.Peek()

		activationRecord.SetItem(variableName, variableValue)

	} else if v, ok := node.(Variable); ok {

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

	} else if b, ok := node.(BinaryOperationNode); ok {

		// BinaryOperationNode
		if b.Operation.Type == constants.PLUS {
			// fmt.Print("adding \n")
			result = i.Visit(b.Left, depth+1) + i.Visit(b.Right, depth+1)
			// fmt.Println("addition result = ", result)

		} else if b.Operation.Type == constants.MINUS {

			result = i.Visit(b.Left, depth+1) - i.Visit(b.Right, depth+1)

		} else if b.Operation.Type == constants.MUL {

			result = i.Visit(b.Left, depth+1) * i.Visit(b.Right, depth+1)

		} else if b.Operation.Type == constants.FLOAT_DIV {

			result = i.Visit(b.Left, depth+1) / i.Visit(b.Right, depth+1)
		} else {
			// integer division
			result = i.Visit(b.Left, depth+1) / i.Visit(b.Right, depth+1)
		}

	}

	// fmt.Printf("\n\n result at Depth %d = %f \n\n", depth, result)
	return result
}

// changes the interpreter's current enclosing scope to its parent's EnclosingScope
func (i *Interpreter) ReleaseScope() {
	// helpers.ColorPrint(
	// 	constants.Green, 1,
	// 	"\n Releasing Scope ", i.CurrentScope,
	// 	"\n New Scope ", i.CurrentScope.EnclosingScope,
	// )
	i.CurrentScope = i.CurrentScope.EnclosingScope
}

func (i *Interpreter) Interpret() float32 {
	tree := i.TextParser.Parse()

	// fmt.Print(tree)
	// fmt.Printf(" type = %t", tree)

	printTree := false

	if printTree {
		helpers.ColorPrint(constants.LightGreen, 1, constants.SpewPrinter.Sdump(tree))
	}

	tree.Scope(i)

	// constants.SpewPrinter.Dump(i.CurrentScope)

	return i.Visit(tree, 1)
}
