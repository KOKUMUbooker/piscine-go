package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, v := range a {
		counts[v]++
	}

	for _, v := range a {
		// Check for counts with odd values
		if counts[v]%2 != 0 {
			return v
		}
	}

	return -1
}
