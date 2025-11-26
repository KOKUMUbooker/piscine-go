package piscine

// Capitalize("~Z0xLxj~^\\2.d") == "~Z0Xlxj~^\\2.D" instead of "~Z0xlxj~^\\2.D"
func Capitalize(s string) string {
	sR := []rune(s)
	var prevSeparator rune
	prevSeparatorIdx := -1
	capRunes := []rune{}

	for i, r := range sR {
		isUpper := (r >= 'A' && r <= 'Z')
		isLower := (r >= 'a' && r <= 'z')
		isNum := (r >= '0' && r <= '9')
		prevRuneIsSeparator := i-prevSeparatorIdx == 1 && !isAlphaNum(prevSeparator)

		if isLower {
			l := r
			if prevRuneIsSeparator {
				l = toUpper(l)
			}
			capRunes = append(capRunes, l)
		} else if isUpper {
			l := r
			if !prevRuneIsSeparator {
				l = toLower(l)
			}
			capRunes = append(capRunes, l)
		} else if isNum {
			capRunes = append(capRunes, r)
		} else {
			capRunes = append(capRunes, r)

			prevSeparator = r
			prevSeparatorIdx = i
		}

	}

	return string(capRunes)
}

func isAlphaNum(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	} else {
		return false
	}
}

func toUpper(r rune) rune {
	return r - ('a' - 'A')
}

func toLower(r rune) rune {
	return r + ('a' - 'A')
}
