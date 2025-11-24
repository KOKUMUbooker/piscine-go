package piscine

func IterativePower(nb int, power int) int {
	if nb <= 0 || power <= 0 {
		return 0
	}

	// Any num^1 = num
	if power == 1 {
		return nb
	}

	result := 1
	for i := 1; i <= power; i++ {
		result = result * nb
	}

	return result
}
