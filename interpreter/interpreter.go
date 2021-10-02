package interpreter

import (
	"fmt"
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

func isSpace(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
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
	for !i.EndOfInput && isSpace(i.CurrentChar) {
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
		fmt.Println("Bad Token")
	}
}

/*
	Parser / Interpreter

	expr -> INTEGER PLUS INTEGER
	expr -> INTEGER MINUS INTEGER

*/
func (i *Interpreter) Parse() int {
	i.CurrentToken = i.GetNextToken()

	leftOperand := i.CurrentToken

	// left token needs to be an integer
	i.ValidateToken(constants.INTEGER)

	operator := i.CurrentToken

	// only works for addition and subtraction for now
	if operator.Type == constants.PLUS {
		i.ValidateToken(constants.PLUS)
	} else {
		i.ValidateToken(constants.MINUS)
	}

	// don't need to do i.GetNextToken() as i.ValidateToken() advances the pointer
	rightOperand := i.CurrentToken

	i.ValidateToken(constants.INTEGER)

	var result int

	if operator.Type == constants.PLUS {
		result = leftOperand.IntegerValue + rightOperand.IntegerValue
	} else {
		result = leftOperand.IntegerValue - rightOperand.IntegerValue
	}

	return result
}
