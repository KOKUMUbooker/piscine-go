package piscine

// IDEA :
// - Squares get bigger as numbers get bigger
// - Repeatedly increase the number whilst squaring it until you hit the target value or go past it
// - If target is hit -> success
// - If target is passed -> no integer square root exists
func Sqrt(nb int) int {
	if nb < 1 {
		return 0
	}

	res := 1
	// Iterate unitl square of number exceeds nb
	for res*res < nb {
		res++
	}

	if res*res == nb {
		return res
	}

	return 0
}
