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
}

func (p *Parser) Init(text string) {
	p.Lexer = LexicalAnalyzer{
		Text: text,
	}

	p.Lexer.Init()

	p.CurrentToken = p.Lexer.GetNextToken()
	// helpers.ColorPrint(constants.LightCyan, 1, constants.SpewPrinter.Sdump(p.CurrentToken))
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
	// fmt.Println("Validating Token ", p.CurrentToken)
	// fmt.Println("Validating against ", tokenType, "\n\n")

	if p.CurrentToken.Type == tokenType {
		p.CurrentToken = p.Lexer.GetNextToken()
		// helpers.ColorPrint(constants.LightCyan, 1, constants.SpewPrinter.Sdump(p.CurrentToken))

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

	case constants.LPAREN:
		p.ValidateToken(constants.LPAREN)
		returningValue = p.Expression()
		p.ValidateToken(constants.RPAREN)

	default:
		returningValue = p.Variable()
	}

	// fmt.Println("\nreturining from Factor = ", returningValue)

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
		node := p.Expression()
		actualParameters = append(actualParameters, node)
	}

	// could be any number of parameters delimited by a comma
	for p.CurrentToken.Type == constants.COMMA {
		p.ValidateToken(constants.COMMA)
		node := p.Expression()
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

	if p.CurrentToken.Type == constants.LPAREN {
		p.ValidateToken(constants.LPAREN)
		parametersList = p.FormalParametersList()
		p.ValidateToken(constants.RPAREN)
	}

	p.ValidateToken(constants.LCURLY)
	functionBlock := p.Program()
	p.ValidateToken(constants.RCURLY)

	function := FunctionDeclaration{
		FunctionName:     functionName,
		FunctionBlock:    functionBlock,
		FormalParameters: parametersList,
	}

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

	if token.Type == constants.INTEGER_TYPE {
		p.ValidateToken(constants.INTEGER_TYPE)
	} else {
		p.ValidateToken(constants.FLOAT_TYPE)
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

/*
	statement --> assignment_statement | blank
*/
func (p *Parser) Statement() AbstractSyntaxTree {
	var node AbstractSyntaxTree

	// a variable foo := 3 and a function call foo() start with the same token, IDENTIFIER
	// add some conditionals to distinguish between them

	// helpers.ColorPrint(constants.Yellow, 1, "calling statement", p.CurrentToken)

	if p.CurrentToken.Type == constants.IDENTIFIER {
		// helpers.ColorPrint(constants.Yellow, 1, "gonna call assignment_statement")
		if string(p.Lexer.CurrentChar) == constants.LPAREN_SYMBOL {
			// a function call
			node = p.FunctionCallStatement()
		} else {
			// helpers.ColorPrint(constants.Yellow, 1, "calling assignment_statement")

			// variable definition
			node = p.AssignmentStatement()
		}

	} else if p.CurrentToken.Type == constants.INTEGER || p.CurrentToken.Type == constants.FLOAT {
		node = p.Expression()
	} else {
		node = BlankStatement{
			Token: types.Token{
				Type:  constants.BLANK,
				Value: "",
			},
		}
	}

	return node

}

/*
	assignment_statement --> variable ASSIGN expression
*/
func (p *Parser) AssignmentStatement() AbstractSyntaxTree {
	left := p.Variable()

	token := p.CurrentToken
	p.ValidateToken(constants.ASSIGN)

	right := p.Expression()

	// helpers.ColorPrint(constants.Yellow, 1, "\n\n Variable AssignmentStatement \n\n", left, token, right)

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
	// return p.Expression()
}
