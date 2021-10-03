package interpreter

import (
	"interpreter/constants"
	"strconv"
	"unicode"
)

type LexicalAnalyzer struct {
	Text        string
	Position    int
	CurrentChar byte
	EndOfInput  bool
}

func (lex *LexicalAnalyzer) Init() {
	lex.Position = 0
	lex.CurrentChar = lex.Text[0]
	lex.EndOfInput = false
}

/*
	Advance the pointer into the string
*/
func (lex *LexicalAnalyzer) Advance() {
	lex.Position++

	if lex.Position >= len(lex.Text) {
		lex.EndOfInput = true
	} else {
		lex.CurrentChar = lex.Text[lex.Position]
	}
}

func (lex *LexicalAnalyzer) SkipWhitespace() {
	for !lex.EndOfInput && unicode.IsSpace(rune(lex.CurrentChar)) {
		lex.Advance()
	}
}

func (lex *LexicalAnalyzer) ConstructInteger() int {
	var s string = ""

	for !lex.EndOfInput && unicode.IsDigit(rune(lex.CurrentChar)) {
		s += string(lex.CurrentChar)
		lex.Advance()
	}

	integer, _ := strconv.Atoi(s)

	return integer
}

/*
	The lexical analyzer / scanner / tokenizer which will convert the input string to
	tokens
*/
func (lex *LexicalAnalyzer) GetNextToken() Token {
	for !lex.EndOfInput {
		charToString := string(lex.CurrentChar)

		if unicode.IsSpace(rune(lex.CurrentChar)) {
			lex.SkipWhitespace()
			continue
		}

		if unicode.IsDigit(rune(lex.CurrentChar)) {
			integer := lex.ConstructInteger()

			return Token{
				Type:         constants.INTEGER,
				IntegerValue: integer,
			}
		}

		if charToString == constants.OPERANDS[constants.PLUS] {
			lex.Advance()

			return Token{
				Type:  constants.PLUS,
				Value: constants.OPERANDS[constants.PLUS],
			}
		}

		if charToString == constants.OPERANDS[constants.MINUS] {
			lex.Advance()

			return Token{
				Type:  constants.MINUS,
				Value: constants.OPERANDS[constants.MINUS],
			}
		}

		if charToString == constants.OPERANDS[constants.MUL] {
			lex.Advance()

			return Token{
				Type:  constants.MUL,
				Value: constants.OPERANDS[constants.MUL],
			}
		}

		if charToString == constants.OPERANDS[constants.DIV] {
			lex.Advance()

			return Token{
				Type:  constants.DIV,
				Value: constants.OPERANDS[constants.DIV],
			}
		}

		if charToString == constants.LPAREN_SYMBOL {
			lex.Advance()

			return Token{
				Type:  constants.LPAREN,
				Value: constants.LPAREN_SYMBOL,
			}
		}

		if charToString == constants.RPAREN_SYMBOL {
			lex.Advance()

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
