package helpers

func ValueInSlice(value string, list []string) bool {
	for _, val := range list {
		if val == value {
			return true
		}
	}

	return false
}
