package piscine

/*
1. Split string into runes[]
2. Iterate over rune slice
	- For every rune convert it to int
	- Idea for solution - using Expanded Notation
	  eg 125 = 100 + 20 + 5 = 1
	  	 125 = (1 * 10^2) + (2 * 10^1) + (5 * 10^0)
*/

func BasicAtoi(s string) int {
	var rSlice []rune
	var num int

	for _, r := range s {
		rSlice = append(rSlice, r)
	}

	for i, r := range rSlice {
		digit := convRuneToNum(r)
		exp := len(rSlice) - (i + 1)

		num += digit * pow(10, exp)
	}

	return num
}

func convRuneToNum(r rune) int {
	zeroRune := '0'
	offsetFromZero := r - zeroRune // Assumption being values are positive ie > 0
	return int(offsetFromZero)
}

// calc power using recursion
func pow(n, m int) int {
	if m == 0 {
		return 1
	}

	return n * pow(n, m-1)
}
