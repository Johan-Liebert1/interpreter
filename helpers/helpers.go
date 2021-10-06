package helpers

import "unicode"

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
