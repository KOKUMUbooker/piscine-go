package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	// Any num^0 = 1
	if power == 0 {
		return 1
	}

	return nb * RecursivePower(nb, power-1)
}
