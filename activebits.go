package piscine

func ActiveBits(n int) int {
	base := 2

	if n == 0 {
		return 0
	}

	count := 0
	for quot := n; quot > 0; quot /= base {
		mod := quot % base
		if mod == 1 {
			count++
		}
	}

	return count
}
