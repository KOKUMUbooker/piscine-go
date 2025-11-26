package piscine

// eg
// strings.Index("chicken", "ken") = 4
// strings.Index("chicken", "dmr") = -1

// fmt.Println(piscine.Index("Hello!", "l"))
// fmt.Println(piscine.Index("Salut!", "alu"))
// fmt.Println(piscine.Index("Ola!", "hOl"))
func Index(s string, toFind string) int {
	if len(s) == 0 && len(toFind) == 0 {
		return 0
	}

	if len(toFind) == 0 {
		return 0
	}

	sRSlice := []rune(s)
	tFRSlice := []rune(toFind)
	matchRSlice := []rune{}

	index := -1

	for i := 0; i < len(sRSlice); i++ {
		sR := sRSlice[i]
		matchLen := len(matchRSlice) // Should point to the next value that should be picked in tFRSlice thus no -1 needed

		tFR := tFRSlice[matchLen]

		if sR == tFR {
			matchRSlice = append(matchRSlice, sR)

			if len(matchRSlice) == len(tFRSlice) {
				// 2 = 2-(1-1)
				index = i - (len(matchRSlice) - 1)
				break
			}
		} else if sR != tFR && len(matchRSlice) > 0 {
			matchRSlice = []rune{} // Delete everything in the matchSlice if subsequent chars don't match
		}
	}

	return index
}
