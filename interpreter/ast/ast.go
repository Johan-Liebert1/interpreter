package ast

type AbstractSyntaxTree interface {
	TraverseTree(traversalType string)
}
