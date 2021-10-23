package interpreter

import "programminglang/types"

type LogicalNode struct {
	Left            AbstractSyntaxTree // comparison node
	LogicalOperator types.Token        // and , or , not
	Right           AbstractSyntaxTree // comparison node
}

func (cn LogicalNode) GetToken() types.Token {
	return cn.LogicalOperator
}

func (fn LogicalNode) Scope(i *Interpreter) {}
