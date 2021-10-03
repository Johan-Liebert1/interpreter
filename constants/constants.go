package constants

const (
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	MINUS   = "MINUS"
	MUL     = "MUL"
	DIV     = "DIV"
	EOF     = "EOF"
)

const (
	PLUS_SYMBOL  = "+"
	MINUS_SYMBOL = "-"
	MUL_SYMBOL   = "*"
	DIV_SYMBOL   = "/"
)

var OPERANDS = map[string]string{
	PLUS:  PLUS_SYMBOL,
	MINUS: MINUS_SYMBOL,
	MUL:   MUL_SYMBOL,
	DIV:   DIV_SYMBOL,
}

var NUMBER_ASCII_RANGE = [2]int{48, 57}
