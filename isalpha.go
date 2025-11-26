package piscine

func IsAlpha(s string) bool {
	isAlpha := true
	for _, r := range s {
		isAlphaNum := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
		if !isAlphaNum {
			isAlpha = false
			break
		}
	}

	return isAlpha
}
