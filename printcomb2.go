package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	first := true

	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)

		for j := '0'; j <= '9'; j++ {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

			first = false
			z01.PrintRune(j)
		}
	}

	z01.PrintRune('\n')
}
