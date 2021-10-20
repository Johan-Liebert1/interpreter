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
	Left  AbstractSyntaxTree
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
		errorMessage := fmt.Sprintf("AssignmentStatement, %s is not defined", variableName)

		semanticError := errors.SemanticError{
			ErrorCode: constants.ERROR_VARAIBLE_NOT_DEFINED,
			Token:     as.Left.GetToken(),
			Message:   errorMessage,
		}

		semanticError.Print()
	}

	as.Right.Scope(i)
}

func (bs BlankStatement) GetToken() types.Token {
	return bs.Token
}
func (bs BlankStatement) Scope(_ *Interpreter) {}
