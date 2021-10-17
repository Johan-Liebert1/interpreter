package errors

import (
	"programminglang/constants"
	"programminglang/helpers"
	"programminglang/types"
)

type Error interface {
	Print()
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

func (lxe *LexerError) Print() {
	helpers.ColorPrint(constants.Red, 1, "LexerError: ", lxe.Message)
}

func (pe *ParseError) Print() {
	helpers.ColorPrint(constants.Red, 1, "ParseError: ", pe.Message)
}

func (se *SemanticError) Print() {
	helpers.ColorPrint(constants.Red, 1, "SemanticError: ", se.Message)
}
