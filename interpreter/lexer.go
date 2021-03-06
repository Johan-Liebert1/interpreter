package interpreter

import (
	"fmt"
	"strconv"
	"unicode"

	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/interpreter/errors"
	"programminglang/types"
)

type LexicalAnalyzer struct {
	Text        string
	Position    int
	CurrentChar byte
	EndOfInput  bool

	// for error handling
	LineNumber int
	Column     int
}

func (lex *LexicalAnalyzer) Init() {
	lex.Position = 0
	lex.CurrentChar = lex.Text[0]
	lex.EndOfInput = false
	lex.LineNumber = 1
	lex.Column = 1

	// helpers.ColorPrint(constants.Green, 2, "lexer initialized")
}

func (lex *LexicalAnalyzer) Error() {
	message := fmt.Sprintf(
		"Lexer Error: Character %s\n\tLine: %d\n\tColumn: %d\n",
		string(lex.CurrentChar),
		lex.LineNumber,
		lex.Column,
	)

	errors.ShowError(
		constants.LEXER_ERROR,
		"",
		message,
		types.Token{},
	)

}

/*
	Advance the pointer into the string
*/
func (lex *LexicalAnalyzer) Advance() {
	if string(lex.CurrentChar) == "\n" {
		lex.LineNumber++
		lex.Column = 0
	}

	lex.Position++

	if lex.Position >= len(lex.Text) {
		lex.EndOfInput = true
	} else {
		lex.CurrentChar = lex.Text[lex.Position]
		lex.Column++
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

	// helpers.ColorPrint(constants.LightGreen, 1, "constructed integer ", s)

	return s
}

func (lex *LexicalAnalyzer) ConstructNumber() types.Token {
	integerPart := lex.ConstructInteger()

	// helpers.ColorPrint(constants.LightCyan, 1, "integerPart = ", integerPart)

	if string(lex.CurrentChar) == constants.DOT_SYMBOL {
		// is a floating point number
		s := string(lex.CurrentChar) // the dot

		lex.Advance() // start from the next digit

		fractionalPart := lex.ConstructInteger()

		realNumber, _ := strconv.ParseFloat(integerPart+s+fractionalPart, 64)
		// helpers.ColorPrint(constants.LightCyan, 1, "fraction = ", fractionalPart, " realNumber = ", realNumber)

		return types.Token{
			Type:       constants.FLOAT,
			FloatValue: float32(realNumber),
			LineNumber: lex.LineNumber,
			Column:     lex.Column,
		}

	}

	integer, _ := strconv.Atoi(integerPart)

	return types.Token{
		Type:         constants.INTEGER,
		IntegerValue: integer,
		LineNumber:   lex.LineNumber,
		Column:       lex.Column,
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

	token := lex.GetToken(constants.IDENTIFIER, identifier)

	return token
}

func (lex *LexicalAnalyzer) ConstructString(quote string) types.Token {
	str := ""

	currentChar := string(lex.Text[lex.Position])

	for !lex.EndOfInput && currentChar != quote {
		str += currentChar

		lex.Advance()
		currentChar = string(lex.Text[lex.Position])
	}

	// for the last quote
	lex.Advance()

	token := lex.GetToken(constants.STRING, str)

	return token
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

			// fmt.Println("Constructed Identifier = ", identifier)

			return identifier

		}

		if helpers.ValueInSlice(charToString, constants.QUOTES_SLICE) {
			lex.Advance()

			token := lex.ConstructString(charToString)

			return token
		}

		if charToString == constants.COLON_SYMBOL {
			peekPos := lex.Peek()

			// fmt.Println("peekPos = ", peekPos)

			if peekPos != -1 {
				if string(lex.Text[lex.Position]) == constants.COLON_SYMBOL &&
					string(lex.Text[peekPos]) == constants.EQUAL_SYMBOL {
					token := lex.GetToken(constants.ASSIGN, constants.ASSIGN_SYMBOL)
					lex.Advance()
					lex.Advance()

					return token
				}
			}

			// just a colon
			token := lex.GetToken(constants.COLON, constants.COLON_SYMBOL)
			lex.Advance()
			return token
		}

		if charToString == constants.SEMI_COLON_SYMBOL {
			token := lex.GetToken(constants.SEMI_COLON, constants.SEMI_COLON_SYMBOL)
			lex.Advance()
			return token
		}

		if charToString == constants.GREATER_THAN_SYMBOL {
			// need to peek for an equal sign
			peekPos := lex.Peek()

			if peekPos != -1 {
				if string(lex.Text[lex.Position]) == constants.GREATER_THAN_SYMBOL &&
					string(lex.Text[peekPos]) == constants.EQUAL_SYMBOL {
					token := lex.GetToken(constants.GREATER_THAN_EQUAL_TO, constants.GREATER_THAN_EQUAL_TO_SYMBOL)

					lex.Advance()
					lex.Advance()

					return token
				}
			}

			token := lex.GetToken(constants.GREATER_THAN, constants.GREATER_THAN_SYMBOL)
			lex.Advance()
			return token
		}

		if charToString == constants.LESS_THAN_SYMBOL {
			// need to peek for an equal sign

			peekPos := lex.Peek()

			if peekPos != -1 {
				if string(lex.Text[lex.Position]) == constants.LESS_THAN_SYMBOL &&
					string(lex.Text[peekPos]) == constants.EQUAL_SYMBOL {
					token := lex.GetToken(constants.LESS_THAN_EQUAL_TO, constants.LESS_THAN_EQUAL_TO_SYMBOL)

					lex.Advance()
					lex.Advance()

					return token
				}
			}

			token := lex.GetToken(constants.LESS_THAN, constants.LESS_THAN_SYMBOL)
			lex.Advance()
			return token
		}

		if charToString == constants.EQUAL_SYMBOL {
			// need to peek for an equal sign

			peekPos := lex.Peek()

			if peekPos != -1 {
				if string(lex.Text[lex.Position]) == constants.EQUAL_SYMBOL &&
					string(lex.Text[peekPos]) == constants.EQUAL_SYMBOL {
					token := lex.GetToken(constants.EQUALITY, constants.EQUALITY_SYMBOL)

					lex.Advance()
					lex.Advance()

					return token
				}
			}
		}

		if charToString == constants.EXCLAMATION_SYMBOL {
			// need to peek for an equal sign
			peekPos := lex.Peek()

			if peekPos != -1 {
				if string(lex.Text[lex.Position]) == constants.EXCLAMATION_SYMBOL &&
					string(lex.Text[peekPos]) == constants.EQUAL_SYMBOL {
					token := lex.GetToken(constants.NOT_EQUAL_TO, constants.NOT_EQUAL_TO_SYMBOL)

					lex.Advance()
					lex.Advance()

					return token
				}
			}
		}

		if charToString == constants.OPERANDS[constants.PLUS] {
			token := lex.GetToken(constants.PLUS, constants.OPERANDS[constants.PLUS])
			lex.Advance()
			return token
		}

		if charToString == constants.OPERANDS[constants.MINUS] {
			token := lex.GetToken(constants.MINUS, constants.OPERANDS[constants.MINUS])
			lex.Advance()
			return token
		}

		if charToString == constants.OPERANDS[constants.MUL] {
			token := lex.GetToken(constants.MUL, constants.OPERANDS[constants.MUL])
			lex.Advance()
			return token
		}

		if charToString == constants.OPERANDS[constants.EXPONENT] {
			token := lex.GetToken(constants.EXPONENT, constants.OPERANDS[constants.EXPONENT])
			lex.Advance()
			return token
		}

		// could be an integer division or a float division
		// need to peek
		if charToString == constants.OPERANDS[constants.DIV] {
			peekPos := lex.Peek()

			if peekPos != -1 {
				if string(lex.Text[peekPos]) == constants.OPERANDS[constants.DIV] {
					// integer division
					token := lex.GetToken(constants.INTEGER_DIV, constants.INTEGER_DIV_SYMBOL)

					lex.Advance()
					lex.Advance()

					return token
				}
			} else {
				return types.Token{
					Type: constants.EOF,
				}
			}

			// otherwise float division

			token := lex.GetToken(constants.FLOAT_DIV, constants.FLOAT_DIV_SYMBOL)
			lex.Advance()

			return token
		}

		if charToString == constants.MODULO_SYMBOL {
			token := lex.GetToken(constants.MODULO, constants.MODULO_SYMBOL)
			lex.Advance()
			return token
		}

		if charToString == constants.LPAREN_SYMBOL {
			token := lex.GetToken(constants.LPAREN, constants.LPAREN_SYMBOL)
			lex.Advance()
			return token
		}

		if charToString == constants.RPAREN_SYMBOL {
			token := lex.GetToken(constants.RPAREN, constants.RPAREN_SYMBOL)
			lex.Advance()
			return token
		}

		if charToString == constants.LCURLY_SYMBOL {
			token := lex.GetToken(constants.LCURLY, constants.LCURLY_SYMBOL)
			lex.Advance()
			return token
		}

		if charToString == constants.RCURLY_SYMBOL {
			token := lex.GetToken(constants.RCURLY, constants.RCURLY_SYMBOL)
			lex.Advance()
			return token
		}

		if charToString == constants.COMMA_SYMBOL {
			token := lex.GetToken(constants.COMMA, constants.COMMA_SYMBOL)
			lex.Advance()
			return token
		}

		return lex.GetToken(constants.INVALID, charToString)

	}

	return types.Token{
		Type: constants.EOF,
	}
}

func (lex *LexicalAnalyzer) GetToken(tokenType string, tokenValue string) types.Token {
	return types.Token{
		Type:       tokenType,
		Value:      tokenValue,
		LineNumber: lex.LineNumber,
		Column:     lex.Column,
	}
}

func (lex LexicalAnalyzer) PeekNextToken(num int) types.Token {
	var token types.Token

	for i := 0; i < num; i++ {
		token = lex.GetNextToken()
	}

	return token
}
