package constants

const (
	INTEGER    = "INTEGER"
	PLUS       = "PLUS"
	MINUS      = "MINUS"
	MUL        = "MUL"
	DIV        = "DIV"
	EOF        = "EOF"
	INVALID    = "INVALID"
	LPAREN     = "LPAREN"
	RPAREN     = "RAPREN"
	IDENTIFIER = "IDENTIFIER"
	ASSIGN     = "ASSIGN"
	SEMI_COLON = "SEMI_COLON"
	DOT        = "DOT"
	BLANK      = "BLANK"
)

const (
	PLUS_SYMBOL       = "+"
	MINUS_SYMBOL      = "-"
	MUL_SYMBOL        = "*"
	DIV_SYMBOL        = "/"
	LPAREN_SYMBOL     = "("
	RPAREN_SYMBOL     = ")"
	EQUAL_SYMBOL      = "="
	COLON_SYMBOL      = ":"
	SEMI_COLON_SYMBOL = ";"
	DOT_SYMBOL        = "."
	ASSIGN_SYMBOL     = ":="
)

/*
	Maps "PLUS" to "+", "MINUS" to "-", "MUL" to "*" and "DIV" to "/"
*/
var OPERANDS = map[string]string{
	PLUS:  PLUS_SYMBOL,
	MINUS: MINUS_SYMBOL,
	MUL:   MUL_SYMBOL,
	DIV:   DIV_SYMBOL,
}

var PLUS_MINUS_SLICE = []string{PLUS, MINUS}
var MUL_DIV_SLICE = []string{MUL, DIV}

var NUMBER_ASCII_RANGE = [2]int{48, 57}
