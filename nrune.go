package piscine

func NRune(s string, n int) rune {
	if n > len(s)-1 {
		return 0
	}

	iVal := n
	if n < 0 {
		iVal = len(s) + n
	}

	var runeVal rune
	for i, r := range s {
		if i == iVal {
			runeVal = r
			break
		}
	}

	return runeVal
}
