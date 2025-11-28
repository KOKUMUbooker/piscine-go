package piscine

func Split(s, sep string) []string {
	sepStr := []string{}
	sR := []rune(s)
	sepR := []rune(sep)

	sRLen := len(sR)
	spLen := len(sepR)

	if len(sep) == 0 {
		return []string{s}
	}

	i := 0
	end := 0

	for i < sRLen {
		r := sR[i]
		pendLen := len(sR[i:])
		pendHoldsSep := pendLen >= len(sep)

		if r == sepR[0] && pendHoldsSep && string(sR[i:i+spLen]) == sep {
			i += spLen
			end += spLen
			continue
		}

		for {
			if end >= sRLen {
				break
			}

			eR := sR[end]
			ePendLen := len(sR[end:])
			ePendHoldsSep := ePendLen >= len(sep)

			if eR == sepR[0] && ePendHoldsSep && string(sR[end:end+spLen]) == sep {
				break
			}
			end++
		}

		word := string(sR[i:end])
		sepStr = append(sepStr, word)
		i = end
	}

	return sepStr
}
