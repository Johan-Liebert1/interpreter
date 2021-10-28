package interpreter

import "programminglang/types"

type ComparisonNode struct {
	Left       AbstractSyntaxTree // left hand of comparison node
	Comparator types.Token        // the Comparator ie. > , < , == , >= , <=
	Right      AbstractSyntaxTree // right hand of comparison node
}

func (cn ComparisonNode) GetToken() types.Token {
	return cn.Comparator
}

func (cn ComparisonNode) GetLeftOperandToken() types.Token {
	var token types.Token = cn.Left.GetToken()

	if b, ok := cn.Left.(BinaryOperationNode); ok {
		token = b.GetLeftOperandToken()
	}

	return token
}

func (cn ComparisonNode) GetRightOperandToken() types.Token {
	var token types.Token = cn.Right.GetToken()

	if b, ok := cn.Right.(BinaryOperationNode); ok {
		token = b.GetRightOperandToken()
	}

	return token
}

func (fn ComparisonNode) Scope(i *Interpreter) {}
