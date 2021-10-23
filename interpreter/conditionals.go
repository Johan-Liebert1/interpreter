package interpreter

import "programminglang/types"

type ConditionalStatement struct {
	Type             string // if, elif, else
	Token            types.Token
	Conditionals     AbstractSyntaxTree     // LogicalNode
	ConditionalBlock AbstractSyntaxTree     // Program node
	Ladder           []ConditionalStatement // for if-elif-else ladder
}

func (cs ConditionalStatement) GetToken() types.Token {
	return cs.Token
}

func (cs ConditionalStatement) Scope(_ *Interpreter) {}
