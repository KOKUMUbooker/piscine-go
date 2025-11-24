package piscine

// Factorial - product of all positive ints less that or equal to a number
// eg 5! = 5 * 4 * 3 * 2 *1
func IterativeFactorial(nb int) int {
	// Factorial of 25 results in a int overflow on the positive side
	if nb < 0 || nb >= 25 {
		return 0
	}

	if nb == 1 || nb == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= nb; i++ {
		result = result * i
	}

	return result
}
