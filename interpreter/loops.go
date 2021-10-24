package interpreter

import "programminglang/types"

// loops that range over some values. Ex - loop from 1 to 10 with id
type RangeLoop struct {
	IdentifierToken types.Token
	Low             AbstractSyntaxTree
	High            AbstractSyntaxTree
	Block           AbstractSyntaxTree
}

func (rl RangeLoop) GetToken() types.Token {
	return rl.IdentifierToken
}

func (rl RangeLoop) Scope(_ *Interpreter) {}
