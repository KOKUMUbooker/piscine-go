package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	// f Returns positive int if 1st > 2nd,returns 0 if equal, returns -1 otherwise
	if len(a) < 2 {
		return true
	}

	asc := true
	desc := true

	for i := 0; i < len(a)-1; i++ {
		cmp := f(a[i], a[i+1])

		if cmp > 0 { // Check 1st vs 2nd for ascending order
			asc = false
		}
		if cmp < 0 { // Check 1st vs 2nd for descending order
			desc = false
		}
	}

	return asc || desc
}

func F(a, b int) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}

	return -1
}
