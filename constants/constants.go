package constants

const (
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	MINUS   = "MINUS"
	EOF     = "EOF"
)

var OPERANDS = map[string]string{
	PLUS:  "+",
	MINUS: "-",
}

var NUMBER_ASCII_RANGE = [2]int{48, 57}
