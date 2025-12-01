package piscine

func Map(f func(int) bool, a []int) []bool {
	res := []bool{}
	for _, n := range a {
		res = append(res, f(n))
	}

	return res
}
