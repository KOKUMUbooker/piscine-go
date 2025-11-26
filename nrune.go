package piscine

// NRune("Whatshisname", -2) == 'm' instead of '\x00'
// piscine.NRune("Ola!", 4)
func NRune(s string, n int) rune {
	if n > len(s) || n < -1 {
		return 0
	}

	iVal := n - 1 // eg n = 2 should mean 2nd letter ie i = 1
	if n == -1 {
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
