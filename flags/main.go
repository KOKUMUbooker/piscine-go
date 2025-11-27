package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	args = args[1:]
	l := 0
	for i := range args {
		l = i + 1
	}
	if l == 0 || args[0] == "-h" || args[0] == "--help" {
		Usage()
	} else {
		var runes []rune
		order := false
		order_prev := false
		insert := false

		for _, arg := range args {
			length := len([]rune(arg))

			if length > 9 && arg[:9] == "--insert=" {
				runes_temp := []rune(arg[9:])

				for i := 0; i < length-9; i++ {
					runes = append(runes, runes_temp[i])
				}
				insert = true

			} else if length > 3 && arg[:3] == "-i=" {
				runes_temp := []rune(arg[3:])

				for i := 0; i < length-3; i++ {
					runes = append(runes, runes_temp[i])
				}
				insert = true

			} else if (length == 7 && arg[:7] == "--order") || (length == 2 && arg[:2] == "-o") {
				order = true
				order_prev = true
				continue
			} else if order_prev {
				runes_temp := []rune(arg)
				for i := range runes_temp {
					runes = append(runes, runes_temp[i])
				}
			} else {
				runes_print := []rune(arg)
				if order {
					for i := range runes_print {
						runes = append(runes, runes_print[i])
					}
				}
				for i := range runes_print {
					z01.PrintRune(runes_print[i])
				}
			}
			order_prev = false

		}
		if order {
			bubbleSort(runes)
		}

		if order || insert {
			for i := range runes {
				z01.PrintRune(runes[i])
			}
		}

		z01.PrintRune(10)
	}
}

func Usage() {
	fmt.Fprintf(os.Stderr, "--insert\n")
	fmt.Fprintf(os.Stderr, "  -i\n")
	fmt.Fprintf(os.Stderr, "\t This flag inserts the string into the string passed as argument.\n")

	fmt.Fprintf(os.Stderr, "--order\n")
	fmt.Fprintf(os.Stderr, "  -o\n")
	fmt.Fprintf(os.Stderr, "\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func bubbleSort(r []rune) {
	n := len(r)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
}
