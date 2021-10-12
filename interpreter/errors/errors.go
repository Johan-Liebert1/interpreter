package errors

import (
	"fmt"
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
	fmt.Println(lxe.Message)
}

func (pe *ParseError) Print() {
	fmt.Println(pe.Message)
}

func (se *SemanticError) Print() {
	fmt.Println(se.Message)
}
