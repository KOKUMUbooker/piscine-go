package piscine

func LoafOfBread(str string) string {
	clean := []rune{}
	for _, r := range str {
		if r != ' ' {
			clean = append(clean, r)
		}
	}

	if len(clean) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0

	for i < len(clean) {
		if i+5 > len(clean) {
			break
		}

		grp := string(clean[i : i+5])
		result += grp

		if i+6 < len(clean) {
			result += " "
		}

		i += 6
	}

	return result + "\n"
}
