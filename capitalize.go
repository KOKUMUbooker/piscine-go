package piscine

func Capitalize(s string) string {
	sR := []rune(s)
	var prevSeparator rune
	prevSeparatorIdx := -1
	capRunes := []rune{}

	for i, r := range sR {
		isUpper := (r >= 'A' && r <= 'Z')
		isLower := (r >= 'a' && r <= 'z')
		prevRuneIsSeparator := i-prevSeparatorIdx == 1 && !isAlpha(prevSeparator)

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
		} else {
			capRunes = append(capRunes, r)

			prevSeparator = r
			prevSeparatorIdx = i
		}

	}

	return string(capRunes)
}

func isAlpha(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
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
