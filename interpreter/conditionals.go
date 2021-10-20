package interpreter

import "programminglang/types"

type ConditionalStatement struct {
	Type         string // if, elif, else
	Token        types.Token
	Conditionals AbstractSyntaxTree //
}
