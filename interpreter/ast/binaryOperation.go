package ast

import "interpreter/interpreter"

type Number struct {
	Token interpreter.Token
	Value int
}

type BinaryOperationNode struct {
	Left      Number
	Operation interpreter.Token
	Right     Number
}
