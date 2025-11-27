package piscine

func SplitWhiteSpaces(s string) []string {
	sR := []rune(s)
	sRLen := len(sR)
	strSlice := []string{}

	i := 0
	end := 0

	for i < sRLen {
		if sR[i] == ' ' {
			i++
			end++
			continue
		}

		for sR[end] != ' ' {
			end++
			if end == sRLen {
				break
			}
		}

		word := ""
		if end == sRLen {
			word = string(sR[i:end])
		} else {
			word = string(sR[i:end])
		}

		strSlice = append(strSlice, word)
		i = end
	}

	return strSlice
}
