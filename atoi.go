package piscine

/*
1. Split string into runes[]
2. Iterate over rune slice
    - For every rune convert it to int
    - Idea for solution - using Expanded Notation
      eg 125 = 100 + 20 + 5 = 1
           125 = (1 * 10^2) + (2 * 10^1) + (5 * 10^0)
*/

func Atoi(s string) int {
	rSlice := []rune(s)

	var num int

	isNegative := false
	hasPassedFirstDigit := false

	for i, r := range rSlice {
		isNumber := r >= '0' && r <= '9'

		// If we started collecting numbers & a non-digit character
		// appears in the middle return 0
		if num > 0 && !isNumber {
			return 0
		}

		// If were in the 2nd value and its not a number return 0
		if i == 1 && !isNumber {
			return 0
		}

		if !isNumber {
			continue
		}

		// Safely check for '-' on the immediate previous value
		if !hasPassedFirstDigit && i == 1 && rSlice[i-1] == '-' {
			isNegative = true
		}

		hasPassedFirstDigit = true

		digit := int(r - '0')
		exp := len(rSlice) - (i + 1)

		num += digit * pow3(10, exp)
	}

	if isNegative {
		return -num
	}

	return num
}

// calc power using recursion
func pow3(n, m int) int {
	if m == 0 {
		return 1
	}

	return n * pow3(n, m-1)
}
