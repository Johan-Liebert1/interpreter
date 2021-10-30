package interpreter

import (
	"programminglang/types"
)

type AbstractSyntaxTree interface {
	GetToken() types.Token
	Scope(i *Interpreter)
}
