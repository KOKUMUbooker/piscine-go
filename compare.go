package piscine

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
			res := compareRunes(r, aR[i])
			if res != 0 {
				break
			}
		}
		if res == 0 { // if the 2 diff-len strings are equal to the minLen of both, just say a > b
			res = 1
		}

	} else { // strings.Compare("abc", "abcd") = -1
		for i, r := range aR {
			res := compareRunes(r, bR[i])
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

func compareRunes(x, y rune) int {
	if x == y {
		return 0
	} else if x > y {
		return 1
	} else {
		return -1
	}
}
