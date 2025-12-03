package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	sentence := splitString(str)

	res := map[string]int{}

	for _, word := range sentence {
		res[word]++
	}

	return res
}

func splitString(s string) []string {
	i := 0

	sR := []rune(s)
	sRLen := len(sR)

	res := []string{}
	for i < sRLen {
		end := i + 1
		for !isUpper(sR[end]) {
			end++
			if end == sRLen {
				break
			}
		}

		wordR := sR[i:end]
		word := removeSpace(string(wordR))
		res = append(res, word)
		if end == sRLen {
			break
		}
		i = end
	}

	return res
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func removeSpace(s string) string {
	res := ""
	for _, r := range s {
		if r != ' ' {
			res += string(r)
		}
	}

	return res
}
