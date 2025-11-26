package piscine

func AtoiBase(s string, base string) int {
	if !isBaseValid2(base) {
		return 0
	}

	// "125", "0123456789" -> 125
	// "7D", "0123456789ABCDEF" -> 125
	digits := []int{}
	baseVal := len(base)
	var num int

	for _,r := range s {
		pos := index(base,r)
		digits = append(digits, pos)
	}	

	for i,digit := range digits {
		op :=  pow5(baseVal,len(digits) - (i + 1))
		num += (digit * op)
	}

	return  1
}

func pow5(base,exp int) int {
	if exp == 0 {
		return 1
	}

	return  base * pow5(base,exp-1)
}

func index(base string, r rune) int {
	pos := -1
	for i,sliceR := range base {
		if sliceR == r {
			pos = i 
			break
		}
	}
	return  pos
}

func isBaseValid2(base string) bool {
	rCounts := map[rune]int{}
	var hasPlusOrMinus bool

	if len(base) < 2 {
		return false
	}

	for _, r := range base {
		if r == '-' || r == '+' {
			hasPlusOrMinus = true
		}

		rCounts[r]++
	}

	if hasPlusOrMinus {
		return false
	}

	isValid := true
	for _, counts := range rCounts {
		if counts > 1 {
			isValid = false
			break
		}
	}

	return isValid
}
