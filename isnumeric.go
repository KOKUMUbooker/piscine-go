package piscine

func IsNumeric(s string) bool {
	isNum := true

	for _, r := range s {
		if !(r >= '0' && r <= '9') {
			isNum = false
			break
		}
	}

	return isNum
}
