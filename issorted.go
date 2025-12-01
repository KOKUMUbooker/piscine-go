package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	asc := true
	desc := true

	for i := 0; i < len(a)-1; i++ {
		cmp := f(a[i], a[i+1])

		if cmp > 0 {
			asc = false
		}
		if cmp < 0 {
			desc = false
		}
	}

	return asc || desc
}

func F(a, b int) int {
	if a == b {
		return 0
	}

	if a > b {
		return 1
	}

	return -1
}
