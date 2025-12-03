package piscine

func ShoppingListSort(slice []string) []string {
	bbSort(slice)

	return slice
}

func bbSort(s []string) {
	n := len(s)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			a := len(s[j])
			b := len(s[j+1])
			if a > b {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
