package constants

import (
	"programminglang/types"

	"github.com/davecgh/go-spew/spew"
)

const (
	INTEGER               = "INTEGER"
	FLOAT                 = "FLOAT"
	STRING                = "STRING"
	BOOLEAN               = "BOOLEAN"
	PLUS                  = "PLUS"
	MINUS                 = "MINUS"
	MUL                   = "MUL"
	EXPONENT              = "EXPONENT"
	INTEGER_DIV           = "INTEGER_DIV"
	FLOAT_DIV             = "FLOAT_DIV"
	MODULO                = "MODULO"
	EQUALITY              = "EQUALITY"
	NOT_EQUAL_TO          = "NOT_EQUAL_TO"
	GREATER_THAN          = "GREATER_THAN"
	LESS_THAN             = "LESS_THAN"
	GREATER_THAN_EQUAL_TO = "GREATER_THAN_EQUAL_TO"
	LESS_THAN_EQUAL_TO    = "LESS_THAN_EQUAL_TO"
	DIV                   = "DIV"
	EOF                   = "EOF"
	INVALID               = "INVALID"
	LPAREN                = "LPAREN"
	RPAREN                = "RAPREN"
	LCURLY                = "LCURLY"
	RCURLY                = "RCURLY"
	IDENTIFIER            = "IDENTIFIER"
	ASSIGN                = "ASSIGN"
	SEMI_COLON            = "SEMI_COLON"
	COLON                 = "COLON"
	DOT                   = "DOT"
	BLANK                 = "BLANK"
	COMMA                 = "COMMA"
	SINGLE_QUOTE          = "SINGLE_QUOTE"
	DOUBLE_QOUTE          = "DOUBLE_QOUTE"
)

const (
	PLUS_SYMBOL                  = "+"
	MINUS_SYMBOL                 = "-"
	MUL_SYMBOL                   = "*"
	EXPONENT_SYMBOL              = "^"
	INTEGER_DIV_SYMBOL           = "//"
	FLOAT_DIV_SYMBOL             = "/"
	MODULO_SYMBOL                = "%"
	EQUALITY_SYMBOL              = "=="
	NOT_EQUAL_TO_SYMBOL          = "!="
	GREATER_THAN_SYMBOL          = ">"
	LESS_THAN_SYMBOL             = "<"
	GREATER_THAN_EQUAL_TO_SYMBOL = ">="
	LESS_THAN_EQUAL_TO_SYMBOL    = "<="
	LPAREN_SYMBOL                = "("
	LCURLY_SYMBOL                = "{"
	RPAREN_SYMBOL                = ")"
	RCURLY_SYMBOL                = "}"
	EQUAL_SYMBOL                 = "="
	COLON_SYMBOL                 = ":"
	SEMI_COLON_SYMBOL            = ";"
	DOT_SYMBOL                   = "."
	EXCLAMATION_SYMBOL           = "!"
	ASSIGN_SYMBOL                = ":="
	COMMENT_SYMBOL               = "#"
	COMMA_SYMBOL                 = ","
	SINGLE_QUOTE_SYMBOL          = "\""
	DOUBLE_QOUTE_SYMBOL          = "'"
)

// keywords
const (
	LET          = "let"
	INTEGER_TYPE = "int"
	FLOAT_TYPE   = "float"
	STRING_TYPE  = "str"
	BOOLEAN_TYPE = "bool"
	DEFINE       = "define"
	IF           = "if"
	ELSE_IF      = "elif"
	ELSE         = "else"
	AND          = "and"
	OR           = "or"
	NOT          = "not"
	LOOP         = "loop"
	FROM         = "from"
	TO           = "to"
	USING        = "using"
	RETURN       = "return"
	TRUE         = "true"
	FALSE        = "false"
)

// symbol types
const (
	BUILT_IN_TYPE = "BUILT_IN_TYPE"
	VARIABLE_TYPE = "VARIABLE_TYPE"
	FUNCTION_TYPE = "FUNCTION_TYPE"
)

// predefined functions
const (
	PRINT_OUTPUT = "output"
)

// error codes
const (
	ERROR_UNEXPECTED_TOKEN     = "Unexpected Token"
	ERROR_ID_NOT_FOUND         = "Identifier not found"
	ERROR_DUPLICATE_ID         = "Duplicate identifier found"
	ERROR_VARAIBLE_NOT_DEFINED = "Variable not defined"
	INVALID_SYNTAX             = "Invalid Syntax"
	LOGICAL_ERROR              = "Logical Error"
)

// error types
const (
	LEXER_ERROR    = "LEXER_ERROR"
	PARSER_ERROR   = "PARSER_ERROR"
	SEMANTIC_ERROR = "SEMANTIC_ERROR"
	RUNTIME_ERROR  = "RUNTIME_ERROR"
	TYPE_ERROR     = "TYPE_ERROR"
)

// activation record keys
const (
	AR_PROGRAM  = "AR_PROGRAM"
	AR_FUNCTION = "AR_FUNCTION"
	AR_LOOP     = "AR_LOOP"
	AR_IF       = "AR_IF"
	AR_ELSE_IF  = "AR_ELSE_IF"
	AR_ELSE     = "AR_ELSE"

	AR_KEY_TYPE  = "varType"
	AR_KEY_VALUE = "varValue"
)

/*
	Maps "PLUS" to "+", "MINUS" to "-", "MUL" to "*" and "DIV" to "/"
*/
var OPERANDS = map[string]string{
	PLUS:     PLUS_SYMBOL,
	MINUS:    MINUS_SYMBOL,
	MUL:      MUL_SYMBOL,
	DIV:      FLOAT_DIV_SYMBOL,
	EXPONENT: EXPONENT_SYMBOL,
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

	STRING_TYPE: {
		Type:  STRING_TYPE,
		Value: STRING_TYPE,
	},

	BOOLEAN_TYPE: {
		Type:  BOOLEAN_TYPE,
		Value: BOOLEAN_TYPE,
	},

	DEFINE: {
		Type:  DEFINE,
		Value: DEFINE,
	},

	AND: {
		Type:  AND,
		Value: AND,
	},

	OR: {
		Type:  OR,
		Value: OR,
	},

	NOT: {
		Type:  NOT,
		Value: NOT,
	},

	IF: {
		Type:  IF,
		Value: IF,
	},

	ELSE_IF: {
		Type:  ELSE_IF,
		Value: ELSE_IF,
	},

	ELSE: {
		Type:  ELSE,
		Value: ELSE,
	},

	LOOP: {
		Type:  LOOP,
		Value: LOOP,
	},

	FROM: {
		Type:  FROM,
		Value: FROM,
	},

	TO: {
		Type:  TO,
		Value: TO,
	},

	USING: {
		Type:  USING,
		Value: USING,
	},

	RETURN: {
		Type:  RETURN,
		Value: RETURN,
	},

	TRUE: {
		Type:  TRUE,
		Value: TRUE,
	},

	FALSE: {
		Type:  FALSE,
		Value: FALSE,
	},
}

var PLUS_MINUS_SLICE = []string{PLUS, MINUS}
var MUL_DIV_SLICE = []string{MUL, INTEGER_DIV, FLOAT_DIV, MODULO, EXPONENT}
var COMPARATORS_SLICE = []string{GREATER_THAN, LESS_THAN, GREATER_THAN_EQUAL_TO, LESS_THAN_EQUAL_TO, EQUALITY, NOT_EQUAL_TO}
var LOGICAL_OPERANDS_SLICE = []string{AND, OR, NOT}
var CONDITIONAL_KEYWORDS = []string{ELSE_IF, ELSE}
var QUOTES_SLICE = []string{DOUBLE_QOUTE_SYMBOL, SINGLE_QUOTE_SYMBOL}

var SpewPrinter = spew.ConfigState{Indent: "    "}

// colors
const (
	Black   = "\u001b[30;1m"
	Red     = "\u001b[31;1m"
	Green   = "\u001b[32;1m"
	Yellow  = "\u001b[33;1m"
	Blue    = "\u001b[34;1m"
	Magenta = "\u001b[35;1m"
	Cyan    = "\u001b[36;1m"
	White   = "\u001b[37;1m"
	Reset   = "\u001b[0m"
)

const (
	LightBlack   = "\u001b[30m"
	LightRed     = "\u001b[31m"
	LightGreen   = "\u001b[32m"
	LightYellow  = "\u001b[33m"
	LightBlue    = "\u001b[34m"
	LightMagenta = "\u001b[35m"
	LightCyan    = "\u001b[36m"
	LightWhite   = "\u001b[37m"
)
