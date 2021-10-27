package errors

import (
	"fmt"
	"os"
	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/types"
)

type ErrorInterface interface {
	PrintError()
}

// Errors found in LexicalAnalyzer
type LexerError struct {
	ErrorCode string
	Token     types.Token
	Message   string
}

// Errors found while parsing
type ParseError struct {
	ErrorCode string
	Token     types.Token
	Message   string
}

// Errors found while parsing
type SemanticError struct {
	ErrorCode string
	Token     types.Token
	Message   string
}

type RuntimeError struct {
	ErrorCode string
	Token     types.Token
	Message   string
}

type TypeError struct {
	ErrorCode string
	Token     types.Token
	Message   string
}

func (lxe *LexerError) PrintError() {
	helpers.ColorPrint(constants.Red, 1, 1, fmt.Sprintf("LexerError: %s. %s", lxe.Message, lxe.Token.PrintLineCol()))
	os.Exit(1)
}

func (pe *ParseError) PrintError() {
	helpers.ColorPrint(constants.Red, 1, 1, fmt.Sprintf("ParseError: %s. %s", pe.Message, pe.Token.PrintLineCol()))
	os.Exit(1)
}

func (se *SemanticError) PrintError() {
	helpers.ColorPrint(constants.Red, 1, 1, fmt.Sprintf("SemanticError: %s. %s", se.Message, se.Token.PrintLineCol()))
	os.Exit(1)
}

func (re *RuntimeError) PrintError() {
	helpers.ColorPrint(constants.Red, 1, 1, fmt.Sprintf("RuntimeError: %s. %s", re.Message, re.Token.PrintLineCol()))
	os.Exit(1)
}

func (te *TypeError) PrintError() {
	helpers.ColorPrint(constants.Red, 1, 1, fmt.Sprintf("TypeError: %s. %s", te.Message, te.Token.PrintLineCol()))
	os.Exit(1)
}

func ShowError(errorType string, errorCode string, message string, token types.Token) {
	var e ErrorInterface

	if errorType == constants.LEXER_ERROR {
		e = &LexerError{
			ErrorCode: errorCode,
			Token:     token,
			Message:   message,
		}
	} else if errorType == constants.PARSER_ERROR {
		e = &ParseError{
			ErrorCode: errorCode,
			Token:     token,
			Message:   message,
		}
	} else if errorType == constants.SEMANTIC_ERROR {
		e = &SemanticError{
			ErrorCode: errorCode,
			Token:     token,
			Message:   message,
		}
	} else if errorType == constants.RUNTIME_ERROR {
		e = &RuntimeError{
			ErrorCode: errorCode,
			Token:     token,
			Message:   message,
		}
	} else if errorType == constants.TYPE_ERROR {
		e = &TypeError{
			ErrorCode: errorCode,
			Token:     token,
			Message:   message,
		}
	}

	e.PrintError()
}
