package piscine

func Unmatch(a []int) int {
	res := -1

	for i := 0; i < len(a); i++ {
		present := false
		for j := 0; j < len(a); j++ {
			if j != i && a[i] == a[j] {
				present = true
			}
		}

		if !present {
			return a[i]
		}
	}

	return res
}
