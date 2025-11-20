package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	first := true

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					// If ij similar to kl skip iteration
					if i == k && j == l {
						continue
					}

					if !first { // Only add ", " if not in first set of iteration
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

					first = false
					z01.PrintRune(i)
					z01.PrintRune(j)

					z01.PrintRune(' ')

					z01.PrintRune(k)
					z01.PrintRune(l)
				}
			}
		}
	}

	z01.PrintRune('\n')
}
