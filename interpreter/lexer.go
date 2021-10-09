package interpreter

import (
	"fmt"
	"strconv"
	"unicode"

	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/types"
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

// skip all the whitespaces between two tokens
func (lex *LexicalAnalyzer) SkipWhitespace() {
	for !lex.EndOfInput && unicode.IsSpace(rune(lex.CurrentChar)) {
		lex.Advance()
	}
}

// skip a comment
func (lex *LexicalAnalyzer) SkipComment() {
	for !lex.EndOfInput && string(lex.CurrentChar) != "\n" {
		lex.Advance()
	}

	lex.Advance() // for the new line character
}

func (lex *LexicalAnalyzer) ConstructInteger() string {
	var s string = ""

	for !lex.EndOfInput && unicode.IsDigit(rune(lex.CurrentChar)) {
		s += string(lex.CurrentChar)
		lex.Advance()
	}

	return s
}

func (lex *LexicalAnalyzer) ConstructNumber() types.Token {
	integerPart := lex.ConstructInteger()

	if string(lex.CurrentChar) == "." {
		// is a floating point number
		s := string(lex.CurrentChar)

		fractionalPart := lex.ConstructInteger()

		realNumber, _ := strconv.ParseFloat(integerPart+s+fractionalPart, 32)

		return types.Token{
			Type:       constants.FLOAT,
			FloatValue: float32(realNumber),
		}

	}

	integer, _ := strconv.Atoi(integerPart)

	return types.Token{
		Type:         constants.INTEGER,
		IntegerValue: integer,
	}
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

	if token, ok := constants.RESERVED[identifier]; ok {
		// is a reserved keyword
		return token
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

		if charToString == constants.COMMENT_SYMBOL {
			lex.Advance()
			lex.SkipComment()
			continue
		}

		// starts with a digit, is a number
		if unicode.IsDigit(rune(lex.CurrentChar)) {
			return lex.ConstructNumber()
		}

		// starts with a letter, is an identifier
		if unicode.IsLetter(rune(lex.CurrentChar)) {
			identifier := lex.Identifier()

			fmt.Println("Constructed Identifier = ", identifier)

			return identifier

		}

		if charToString == constants.COLON_SYMBOL {
			peekPos := lex.Peek()

			fmt.Println("peekPos = ", peekPos)

			if peekPos != -1 {
				if string(lex.Text[lex.Position]) == constants.COLON_SYMBOL &&
					string(lex.Text[peekPos]) == constants.EQUAL_SYMBOL {

					lex.Advance()
					lex.Advance()

					return types.Token{
						Type:  constants.ASSIGN,
						Value: constants.ASSIGN_SYMBOL,
					}
				}
			}

			lex.Advance()

			// just a colon
			return types.Token{
				Type:  constants.COLON,
				Value: constants.COLON_SYMBOL,
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

		// could be an integer division or a float division
		// need to peek
		if charToString == constants.OPERANDS[constants.DIV] {
			peekPos := lex.Peek()

			if peekPos != -1 {
				if string(lex.Text[peekPos]) == constants.OPERANDS[constants.DIV] {
					// integer division
					lex.Advance()
					lex.Advance()

					return types.Token{
						Type:  constants.INTEGER_DIV,
						Value: constants.INTEGER_DIV_SYMBOL,
					}
				}
			} else {
				return types.Token{
					Type: constants.EOF,
				}
			}

			// otherwise float division
			lex.Advance()

			// fmt.Println("adding a division token")

			return types.Token{
				Type:  constants.FLOAT_DIV,
				Value: constants.FLOAT_DIV_SYMBOL,
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

		if charToString == constants.COMMA_SYMBOL {
			lex.Advance()

			return types.Token{
				Type:  constants.COMMA,
				Value: constants.COMMA_SYMBOL,
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
