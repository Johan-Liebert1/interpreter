package interpreter

import "fmt"

type Token struct {
	Type         string
	Value        string
	IntegerValue int
}

func (token Token) Print() string {
	return fmt.Sprintf("Type = %s, Value = %s", token.Type, token.Value)
}
