package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "Invalid Output\n"
	}

	sR := []rune(str)
	sRLen := len(sR)
	if sRLen < 5 {
		return "Invalid Output\n"
	}

	sentence := []string{}
	i := 0
	for i < sRLen {
		end := i

		wordR := []rune{}
		for end < sRLen {
			if sR[end] == ' ' && len(sR) != 5 {
				end++
				continue
			}
			wordR = append(wordR, sR[end])
			end++
			if len(wordR) == 5 {
				sentence = append(sentence, string(wordR))
				break
			}
			if end == sRLen {
				sentence = append(sentence, string(wordR))
				break
			}
		}

		i = end + 1
	}

	res := ""
	for i, s := range sentence {
		res += removeSpace(s)
		if i != len(sentence)-1 {
			res += " "
		}
	}

	return res + "\n"
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
