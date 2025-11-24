package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	// Any num^1 = num
	if power == 1 {
		return nb
	}

	return nb * RecursivePower(nb, power-1)
}
