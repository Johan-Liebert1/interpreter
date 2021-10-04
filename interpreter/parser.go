package interpreter

import (
	"log"

	"interpreter/constants"
	"interpreter/helpers"
)

type Parser struct {
	Lexer        LexicalAnalyzer
	CurrentToken Token
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
func (p *Parser) Term() int {
	result := p.Factor()

	for helpers.ValueInSlice(p.CurrentToken.Type, constants.MUL_DIV_SLICE) {
		switch p.CurrentToken.Type {
		case constants.DIV:
			p.ValidateToken(constants.DIV)
			result /= p.Factor()

		case constants.MUL:
			p.ValidateToken(constants.MUL)
			result *= p.Factor()
		}
	}

	return result
}

/*
	FACTOR --> INTEGER | LPAREN EXPRESSION RPAREN
*/
func (p *Parser) Factor() int {
	token := p.CurrentToken

	var result int

	switch token.Type {
	case constants.INTEGER:
		p.ValidateToken(constants.INTEGER)
		result = token.IntegerValue

	case constants.LPAREN:
		p.ValidateToken(constants.LPAREN)
		result = p.Expression()
		p.ValidateToken(constants.RPAREN)
	}

	return result
}

/*
	Parser / Parser

	EXPRESSION --> TERM ((PLUS | MINUS) TERM)*
*/
func (p *Parser) Expression() int {
	var result int = p.Term()

	for helpers.ValueInSlice(p.CurrentToken.Type, constants.PLUS_MINUS_SLICE) {

		switch p.CurrentToken.Value {
		case constants.PLUS_SYMBOL:
			// this will advance the pointer
			p.ValidateToken(constants.PLUS)
			result += p.Term()

		case constants.MINUS_SYMBOL:
			// this will advance the pointer
			p.ValidateToken(constants.MINUS)
			result -= p.Term()
		}
	}

	return result
}

func (p *Parser) Parse() int {
	return p.Expression()
}
