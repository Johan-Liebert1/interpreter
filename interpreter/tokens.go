package interpreter

import "fmt"

type Token struct {
	Type         string
	Value        string
	IntegerValue int
}

func (token Token) Print() {
	fmt.Printf("Type = %s, Value = %s", token.Type, token.Value)
}
