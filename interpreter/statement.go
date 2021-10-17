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

func (cs CompoundStatement) Op() types.Token {
	return cs.Token
}
func (cs CompoundStatement) LeftOperand() AbstractSyntaxTree {
	return cs.Children[0]
}
func (cs CompoundStatement) RightOperand() AbstractSyntaxTree {
	return cs.Children[0]
}
func (cs CompoundStatement) GetChildren() []AbstractSyntaxTree {
	return cs.Children
}
func (cs CompoundStatement) Scope(i *Interpreter) {
	for _, child := range cs.Children {
		child.Scope(i)
	}
}

func (v AssignmentStatement) Op() types.Token {
	return v.Token
}
func (v AssignmentStatement) LeftOperand() AbstractSyntaxTree {
	return v.Left
}
func (v AssignmentStatement) RightOperand() AbstractSyntaxTree {
	return v.Right
}
func (as AssignmentStatement) Scope(i *Interpreter) {
	variableName := as.Left.Op().Value
	_, exists := i.CurrentScope.LookupSymbol(variableName, false)

	if !exists {
		errorMessage := fmt.Sprintf("AssignmentStatement, %s is not defined", variableName)

		semanticError := errors.SemanticError{
			ErrorCode: constants.ERROR_VARAIBLE_NOT_DEFINED,
			Token:     as.Left.Op(),
			Message:   errorMessage,
		}

		semanticError.Print()
	}

	as.Right.Scope(i)
}

func (bs BlankStatement) Op() types.Token {
	return bs.Token
}
func (bs BlankStatement) LeftOperand() AbstractSyntaxTree {
	return bs
}
func (bs BlankStatement) RightOperand() AbstractSyntaxTree {
	return bs
}
func (bs BlankStatement) Scope(_ *Interpreter) {}
