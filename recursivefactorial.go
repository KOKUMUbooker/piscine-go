package piscine

// Factorial - product of all positive ints less that or equal to a number
// eg 5! = 5 * 4 * 3 * 2 *1
func RecursiveFactorial(nb int) int {
	// Factorial of 25 results in a int overflow on the positive side
	// 12! is last factorial that wont cause overflow for 32 bit systems
	// 20! is the last factorial that wont cause overflow fot 64 bit systems
	if nb < 0 || nb >= 21 {
		return 0
	}

	// Base case
	if nb == 1 || nb == 0 {
		return 1
	}

	return nb * RecursiveFactorial(nb-1)
}
