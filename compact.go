package piscine

func Compact(ptr *[]string) int {
	unique := map[string]int{}

	for _, s := range *ptr {
		unique[s]++
	}

	var str string
	res := []string{}
	for k := range unique {
		if k != str {
			res = append(res, k)
		}
	}

	*ptr = res

	return unique[str]
}
