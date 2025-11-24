package piscine

// Fibonacci - sequence of series where each num is a sum of 2 preceeding FIBONACCI NUMBERS
// FORMULA : f(n) = f(n-1) + f(n-2) for n > 1
func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}

	// Fib[1] = 1
	if index == 1 {
		return 1
	}

	// Fib[0] = 0
	if index == 0 {
		return 0
	}

	// Will get factorials for F[2] and onwards
	return Fibonacci(index-1) + Fibonacci(index-2)
}
