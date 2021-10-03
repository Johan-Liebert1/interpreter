package interpreter

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"interpreter/constants"
	"interpreter/helpers"
)

type Interpreter struct {
	Text         string
	Position     int
	CurrentToken Token
	CurrentChar  byte
	EndOfInput   bool
}

func (i *Interpreter) Init(text string) {
	i.Text = text
	i.Position = 0
	i.CurrentChar = text[0]
	i.EndOfInput = false
	i.CurrentToken = i.GetNextToken()
}

/*
	Advance the pointer into the string
*/
func (i *Interpreter) Advance() {
	i.Position++

	if i.Position >= len(i.Text) {
		i.EndOfInput = true
	} else {
		i.CurrentChar = i.Text[i.Position]
	}
}

func (i *Interpreter) SkipWhitespace() {
	for !i.EndOfInput && unicode.IsSpace(rune(i.CurrentChar)) {
		i.Advance()
	}
}

func (i *Interpreter) ConstructInteger() int {
	var s string = ""

	for !i.EndOfInput && unicode.IsDigit(rune(i.CurrentChar)) {
		s += string(i.CurrentChar)
		i.Advance()
	}

	integer, _ := strconv.Atoi(s)

	return integer
}

/*
	The lexical analyzer / scanner / tokenizer which will convert the input string to
	tokens
*/
func (i *Interpreter) GetNextToken() Token {
	for !i.EndOfInput {
		charToString := string(i.CurrentChar)

		if unicode.IsSpace(rune(i.CurrentChar)) {
			i.SkipWhitespace()
			continue
		}

		if unicode.IsDigit(rune(i.CurrentChar)) {
			integer := i.ConstructInteger()

			return Token{
				Type:         constants.INTEGER,
				IntegerValue: integer,
			}
		}

		if charToString == constants.OPERANDS[constants.PLUS] {
			i.Advance()

			return Token{
				Type:  constants.PLUS,
				Value: constants.OPERANDS[constants.PLUS],
			}
		}

		if charToString == constants.OPERANDS[constants.MINUS] {
			i.Advance()

			return Token{
				Type:  constants.MINUS,
				Value: constants.OPERANDS[constants.MINUS],
			}
		}

		if charToString == constants.OPERANDS[constants.MUL] {
			i.Advance()

			return Token{
				Type:  constants.MUL,
				Value: constants.OPERANDS[constants.MUL],
			}
		}

		if charToString == constants.OPERANDS[constants.DIV] {
			i.Advance()

			return Token{
				Type:  constants.DIV,
				Value: constants.OPERANDS[constants.DIV],
			}
		}

		if charToString == constants.LPAREN_SYMBOL {
			i.Advance()

			return Token{
				Type:  constants.LPAREN,
				Value: constants.LPAREN_SYMBOL,
			}
		}

		if charToString == constants.RPAREN_SYMBOL {
			i.Advance()

			return Token{
				Type:  constants.RPAREN,
				Value: constants.RPAREN_SYMBOL,
			}
		}

		return Token{
			Type:  constants.INVALID,
			Value: charToString,
		}

	}

	return Token{
		Type: constants.EOF,
	}
}

/*
	Validate whether the current token maches the token type passed in.

	If valid advances the parser pointer.

	If not valid, prints a fatal error and exits
*/
func (i *Interpreter) ValidateToken(tokenType string) {
	if i.CurrentToken.Type == tokenType {
		i.CurrentToken = i.GetNextToken()
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
		fmt.Println("result of i.Expression() in Factor = ", result)
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
