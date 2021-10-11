package interpreter

import (
	"log"

	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/interpreter/ast"
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
}

func (p *Parser) Error(tokenType string) {
	log.Fatal(
		"Bad Token",
		"\nCurrent Token: ", p.CurrentToken.Print(),
		"\nToken Type to check with ", tokenType,
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
		// fmt.Println("\n\n", p.CurrentToken, "\n\n")
	} else {
		p.Error(tokenType)
	}
}

/*
	1. Gets the current token

	2. Validates the current token as integer

	3. Returns the IntegerValue of the token

	TERM --> FACTOR ((MUL | DIV) FACTOR)*
*/
func (p *Parser) Term() ast.AbstractSyntaxTree {
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

		returningValue = ast.BinaryOperationNode{
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
func (p *Parser) Factor() ast.AbstractSyntaxTree {
	token := p.CurrentToken

	var returningValue ast.AbstractSyntaxTree

	switch token.Type {
	case constants.PLUS:
		p.ValidateToken(constants.PLUS)
		returningValue = ast.UnaryOperationNode{
			Operation: token,
			Operand:   p.Factor(),
		}

	case constants.MINUS:
		p.ValidateToken(constants.MINUS)
		returningValue = ast.UnaryOperationNode{
			Operation: token,
			Operand:   p.Factor(),
		}

	case constants.INTEGER:
		p.ValidateToken(constants.INTEGER)
		returningValue = ast.IntegerNumber{
			Token: token,
			Value: token.IntegerValue,
		}

	case constants.FLOAT:
		p.ValidateToken(constants.FLOAT)
		returningValue = ast.FloatNumber{
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
func (p *Parser) Expression() ast.AbstractSyntaxTree {
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

		result = ast.BinaryOperationNode{
			Left:      result,
			Operation: currentToken,
			Right:     p.Term(),
		}
	}

	return result
}

func (p *Parser) Program() ast.AbstractSyntaxTree {
	declarationNodes := p.Declarations()
	compoundStatementNodes := p.CompoundStatement()

	node := ast.Program{
		Declarations:      declarationNodes,
		CompoundStatement: compoundStatementNodes,
	}

	return node
}

// declarations --> LET (variable_declaration SEMI)+ | blank
func (p *Parser) Declarations() []ast.AbstractSyntaxTree {
	var declarations []ast.AbstractSyntaxTree

	// variables are defined as, let varialble_name(s) : variable_type;
	if p.CurrentToken.Type == constants.LET {
		// this is messed up. there is no type called constants.LET
		p.ValidateToken(constants.LET)

		for p.CurrentToken.Type == constants.IDENTIFIER {
			varDeclaration := p.VariableDeclaration()
			declarations = append(declarations, varDeclaration...)
			p.ValidateToken(constants.SEMI_COLON)
		}

	}

	for p.CurrentToken.Type == constants.DEFINE {
		p.ValidateToken(constants.DEFINE)
		// proc_name = self.current_token.value
		// self.eat(ID)
		// self.eat(SEMI)
		// block_node = self.block()
		// proc_decl = ProcedureDecl(proc_name, block_node)
		// declarations.append(proc_decl)
		// self.eat(SEMI)

		functionName := p.CurrentToken.Value

		p.ValidateToken(constants.IDENTIFIER)

		functionBlock := p.Program()

		function := ast.FunctionDeclaration{
			FunctionName:  functionName,
			FunctionBlock: functionBlock,
		}

		declarations = append(declarations, function)
	}

	return declarations
}

// variable_declaration --> ID (COMMA ID)* COLON var_type
func (p *Parser) VariableDeclaration() []ast.AbstractSyntaxTree {
	// current node is a variable node
	variableNodes := []ast.AbstractSyntaxTree{ast.Variable{Token: p.CurrentToken, Value: p.CurrentToken.Value}}
	p.ValidateToken(constants.IDENTIFIER)

	// variables can be separated by comma so keep iterating while there's a comma
	for p.CurrentToken.Type == constants.COMMA {
		p.ValidateToken(constants.COMMA)

		variableNodes = append(variableNodes, ast.Variable{Token: p.CurrentToken, Value: p.CurrentToken.Value})

		p.ValidateToken(constants.IDENTIFIER)
	}

	// var variableName : variableType
	// variable name and type will be separated by a colon
	p.ValidateToken(constants.COLON)

	variableType := p.VarType()

	// make a new slice to store all the variable declarations
	var variableDeclarations []ast.AbstractSyntaxTree

	for _, node := range variableNodes {
		newVarDeclr := ast.VariableDeclaration{
			VariableNode: node,
			TypeNode:     variableType,
		}

		variableDeclarations = append(variableDeclarations, newVarDeclr)
	}

	return variableDeclarations
}

// formal_parameter_list --> formal_parameters | formal_parameters SEMI_COLON formal_parameter_list
// func (p *Parser) FormalParametersList() ast.AbstractSyntaxTree {}

// formal_parameters --> ID (COMMA ID)* COLON type_spec
// func (p *Parser) FormalParameters() ast.AbstractSyntaxTree {}

// var_type --> INTEGER_TYPE | FLOAT_TYPE
func (p *Parser) VarType() ast.AbstractSyntaxTree {
	token := p.CurrentToken

	if token.Type == constants.INTEGER_TYPE {
		p.ValidateToken(constants.INTEGER_TYPE)
	} else {
		p.ValidateToken(constants.FLOAT_TYPE)
	}

	return ast.VariableType{
		Token: token,
	}

}

func (p *Parser) CompoundStatement() ast.AbstractSyntaxTree {
	nodes := p.StatementList()

	root := ast.CompoundStatement{}

	root.Children = append(root.Children, nodes...)

	return root
}

// statement_list --> statement SEMI_COLON | statement SEMI_COLON statement_list
func (p *Parser) StatementList() []ast.AbstractSyntaxTree {
	node := p.Statement()

	results := []ast.AbstractSyntaxTree{node}

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
func (p *Parser) Statement() ast.AbstractSyntaxTree {
	var node ast.AbstractSyntaxTree

	if p.CurrentToken.Type == constants.IDENTIFIER {
		node = p.AssignmentStatement()
	} else if p.CurrentToken.Type == constants.INTEGER || p.CurrentToken.Type == constants.FLOAT {
		node = p.Expression()
	} else {
		node = ast.BlankStatement{
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
func (p *Parser) AssignmentStatement() ast.AbstractSyntaxTree {
	left := p.Variable()

	token := p.CurrentToken
	p.ValidateToken(constants.ASSIGN)

	right := p.Expression()

	return ast.AssignmentStatement{
		Left:  left,
		Token: token,
		Right: right,
	}
}

/*
	variable --> ID
*/
func (p *Parser) Variable() ast.AbstractSyntaxTree {
	variable := ast.Variable{
		Token: p.CurrentToken,
		Value: p.CurrentToken.Value,
	}

	p.ValidateToken(constants.IDENTIFIER)

	return variable
}

func (p *Parser) Parse() ast.AbstractSyntaxTree {
	return p.Program()
	// return p.Expression()
}
