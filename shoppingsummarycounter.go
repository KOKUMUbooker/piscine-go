package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	words := splitString(str)
	res := make(map[string]int)

	for _, w := range words {
		res[w]++
	}

	return res
}

func splitString(s string) []string {
	res := []string{}
	current := ""

	for _, r := range s {
		if r == ' ' {
			res = append(res, current)
			current = ""
		} else {
			current += string(r)
		}
	}

	res = append(res, current)

	return res
}
