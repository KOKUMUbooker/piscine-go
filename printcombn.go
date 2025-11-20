package piscine

import "github.com/01-edu/z01"

func PrintCombN() {
	isFirst := true

	// Loop 1: First digit (i) - Iterates 10 times ('0' to '9')
	for i := '0'; i <= '9'; i++ {
		// Loop 2: Second digit (j) - Iterates 10 times ('0' to '9')
		for j := '0'; j <= '9'; j++ {
			// Loop 3: Third digit (k) - Iterates 10 times ('0' to '9')
			for k := '0'; k <= '9'; k++ {
				if i < j && j < k {

					if !isFirst {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

					isFirst = false

					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
