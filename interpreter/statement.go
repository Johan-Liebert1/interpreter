package interpreter

import (
	"fmt"
	"programminglang/constants"
	"programminglang/interpreter/errors"
	"programminglang/types"
)

type CompoundStatement struct {
	Token    types.Token
	Children []AbstractSyntaxTree
}

type AssignmentStatement struct {
	Left  AbstractSyntaxTree // variable struct
	Token types.Token
	Right AbstractSyntaxTree
}

type BlankStatement struct {
	Token types.Token
}

func (cs CompoundStatement) GetToken() types.Token {
	return cs.Token
}
func (cs CompoundStatement) Scope(i *Interpreter) {
	for _, child := range cs.Children {
		child.Scope(i)
	}
}

func (v AssignmentStatement) GetToken() types.Token {
	return v.Token
}
func (as AssignmentStatement) Scope(i *Interpreter) {
	variableName := as.Left.GetToken().Value
	_, exists := i.CurrentScope.LookupSymbol(variableName, false)

	if !exists {
		errors.ShowError(
			constants.SEMANTIC_ERROR,
			constants.ERROR_VARAIBLE_NOT_DEFINED,
			fmt.Sprintf("AssignmentStatement, %s is not defined", variableName),
			as.Left.GetToken(),
		)
	}

	as.Right.Scope(i)
}

func (bs BlankStatement) GetToken() types.Token {
	return bs.Token
}
func (bs BlankStatement) Scope(_ *Interpreter) {}
