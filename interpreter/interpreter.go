package interpreter

import (
	"log"

	"interpreter/constants"
	"interpreter/helpers"
)

type Interpreter struct {
	Lexer        LexicalAnalyzer
	CurrentToken Token
}

func (i *Interpreter) Init(text string) {
	i.Lexer = LexicalAnalyzer{
		Text: text,
	}

	i.Lexer.Init()

	i.CurrentToken = i.Lexer.GetNextToken()
}

/*
	Validate whether the current token maches the token type passed in.

	If valid advances the parser pointer.

	If not valid, prints a fatal error and exits
*/
func (i *Interpreter) ValidateToken(tokenType string) {
	if i.CurrentToken.Type == tokenType {
		i.CurrentToken = i.Lexer.GetNextToken()
	} else {
		log.Fatal(
			"Bad Token",
			"\nCurrent Token: ", i.CurrentToken.Print(),
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
func (i *Interpreter) Term() int {
	result := i.Factor()

	for helpers.ValueInSlice(i.CurrentToken.Type, constants.MUL_DIV_SLICE) {
		switch i.CurrentToken.Type {
		case constants.DIV:
			i.ValidateToken(constants.DIV)
			result /= i.Factor()

		case constants.MUL:
			i.ValidateToken(constants.MUL)
			result *= i.Factor()
		}
	}

	return result
}

/*
	FACTOR --> INTEGER | LPAREN EXPRESSION RPAREN
*/
func (i *Interpreter) Factor() int {
	token := i.CurrentToken

	var result int

	switch token.Type {
	case constants.INTEGER:
		i.ValidateToken(constants.INTEGER)
		result = token.IntegerValue

	case constants.LPAREN:
		i.ValidateToken(constants.LPAREN)
		result = i.Expression()
		i.ValidateToken(constants.RPAREN)
	}

	return result
}

/*
	Parser / Interpreter

	EXPRESSION --> TERM ((PLUS | MINUS) TERM)*
*/
func (i *Interpreter) Expression() int {
	var result int = i.Term()

	for helpers.ValueInSlice(i.CurrentToken.Type, constants.PLUS_MINUS_SLICE) {

		switch i.CurrentToken.Value {
		case constants.PLUS_SYMBOL:
			// this will advance the pointer
			i.ValidateToken(constants.PLUS)
			result += i.Term()

		case constants.MINUS_SYMBOL:
			// this will advance the pointer
			i.ValidateToken(constants.MINUS)
			result -= i.Term()
		}
	}

	return result
}
