package piscine

func Rot14(s string) string {
	res := ""

	for _, r := range s {
		if isUpper(r) {
			targetCh := getNextChar(r, 'Z', 'A')
			res += string(targetCh)
		} else if isLower(r) {
			targetCh := getNextChar(r, 'z', 'a')
			res += string(targetCh)
		} else {
			res += string(r)
		}
	}

	return res
}

func getNextChar(r, upper, lower rune) rune {
	targ := r + 14

	if targ > upper {
		// amount past the upper bound
		over := targ - upper - 1
		return lower + over // wrap around from lower
	}

	return targ
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}
