package piscine

func StringToIntSlice(str string) []int {
	if len(str) == 0 {
		return nil
	}

	result := []int{}

	for _, r := range str {
		result = append(result, int(r))
	}

	return result
}
