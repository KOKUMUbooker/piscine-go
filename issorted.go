package piscine

// Allow for sorting checks either in ascending or descending
func IsSorted(f func(a, b int) int, a []int) bool {
	// f Returns positive int if 1st > 2nd,returns 0 if equal, returns -1 otherwise
	sortedAsc := append([]int{}, a...)
	sortedDesc := append([]int{}, a...)

	n := len(sortedAsc)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if sortedAsc[j] > sortedAsc[j+1] {
				sortedAsc[j], sortedAsc[j+1] = sortedAsc[j+1], sortedAsc[j]
			}
		}
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if sortedDesc[j] < sortedDesc[j+1] {
				sortedDesc[j], sortedDesc[j+1] = sortedDesc[j+1], sortedDesc[j]
			}
		}
	}

	sortedAscBool := true
	sortedDescBool := true
	for i, n := range a {
		cmpAsc := f(n, sortedAsc[i])
		cmpDesc := f(n, sortedDesc[i])

		notEqualAsc := cmpAsc != 0
		notEqualDesc := cmpDesc != 0

		if sortedAscBool && notEqualAsc {
			sortedAscBool = false
		}

		if sortedDescBool && notEqualDesc {
			sortedDescBool = false
		}
	}

	return sortedAscBool || sortedDescBool
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
