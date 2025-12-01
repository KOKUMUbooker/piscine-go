package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	// f returns 0 if equal, returns -1 otherwise
	sorted := append([]int{}, a...)

	n := len(sorted)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	for i, n := range a {
		areEqual := f(n, sorted[i])
		if areEqual != 0 {
			return false
		}
	}

	return true
}

func F(a, b int) int {
	if a == b {
		return 0
	}
	return -1
}
