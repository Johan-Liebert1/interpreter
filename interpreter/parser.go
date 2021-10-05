package interpreter

import (
	"fmt"
	"log"

	"interpreter/constants"
	"interpreter/helpers"
	"interpreter/interpreter/ast"
	"interpreter/types"
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

/*
	Validate whether the current token maches the token type passed in.

	If valid advances the parser pointer.

	If not valid, prints a fatal error and exits
*/
func (p *Parser) ValidateToken(tokenType string) {
	if p.CurrentToken.Type == tokenType {
		p.CurrentToken = p.Lexer.GetNextToken()
	} else {
		log.Fatal(
			"Bad Token",
			"\nCurrent Token: ", p.CurrentToken.Print(),
			"\nToken Type to check with ", tokenType,
		)
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

		fmt.Println("current token in term is saved")

		switch p.CurrentToken.Type {
		case constants.DIV:
			p.ValidateToken(constants.DIV)

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
	FACTOR --> INTEGER | LPAREN EXPRESSION RPAREN
*/
func (p *Parser) Factor() ast.AbstractSyntaxTree {
	token := p.CurrentToken

	var returningValue ast.AbstractSyntaxTree

	switch token.Type {
	case constants.INTEGER:
		p.ValidateToken(constants.INTEGER)
		returningValue = ast.Number{
			Token: token,
			Value: token.IntegerValue,
		}

	case constants.LPAREN:
		p.ValidateToken(constants.LPAREN)
		returningValue = p.Expression()
		p.ValidateToken(constants.RPAREN)
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

func (p *Parser) Parse() ast.AbstractSyntaxTree {
	return p.Expression()
}
