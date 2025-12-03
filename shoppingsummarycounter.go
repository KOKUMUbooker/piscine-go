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
	end := 0

	sR := []rune(s)
	sRLen := len(sR)

	res := []string{}
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

		wordR := sR[i:end]
		res = append(res, string(wordR))
		i = end
	}

	return res
}
