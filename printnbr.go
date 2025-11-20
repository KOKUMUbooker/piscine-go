package piscine

import "github.com/01-edu/z01"

/*
	1. Recieve int param
	2. Break int into individual digits
	3. Convert individual digits into rune
	4. Iteratively call PrinRune() on the array of runes
*/

func PrintNbr(n int) {
	splitDigits := splitIntToDigits(n) // Digits will be in reverse
	splitDigitsRunes := make([]rune, 0)

	if n < 0 {
		splitDigitsRunes = append(splitDigitsRunes, '-') // Prepend '-' for negative numbers
	}

	for i := len(splitDigits) - 1; i >= 0; i-- {
		r := rune('0' + splitDigits[i])
		splitDigitsRunes = append(splitDigitsRunes, r)
	}

	for _, r := range splitDigitsRunes {
		z01.PrintRune(r)
	}
}

/*
Splitting int to digits is achieved using 2 expressions - division by 10 & modulus by 10
  - These 2 operation are repeated until the quotient of the division expression is 0
  - Modulus expression gives the split digit under that iteration
*/
func splitIntToDigits(n int) []int {
	digits := make([]int, 0)
	quot := n
	mod := 0

	if n == 0 {
		return []int{0}
	}

	// Do this until quotient is zero
	for quot != 0 {
		mod = -(quot % 10) // Use '-' to address potential overflow issue arising from doing -(n)
		digits = append(digits, mod)
		quot = quot / 10
	}

	return digits
}
