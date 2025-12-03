package piscine

func Abort(a, b, c, d, e int) int {
	l := []int{a, b, c, d, e}
	n := len(l)

	// Sort in ascending using bubble sort
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}

	return l[n/2]
}
