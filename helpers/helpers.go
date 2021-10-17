package helpers

import (
	"fmt"
	"reflect"
	"unicode"

	"programminglang/constants"
)

func ValueInSlice(value string, list []string) bool {
	for _, val := range list {
		if val == value {
			return true
		}
	}

	return false
}

func IsAlphaNum(value byte) bool {
	return unicode.IsLetter(rune(value)) || unicode.IsDigit(rune(value))
}

func GetFloat(value interface{}) (float32, bool) {
	v := reflect.ValueOf(value)
	v = reflect.Indirect(v)
	var floatType = reflect.TypeOf(float32(0))

	if v.Type().ConvertibleTo(floatType) {
		return float32(v.Convert(floatType).Float()), true
	}

	return 0.0, false
}

func ColorPrint(color string, newLines int, toPrint ...interface{}) {
	var nl string = ""

	for i := 0; i < newLines; i++ {
		nl += "\n"
	}

	fmt.Print(nl, color)
	fmt.Print(toPrint...)
	fmt.Print(nl, constants.Reset)
}
