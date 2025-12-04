package piscine

func LoafOfBread(str string) string {
	r := []rune(str)
	if len(r) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0

	for i := 0; i < len(r); i++ {
		if r[i] == ' ' {
			continue
		}

		result += string(r[i])
		count++

		if count == 5 {
			result += " "
			count = -1 // skip next character
		}

		if count == -1 {
			count = 0
			i++ // skip one char
		}
	}

	// Remove the trailing space if it exists
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
