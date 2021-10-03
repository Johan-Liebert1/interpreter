package interpreter

import (
	"log"
	"strconv"
	"unicode"

	"interpreter/constants"
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
	i.CurrentToken = Token{}
	i.CurrentChar = text[0]
	i.EndOfInput = false
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

		if string(i.CurrentChar) == constants.OPERANDS[constants.PLUS] {
			i.Advance()

			return Token{
				Type:  constants.PLUS,
				Value: constants.OPERANDS[constants.PLUS],
			}
		}

		if string(i.CurrentChar) == constants.OPERANDS[constants.MINUS] {
			i.Advance()

			return Token{
				Type:  constants.MINUS,
				Value: constants.OPERANDS[constants.MINUS],
			}
		}

		if string(i.CurrentChar) == constants.OPERANDS[constants.MUL] {
			i.Advance()

			return Token{
				Type:  constants.MUL,
				Value: constants.OPERANDS[constants.MUL],
			}
		}

		if string(i.CurrentChar) == constants.OPERANDS[constants.DIV] {
			i.Advance()

			return Token{
				Type:  constants.DIV,
				Value: constants.OPERANDS[constants.DIV],
			}
		}

	}

	return Token{
		Type: constants.EOF,
	}
}

/*
	Validate whether the current token maches the token type passed in, and if valid,
	advances the parser pointer. If not valid, prints an error
*/
func (i *Interpreter) ValidateToken(tokenType string) {
	if i.CurrentToken.Type == tokenType {
		i.CurrentToken = i.GetNextToken()
	} else {
		// i.error()
		log.Fatal("Bad Token")
	}
}

/*
	1. Gets the current token
	2. Validates the current token as integer
	3. Returns the IntegerValue of the token
*/
func (i *Interpreter) Term() int {
	token := i.CurrentToken

	i.ValidateToken(constants.INTEGER)

	return token.IntegerValue
}

/*
	Parser / Interpreter

	expr -> INTEGER PLUS INTEGER
	expr -> INTEGER MINUS INTEGER

*/
func (i *Interpreter) Parse() int {
	i.CurrentToken = i.GetNextToken()

	var result int = i.Term()

	operator, ok := constants.OPERANDS[i.CurrentToken.Type]

	for ok {

		switch operator {
		case constants.PLUS_SYMBOL:
			// this will advance the pointer
			i.ValidateToken(constants.PLUS)
			result += i.Term()

		case constants.MINUS_SYMBOL:
			// this will advance the pointer
			i.ValidateToken(constants.MINUS)
			result -= i.Term()

		case constants.MUL_SYMBOL:
			// this will advance the pointer
			i.ValidateToken(constants.MUL)
			result *= i.Term()

		case constants.DIV_SYMBOL:
			// this will advance the pointer
			i.ValidateToken(constants.DIV)
			result /= i.Term()
		}

		operator, ok = constants.OPERANDS[i.CurrentToken.Type]

	}

	return result
}
