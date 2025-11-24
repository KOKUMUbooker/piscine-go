package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}

	prime := nb
	for {
		if isPrime(prime) {
			break
		}
		prime++
	}

	return prime
}

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	// Iterate over only odd numbers & check whether one of the iteration values
	// can divide nb without a remainder
	// If found value has more divisors than two(1 & the num itself), thus making it not a prime
	for i := 2; i < nb-1; i++ {
		if nb%i == 0 {
			return false
		}
	}

	return true
}
