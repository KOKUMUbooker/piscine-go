package piscine

func FindNextPrime(nb int) int {
	MaxInt32 := power(2, 31) - 1 // Limit for prime
	if nb < 2 {
		return 2
	}

	prime := nb
	for prime < MaxInt32 {
		if isPrime(prime) {
			return prime
		}
		prime++
	}

	return -1
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

func power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * power(base, exp-1)
}
