package main

func Split(s string) []string {
	i := 0
	end := 0

	sR := []rune(s)
	sRLen := len(sR)
	res := []string{}

	for i < sRLen {
		if sR[i] == '(' {
			enclosedStr := getEnclosedStr(sR, i)
			i += len(enclosedStr)
			end += len(enclosedStr)
			continue
		}

		if sR[i] == ' ' {
			i++
			end++
			continue
		}

		for sR[i] != ' ' {
			end++
			if end == sRLen {
				break
			}
		}

		word := string(sR[i:end]);
		res = append(res, word)
		i = end
	}

	return res
}

func getEnclosedStr(sR []rune, i int) string {
	end := i;
	for end < len(sR) {
		if sR[end] == ')' {
			break
		}
		end++
	}

	return string(sR[i : end+1])
}
