package piscine

func Compact(ptr *[]string) int {
	var str string
	res := []string{}
	count := 0

	for _, s := range *ptr {
		if s == str {
			count++
			continue
		}
		res = append(res, s)
	}

	*ptr = res

	return count
}
