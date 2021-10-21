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

func (fn ComparisonNode) Scope(i *Interpreter) {}
