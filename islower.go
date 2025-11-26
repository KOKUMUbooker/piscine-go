package piscine

func IsLower(s string) bool {
	isLower := true
	for _, r := range s {
		if !(r >= 'a' && r <= 'z') {
			isLower = false
			break
		}
	}

	return isLower
}
