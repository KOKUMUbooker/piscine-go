package piscine

func CollatzCountdown(start int) int {
	if start < 1 {
		return -1 // piscine often wants -1 for invalid input
	}

	steps := 0

	for start != 1 {
		if start%2 == 0 { // if even divide by 2
			start = start / 2
		} else { // if odd multiply by 3 and add 1
			start = 3*start + 1
		}
		steps++
	}

	return steps
}
