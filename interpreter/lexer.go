package interpreter

import (
	"fmt"
	"strconv"
	"unicode"

	"interpreter/constants"
	"interpreter/helpers"
	"interpreter/types"
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

func (lex *LexicalAnalyzer) Peek() int {
	peekPos := lex.Position + 1

	if peekPos > len(lex.Text)-1 {
		return -1
	} else {
		return peekPos
	}
}

/*
	Handles identifiers (variables) and reserved keywords
*/
func (lex *LexicalAnalyzer) Identifier() types.Token {
	identifier := ""

	for !lex.EndOfInput && helpers.IsAlphaNum(lex.CurrentChar) {
		identifier += string(lex.CurrentChar)
		lex.Advance()
	}

	return types.Token{
		Type:  constants.IDENTIFIER,
		Value: identifier,
	}
}

/*
	The lexical analyzer / scanner / tokenizer which will convert the input string to
	tokens
*/
func (lex *LexicalAnalyzer) GetNextToken() types.Token {
	for !lex.EndOfInput {
		charToString := string(lex.CurrentChar)

		if unicode.IsSpace(rune(lex.CurrentChar)) {
			lex.SkipWhitespace()
			continue
		}

		// starts with a number, is a digit
		if unicode.IsDigit(rune(lex.CurrentChar)) {
			integer := lex.ConstructInteger()

			return types.Token{
				Type:         constants.INTEGER,
				IntegerValue: integer,
			}
		}

		// starts with a letter, is an identifier
		if unicode.IsLetter(rune(lex.CurrentChar)) {
			return lex.Identifier()
		}

		if charToString == constants.EQUAL_SYMBOL {
			peekPos := lex.Peek()

			if peekPos != -1 {

				if string(lex.Text[peekPos]) == constants.COLON_SYMBOL {
					lex.Advance()
					lex.Advance()

					return types.Token{
						Type:  constants.ASSIGN,
						Value: ":=",
					}
				} else {
					// throw an error as the syntax is wrong
					break
				}

			} else {
				return types.Token{
					Type: constants.EOF,
				}
			}
		}

		if charToString == constants.SEMI_COLON_SYMBOL {
			lex.Advance()

			return types.Token{
				Type:  constants.SEMI_COLON,
				Value: constants.SEMI_COLON_SYMBOL,
			}
		}

		if charToString == constants.DOT_SYMBOL {
			lex.Advance()

			return types.Token{
				Type:  constants.DOT,
				Value: constants.DOT_SYMBOL,
			}
		}

		if charToString == constants.OPERANDS[constants.PLUS] {
			lex.Advance()

			return types.Token{
				Type:  constants.PLUS,
				Value: constants.OPERANDS[constants.PLUS],
			}
		}

		if charToString == constants.OPERANDS[constants.MINUS] {
			lex.Advance()

			return types.Token{
				Type:  constants.MINUS,
				Value: constants.OPERANDS[constants.MINUS],
			}
		}

		if charToString == constants.OPERANDS[constants.MUL] {
			lex.Advance()

			return types.Token{
				Type:  constants.MUL,
				Value: constants.OPERANDS[constants.MUL],
			}
		}

		if charToString == constants.OPERANDS[constants.DIV] {
			lex.Advance()

			fmt.Println("adding a division token")

			return types.Token{
				Type:  constants.DIV,
				Value: constants.OPERANDS[constants.DIV],
			}
		}

		if charToString == constants.LPAREN_SYMBOL {
			lex.Advance()

			return types.Token{
				Type:  constants.LPAREN,
				Value: constants.LPAREN_SYMBOL,
			}
		}

		if charToString == constants.RPAREN_SYMBOL {
			lex.Advance()

			return types.Token{
				Type:  constants.RPAREN,
				Value: constants.RPAREN_SYMBOL,
			}
		}

		return types.Token{
			Type:  constants.INVALID,
			Value: charToString,
		}

	}

	return types.Token{
		Type: constants.EOF,
	}
}
