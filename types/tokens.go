package types

import "fmt"

type Token struct {
	Type         string
	Value        string
	IntegerValue int
	FloatValue   float32
	LineNumber   int
	Column       int
}

func (token Token) Print() string {
	return fmt.Sprintf(
		"Type: %s, Value: %s, Line: %d, Column: %d",
		token.Type, token.Value, token.LineNumber, token.Column,
	)
}

func (token Token) PrintLineCol() string {
	if token.LineNumber == 0 {
		return ""
	}

	return fmt.Sprintf(
		"Line: %d, Column: %d",
		token.LineNumber, token.Column,
	)
}
