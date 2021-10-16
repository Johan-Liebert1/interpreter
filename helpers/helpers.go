package helpers

import (
	"reflect"
	"unicode"
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
