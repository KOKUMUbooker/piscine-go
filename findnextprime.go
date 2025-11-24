package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}

	prime := nb
	for !isPrime(prime) {
		prime++
	}

	return prime
}

func isPrime(nb int) bool {
	// No prime numbers before 2
	if nb < 2 {
		return false
	}
	// Return early for 2
	if nb == 2 {
		return true
	}
	// Even no. greater than 2 get excluded
	if nb%2 == 0 {
		return false
	}

	// "i*i <= nb" similar to i <= sqrt(nb)
	// Only loop up to the square root of the value
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
