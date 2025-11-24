package piscine

// Factorial - product of all positive ints less that or equal to a number
// eg 5! = 5 * 4 * 3 * 2 *1
func IterativeFactorial(nb int) int {
	if nb <= 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}

	result := 1
	for i := 1; i <= nb; i++ {
		result = result * i
	}

	return result
}
