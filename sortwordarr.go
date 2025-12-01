package piscine

func SortWordArr(a []string) {
	n := len(a)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if isGreater(a[j], a[j+1]) {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func isGreater(a, b string) bool {
	if len(a) == 0 && len(b) == 0 {
		return false
	}

	aR := []rune(a)
	bR := []rune(b)
	if len(a) == 1 || len(b) == 1 {
		return aR[0] > bR[0]
	}

	if len(aR) > len(bR) {
		equal := true
		for i, r := range b {
			if aR[i] != r {
				equal = false
				return aR[i] > r
			}
		}
		if equal { // If character up to minIndex ie len(bR) are same in both, return true since a is longer
			return true
		}
	} else {
		equal := true
		for i, r := range a {
			if r != bR[i] {
				equal = false
				return r > bR[i]
			}
		}
		if equal { // If character up to minIndex ie len(aR) are same in both, return true since b is longer
			return false
		}
	}

	return false
}
