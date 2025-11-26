package piscine

func IsUpper(s string) bool {
	isUpper := true
	for _, r := range s {
		if !(r >= 'A' && r <= 'Z') {
			isUpper = false
		}
	}

	return isUpper
}
