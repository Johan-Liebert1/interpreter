package interpreter

import (
	"fmt"

	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/interpreter/errors"
	"programminglang/types"
)

type Parser struct {
	Lexer        LexicalAnalyzer
	CurrentToken types.Token
	printToken   bool
}

func (p *Parser) Init(text string, printToken bool) {
	p.Lexer = LexicalAnalyzer{
		Text: text,
	}

	p.Lexer.Init()

	p.CurrentToken = p.Lexer.GetNextToken()

	p.printToken = printToken

	if p.printToken {
		helpers.ColorPrint(constants.LightCyan, 1, 1, constants.SpewPrinter.Sdump(p.CurrentToken))
	}
}

func (p *Parser) Error(errorCode string, token types.Token, tokenType string) {
	errors.ShowError(
		constants.PARSER_ERROR,
		errorCode,
		fmt.Sprintf("%s -> %s \nExpected: %s", errorCode, token.Print(), tokenType),
		token,
	)
}

/*
	Validate whether the current token maches the token type passed in.

	If valid advances the parser pointer.

	If not valid, prints a fatal error and exits
*/
func (p *Parser) ValidateToken(tokenType string) {
	if p.CurrentToken.Type == tokenType {
		p.CurrentToken = p.Lexer.GetNextToken()

		if p.printToken {
			helpers.ColorPrint(constants.LightCyan, 1, 1, constants.SpewPrinter.Sdump(p.CurrentToken))
		}

		if p.CurrentToken.Type == constants.INVALID {
			p.Error(constants.ERROR_UNEXPECTED_TOKEN, p.CurrentToken, "")
		}

	} else {
		p.Error(constants.ERROR_UNEXPECTED_TOKEN, p.CurrentToken, tokenType)
	}
}

/*
	1. Gets the current token

	2. Validates the current token as integer

	3. Returns the IntegerValue of the token

	TERM --> FACTOR ((MUL | DIV) FACTOR)*
*/
func (p *Parser) Term() AbstractSyntaxTree {
	returningValue := p.Factor()

	for helpers.ValueInSlice(p.CurrentToken.Type, constants.MUL_DIV_SLICE) {
		currentToken := p.CurrentToken

		// fmt.Println("current token in term is saved")

		switch p.CurrentToken.Type {
		case constants.INTEGER_DIV:
			p.ValidateToken(constants.INTEGER_DIV)

		case constants.FLOAT_DIV:
			p.ValidateToken(constants.FLOAT_DIV)

		case constants.MUL:
			p.ValidateToken(constants.MUL)

		case constants.MODULO:
			p.ValidateToken(constants.MODULO)

		case constants.EXPONENT:
			p.ValidateToken(constants.EXPONENT)
		}

		returningValue = BinaryOperationNode{
			Left:      returningValue,
			Operation: currentToken,
			Right:     p.Factor(),
		}
	}

	// fmt.Println("\nreturinig from p.Term = ", returningValue)

	return returningValue
}

/*
	FACTOR --> ((PLUS | MINUS) FACTOR) | INTEGER | LPAREN EXPRESSION RPAREN
*/
func (p *Parser) Factor() AbstractSyntaxTree {
	token := p.CurrentToken

	var returningValue AbstractSyntaxTree

	switch token.Type {
	case constants.PLUS:
		p.ValidateToken(constants.PLUS)
		returningValue = UnaryOperationNode{
			Operation: token,
			Operand:   p.Factor(),
		}

	case constants.MINUS:
		p.ValidateToken(constants.MINUS)
		returningValue = UnaryOperationNode{
			Operation: token,
			Operand:   p.Factor(),
		}

	case constants.INTEGER:
		p.ValidateToken(constants.INTEGER)
		returningValue = IntegerNumber{
			Token: token,
			Value: token.IntegerValue,
		}

	case constants.FLOAT:
		p.ValidateToken(constants.FLOAT)
		returningValue = FloatNumber{
			Token: token,
			Value: token.FloatValue,
		}

	case constants.STRING:
		p.ValidateToken(constants.STRING)
		returningValue = String{
			Token: token,
			Value: token.Value,
		}

	case constants.TRUE:
		p.ValidateToken(constants.TRUE)
		returningValue = Boolean{
			Token: token,
			Value: true,
		}

	case constants.FALSE:
		p.ValidateToken(constants.FALSE)
		returningValue = Boolean{
			Token: token,
			Value: false,
		}

	case constants.LPAREN:
		p.ValidateToken(constants.LPAREN)
		returningValue = p.Expression()
		p.ValidateToken(constants.RPAREN)

	default:
		if p.Lexer.PeekNextToken(1).Type == constants.LPAREN {
			returningValue = p.FunctionCallStatement()
		} else {
			returningValue = p.Variable()
		}
	}

	return returningValue
}

/*
	Parser / Parser

	EXPRESSION --> TERM ((PLUS | MINUS) TERM)*
*/
func (p *Parser) Expression() AbstractSyntaxTree {
	result := p.Term()

	// fmt.Println("\nin Expression p.Term = ", result)

	for helpers.ValueInSlice(p.CurrentToken.Type, constants.PLUS_MINUS_SLICE) {
		currentToken := p.CurrentToken

		switch p.CurrentToken.Value {
		case constants.PLUS_SYMBOL:
			// this will advance the pointer
			p.ValidateToken(constants.PLUS)

		case constants.MINUS_SYMBOL:
			// this will advance the pointer
			p.ValidateToken(constants.MINUS)
		}

		result = BinaryOperationNode{
			Left:      result,
			Operation: currentToken,
			Right:     p.Term(),
		}
	}

	return result
}

// logical_statement --> NOT* (comparator ((AND | OR) comparator)*)
func (p *Parser) LogicalStatement() AbstractSyntaxTree {
	result := p.ComparisonStatement()

	// helpers.ColorPrint(constants.LightYellow, 1, "curren token in logical statement\n", p.CurrentToken)

	for helpers.ValueInSlice(p.CurrentToken.Type, constants.LOGICAL_OPERANDS_SLICE) {

		currentToken := p.CurrentToken

		switch p.CurrentToken.Value {
		case constants.AND:
			// this will advance the pointer
			p.ValidateToken(constants.AND)

		case constants.OR:
			// this will advance the pointer
			p.ValidateToken(constants.OR)

		case constants.NOT:
			// this will advance the pointer
			p.ValidateToken(constants.NOT)
		}

		result = LogicalNode{
			Left:            result,
			LogicalOperator: currentToken,
			Right:           p.ComparisonStatement(),
		}
	}

	return result
}

// comparison --> expression comparator expression
func (p *Parser) ComparisonStatement() AbstractSyntaxTree {
	result := p.Expression()

	for helpers.ValueInSlice(p.CurrentToken.Type, constants.COMPARATORS_SLICE) {
		currentToken := p.CurrentToken

		switch p.CurrentToken.Value {
		case constants.GREATER_THAN_SYMBOL:
			// this will advance the pointer
			p.ValidateToken(constants.GREATER_THAN)

		case constants.LESS_THAN_SYMBOL:
			// this will advance the pointer
			p.ValidateToken(constants.LESS_THAN)

		case constants.GREATER_THAN_EQUAL_TO_SYMBOL:
			// this will advance the pointer
			p.ValidateToken(constants.GREATER_THAN_EQUAL_TO)

		case constants.LESS_THAN_EQUAL_TO_SYMBOL:
			// this will advance the pointer
			p.ValidateToken(constants.LESS_THAN_EQUAL_TO)

		case constants.EQUALITY_SYMBOL:
			// this will advance the pointer
			p.ValidateToken(constants.EQUALITY)

		case constants.NOT_EQUAL_TO_SYMBOL:
			p.ValidateToken(constants.NOT_EQUAL_TO)
		}

		result = ComparisonNode{
			Left:       result,
			Comparator: currentToken,
			Right:      p.Expression(),
		}
	}

	return result
}

func (p *Parser) Program() AbstractSyntaxTree {
	declarationNodes := p.Declarations()
	compoundStatementNodes := p.CompoundStatement()

	node := Program{
		Declarations:      declarationNodes,
		CompoundStatement: compoundStatementNodes,
	}

	return node
}

// declarations --> LET (variable_declaration SEMI)+ (function)* | blank
func (p *Parser) Declarations() []AbstractSyntaxTree {
	var declarations []AbstractSyntaxTree

	// variables are defined as, let varialble_name(s) : variable_type;
	if p.CurrentToken.Type == constants.LET {
		p.ValidateToken(constants.LET)

		// for p.CurrentToken.Type == constants.IDENTIFIER {
		varDeclaration := p.VariableDeclaration()
		declarations = append(declarations, varDeclaration...)
		p.ValidateToken(constants.SEMI_COLON)
		// }

	}

	// for functions
	for p.CurrentToken.Type == constants.DEFINE {
		functionDeclaration := p.FunctionDeclaration()
		declarations = append(declarations, functionDeclaration)
	}

	return declarations
}

// variable_declaration --> ID (COMMA ID)* COLON var_type
func (p *Parser) VariableDeclaration() []AbstractSyntaxTree {
	// current node is a variable node
	variableNodes := []AbstractSyntaxTree{Variable{Token: p.CurrentToken, Value: p.CurrentToken.Value}}
	p.ValidateToken(constants.IDENTIFIER)

	// variables can be separated by comma so keep iterating while there's a comma
	for p.CurrentToken.Type == constants.COMMA {
		p.ValidateToken(constants.COMMA)

		variableNodes = append(variableNodes, Variable{Token: p.CurrentToken, Value: p.CurrentToken.Value})

		p.ValidateToken(constants.IDENTIFIER)
	}

	// var variableName : variableType
	// variable name and type will be separated by a colon
	p.ValidateToken(constants.COLON)

	variableType := p.VarType()

	// make a new slice to store all the variable declarations
	var variableDeclarations []AbstractSyntaxTree

	for _, node := range variableNodes {
		newVarDeclr := VariableDeclaration{
			VariableNode: node,
			TypeNode:     variableType,
		}

		variableDeclarations = append(variableDeclarations, newVarDeclr)
	}

	return variableDeclarations
}

// function_call --> ID LPAREN (expression (COMMA expression)*)? RPAREN
func (p *Parser) FunctionCallStatement() AbstractSyntaxTree {

	token := p.CurrentToken

	funcName := p.CurrentToken.Value

	p.ValidateToken(constants.IDENTIFIER)
	p.ValidateToken(constants.LPAREN)

	var actualParameters []AbstractSyntaxTree

	if p.CurrentToken.Type != constants.RPAREN {
		// has actual arguments and isn't just function()
		node := p.LogicalStatement()
		actualParameters = append(actualParameters, node)
	}

	// could be any number of parameters delimited by a comma
	for p.CurrentToken.Type == constants.COMMA {
		p.ValidateToken(constants.COMMA)
		node := p.LogicalStatement()
		actualParameters = append(actualParameters, node)
	}

	// all arguments are parsed, now check for a right parenthesis
	p.ValidateToken(constants.RPAREN)

	functionCallNode := FunctionCall{
		FunctionName:     funcName,
		ActualParameters: actualParameters,
		Token:            token,
	}

	return functionCallNode
}

// function --> DEFINE ID LPAREN formal_parameters_list? RPAREN LCURLY block RCURLY
func (p *Parser) FunctionDeclaration() AbstractSyntaxTree {
	p.ValidateToken(constants.DEFINE)

	functionName := p.CurrentToken.Value

	p.ValidateToken(constants.IDENTIFIER)

	var parametersList []FunctionParameters
	var returnStatement AbstractSyntaxTree

	if p.CurrentToken.Type == constants.LPAREN {
		p.ValidateToken(constants.LPAREN)
		parametersList = p.FormalParametersList()
		p.ValidateToken(constants.RPAREN)
	}

	p.ValidateToken(constants.LCURLY)
	functionBlock := p.Program()

	if p.CurrentToken.Type == constants.RETURN {
		p.ValidateToken(constants.RETURN)
		returnStatement = p.LogicalStatement()
	}

	if p.CurrentToken.Type == constants.SEMI_COLON {
		p.ValidateToken(constants.SEMI_COLON)
	}

	p.ValidateToken(constants.RCURLY)

	function := FunctionDeclaration{
		FunctionName:     functionName,
		FunctionBlock:    functionBlock,
		FormalParameters: parametersList,
		ReturningValue:   returnStatement,
	}

	// helpers.ColorPrint(constants.LightGreen, 1, 1, constants.SpewPrinter.Sdump(function))

	return function
}

// formal_parameter_list --> formal_parameters | formal_parameters SEMI_COLON formal_parameter_list
func (p *Parser) FormalParametersList() []FunctionParameters {
	var paramNodes []FunctionParameters

	if p.CurrentToken.Type != constants.IDENTIFIER {
		return paramNodes
	}

	paramNodes = p.FormalParameters()

	for p.CurrentToken.Type == constants.SEMI_COLON {
		p.ValidateToken(constants.SEMI_COLON)
		paramNodes = append(paramNodes, p.FormalParameters()...)
	}

	return paramNodes
}

// formal_parameters --> ID (COMMA ID)* COLON type_spec
func (p *Parser) FormalParameters() []FunctionParameters {
	var paramNodes []FunctionParameters

	paramTokens := []types.Token{p.CurrentToken}

	p.ValidateToken(constants.IDENTIFIER)

	for p.CurrentToken.Type == constants.COMMA {
		p.ValidateToken(constants.COMMA)
		paramTokens = append(paramTokens, p.CurrentToken)
		p.ValidateToken(constants.IDENTIFIER)
	}

	p.ValidateToken(constants.COLON)

	typeNode := p.VarType()

	for _, parameterToken := range paramTokens {
		paramNodes = append(paramNodes, FunctionParameters{
			VariableNode: Variable{
				Token: parameterToken,
				Value: parameterToken.Value,
			},
			TypeNode: typeNode,
		})
	}

	return paramNodes

}

// var_type --> INTEGER_TYPE | FLOAT_TYPE
func (p *Parser) VarType() AbstractSyntaxTree {
	token := p.CurrentToken

	switch token.Type {
	case constants.INTEGER_TYPE:
		p.ValidateToken(constants.INTEGER_TYPE)
	case constants.FLOAT_TYPE:
		p.ValidateToken(constants.FLOAT_TYPE)
	case constants.STRING_TYPE:
		p.ValidateToken(constants.STRING_TYPE)
	case constants.BOOLEAN_TYPE:
		p.ValidateToken(constants.BOOLEAN_TYPE)
	}

	return VariableType{
		Token: token,
	}

}

func (p *Parser) CompoundStatement() AbstractSyntaxTree {
	nodes := p.StatementList()

	root := CompoundStatement{}

	root.Children = append(root.Children, nodes...)

	return root
}

// statement_list --> statement SEMI_COLON | statement SEMI_COLON statement_list
func (p *Parser) StatementList() []AbstractSyntaxTree {
	node := p.Statement()

	results := []AbstractSyntaxTree{node}

	for p.CurrentToken.Type == constants.SEMI_COLON {
		p.ValidateToken(constants.SEMI_COLON)
		results = append(results, p.Statement())
	}

	// if p.CurrentToken.Type == constants.IDENTIFIER {
	// 	p.Error(constants.SEMI_COLON)
	// }

	return results
}

// statement --> assignment_statement | function_call | conditional_statement | blank
func (p *Parser) Statement() AbstractSyntaxTree {
	var node AbstractSyntaxTree

	// a variable foo := 3 and a function call foo() start with the same token, IDENTIFIER
	// add some conditionals to distinguish between them

	// helpers.ColorPrint(constants.Yellow, 1, "calling statement", p.CurrentToken)

	// fmt.Print("next token", p.Lexer.PeekNextToken())

	if p.CurrentToken.Type == constants.IDENTIFIER {
		if string(p.Lexer.CurrentChar) == constants.LPAREN_SYMBOL {
			// helpers.ColorPrint(constants.Yellow, 1, 1, "gonna call function")
			// a function call
			node = p.FunctionCallStatement()
		} else if p.Lexer.PeekNextToken(1).Type == constants.ASSIGN {
			// helpers.ColorPrint(constants.Yellow, 0, 1, "calling assignment_statement")
			// variable definition
			node = p.AssignmentStatement()
		} else {
			// helpers.ColorPrint(constants.Yellow, 1, 1, "calling logical_statement")
			node = p.LogicalStatement()
		}
	} else if p.CurrentToken.Type == constants.IF {
		// helpers.ColorPrint(constants.Yellow, 1, 1, "calling conditional_statement")
		node = p.ConditionalStatement()

	} else if p.CurrentToken.Type == constants.LOOP {
		// helpers.ColorPrint(constants.Yellow, 1, 1, "calling ParseLoop")

		node = p.ParseLoop()

	} else if helpers.ValueInSlice(
		p.CurrentToken.Type,
		[]string{constants.LPAREN, constants.FLOAT, constants.INTEGER, constants.NOT},
	) {
		helpers.ColorPrint(constants.Yellow, 1, 1, "calling LogicalStatement")

		node = p.LogicalStatement()

	} else {
		// helpers.ColorPrint(constants.Yellow, 1, 1, "calling BlankStatement")

		node = BlankStatement{
			Token: types.Token{
				Type:  constants.BLANK,
				Value: "",
			},
		}
	}

	return node

}

// loop --> LOOP FROM expression TO expression WITH variable LCURLY block RCURLY
func (p *Parser) ParseLoop() AbstractSyntaxTree {
	p.ValidateToken(constants.LOOP)

	p.ValidateToken(constants.FROM)

	low := p.Expression()

	p.ValidateToken(constants.TO)

	high := p.Expression()

	p.ValidateToken(constants.USING)

	loopCounter := p.CurrentToken
	p.ValidateToken(constants.IDENTIFIER)

	p.ValidateToken(constants.LCURLY)
	loopBlock := p.Program()
	p.ValidateToken(constants.RCURLY)

	node := RangeLoop{
		IdentifierToken: loopCounter,
		Low:             low,
		High:            high,
		Block:           loopBlock,
	}

	return node
}

/*
conditional_statement --> IF logical_statement LCURLY statement_list RCURLY
(ELIF logical_statement LCURLY statement_list RCURLY)* (ELSE LCURLY statement_list RCURLY){0,1}
*/
func (p *Parser) ConditionalStatement() AbstractSyntaxTree {
	currentToken := p.CurrentToken

	p.ValidateToken(constants.IF)

	condition := p.LogicalStatement()

	p.ValidateToken(constants.LCURLY)
	ifBlock := p.Program()
	p.ValidateToken(constants.RCURLY)

	node := ConditionalStatement{
		Type:             currentToken.Type,
		Token:            currentToken,
		Conditionals:     condition,
		ConditionalBlock: ifBlock,
	}

	for p.CurrentToken.Type == constants.ELSE_IF {
		currentToken := p.CurrentToken

		p.ValidateToken(constants.ELSE_IF)

		condition := p.LogicalStatement()

		p.ValidateToken(constants.LCURLY)
		ifBlock := p.Program()
		p.ValidateToken(constants.RCURLY)

		elseIfNode := ConditionalStatement{
			Type:             currentToken.Type,
			Token:            currentToken,
			Conditionals:     condition,
			ConditionalBlock: ifBlock,
		}

		node.Ladder = append(node.Ladder, elseIfNode)
	}

	if p.CurrentToken.Type == constants.ELSE {
		currentToken := p.CurrentToken

		p.ValidateToken(constants.ELSE)

		p.ValidateToken(constants.LCURLY)
		ifBlock := p.Program()
		p.ValidateToken(constants.RCURLY)

		elseNode := ConditionalStatement{
			Type:             currentToken.Type,
			Token:            currentToken,
			ConditionalBlock: ifBlock,
		}

		node.Ladder = append(node.Ladder, elseNode)
	}

	// helpers.ColorPrint(constants.LightGreen, 1, "conditional_statement ", constants.SpewPrinter.Sdump(node))

	return node
}

/*
	assignment_statement --> variable ASSIGN expression
*/
func (p *Parser) AssignmentStatement() AbstractSyntaxTree {
	left := p.Variable()

	token := p.CurrentToken
	p.ValidateToken(constants.ASSIGN)

	var right AbstractSyntaxTree

	if p.Lexer.PeekNextToken(1).Type == constants.STRING {
		right = p.Factor()
	} else if p.Lexer.PeekNextToken(1).Type == constants.LPAREN {
		// assignment to function call
		right = p.FunctionCallStatement()
	} else {
		right = p.LogicalStatement()
	}

	// helpers.ColorPrint(
	// 	constants.LightYellow, 1, 1,
	// 	"\n\n Variable AssignmentStatement \n\n",
	// 	constants.SpewPrinter.Sdump(left), constants.SpewPrinter.Sdump(token),
	// 	constants.SpewPrinter.Sdump(right),
	// )

	return AssignmentStatement{
		Left:  left,
		Token: token,
		Right: right,
	}
}

/*
	variable --> ID
*/
func (p *Parser) Variable() AbstractSyntaxTree {
	variable := Variable{
		Token: p.CurrentToken,
		Value: p.CurrentToken.Value,
	}

	p.ValidateToken(constants.IDENTIFIER)

	return variable
}

func (p *Parser) Parse() AbstractSyntaxTree {
	return p.Program()
}
