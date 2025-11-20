package piscine

import "github.com/01-edu/z01"

func PrintComb1() {
	first := false

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {

				if i == j && j == k {
					continue
				}

				if i < j && j < k {
					if first { // First time don't prepend ", ", but subsequent cycles prepend it
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

					first = true
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
