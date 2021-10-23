package interpreter

import "programminglang/types"

type LogicalNode struct {
	Left            AbstractSyntaxTree
	LogicalOperator types.Token // and , or , not
	Right           AbstractSyntaxTree
}

func (cn LogicalNode) GetToken() types.Token {
	return cn.LogicalOperator
}

func (fn LogicalNode) Scope(i *Interpreter) {}
