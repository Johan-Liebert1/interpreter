package constants

import (
	"programminglang/types"

	"github.com/davecgh/go-spew/spew"
)

const (
	INTEGER     = "INTEGER"
	FLOAT       = "FLOAT"
	PLUS        = "PLUS"
	MINUS       = "MINUS"
	MUL         = "MUL"
	INTEGER_DIV = "INTEGER_DIV"
	FLOAT_DIV   = "FLOAT_DIV"
	DIV         = "DIV"
	EOF         = "EOF"
	INVALID     = "INVALID"
	LPAREN      = "LPAREN"
	RPAREN      = "RAPREN"
	LCURLY      = "LCURLY"
	RCURLY      = "RCURLY"
	IDENTIFIER  = "IDENTIFIER"
	ASSIGN      = "ASSIGN"
	SEMI_COLON  = "SEMI_COLON"
	COLON       = "COLON"
	DOT         = "DOT"
	BLANK       = "BLANK"
	COMMA       = "COMMA"
)

const (
	PLUS_SYMBOL        = "+"
	MINUS_SYMBOL       = "-"
	MUL_SYMBOL         = "*"
	INTEGER_DIV_SYMBOL = "//"
	FLOAT_DIV_SYMBOL   = "/"
	LPAREN_SYMBOL      = "("
	LCURLY_SYMBOL      = "{"
	RPAREN_SYMBOL      = ")"
	RCURLY_SYMBOL      = "}"
	EQUAL_SYMBOL       = "="
	COLON_SYMBOL       = ":"
	SEMI_COLON_SYMBOL  = ";"
	DOT_SYMBOL         = "."
	ASSIGN_SYMBOL      = ":="
	COMMENT_SYMBOL     = "#"
	COMMA_SYMBOL       = ","
)

// keywords
const (
	LET          = "let"
	INTEGER_TYPE = "int"
	FLOAT_TYPE   = "float"
	DEFINE       = "define"
)

// symbol types
const (
	BUILT_IN_TYPE = "BUILT_IN_TYPE"
	VARIABLE_TYPE = "VARIABLE_TYPE"
	FUNCTION_TYPE = "FUNCTION_TYPE"
)

// error codes
const (
	ERROR_UNEXPECTED_TOKEN = "Unexpected Token"
	ERROR_ID_NOT_FOUND     = "Identifier not found"
	ERROR_DUPLICATE_ID     = "Dupliate identifier found"
)

/*
	Maps "PLUS" to "+", "MINUS" to "-", "MUL" to "*" and "DIV" to "/"
*/
var OPERANDS = map[string]string{
	PLUS:  PLUS_SYMBOL,
	MINUS: MINUS_SYMBOL,
	MUL:   MUL_SYMBOL,
	DIV:   FLOAT_DIV_SYMBOL,
}

var RESERVED = map[string]types.Token{
	LET: {
		Type:  LET,
		Value: LET,
	},

	INTEGER_TYPE: {
		Type:  INTEGER_TYPE,
		Value: INTEGER_TYPE,
	},

	FLOAT_TYPE: {
		Type:  FLOAT_TYPE,
		Value: FLOAT_TYPE,
	},

	DEFINE: {
		Type:  DEFINE,
		Value: DEFINE,
	},
}

var PLUS_MINUS_SLICE = []string{PLUS, MINUS}
var MUL_DIV_SLICE = []string{MUL, INTEGER_DIV, FLOAT_DIV}

var NUMBER_ASCII_RANGE = [2]int{48, 57}

var SpewPrinter = spew.ConfigState{Indent: "\t"}
