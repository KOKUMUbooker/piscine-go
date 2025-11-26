package piscine

// 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b string) int {
	res := 0
	aR := []rune(a)
	bR := []rune(b)

	if a == "" && b == "" {
		res = 0
	} else if len(aR) == len(bR) {
		for i, r := range aR {
			cmp := compareRunes(r, bR[i])
			if cmp != 0 {
				res = cmp
				break
			}
		}
	} else if len(aR) > len(bR) { // strings.Compare("abcd", "abc") = +1
		for i, r := range bR {
			res = compareRunes(aR[i], r)
			if res != 0 {
				break
			}
		}
		if res == 0 { // if the 2 diff-len strings are equal to the minLen of both, just say a > b
			res = 1
		}

	} else { // strings.Compare("abc", "abcd") = -1
		for i, r := range aR {
			res = compareRunes(r, bR[i])
			if res != 0 {
				break
			}
		}
		if res == 0 { // if the 2 diff-len strings are equal to the minLen of both, just say a < b
			res = -1
		}
	}

	return res
}

func compareRunes(a, b rune) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
