package piscine

// piscine.NRune("Ola!", 4)
func NRune(s string, n int) rune {
	if n > len(s) {
		return 0
	}

	iVal := n - 1 // eg n = 2 should mean 2nd letter ie i = 1
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
