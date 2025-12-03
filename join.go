package piscine

func Join(strs []string, sep string) string {
	res := ""

	for i, s := range strs {
		res += s
		if i != len(strs)-1 {
			res += sep
		}
	}

	return res
}
