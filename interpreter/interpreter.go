package interpreter

import (
	"interpreter/constants"
	"strconv"
	"unicode"
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

	// asciiZero := constants.NUMBER_ASCII_RANGE[0]
	// asciiNine := constants.NUMBER_ASCII_RANGE[1]

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
