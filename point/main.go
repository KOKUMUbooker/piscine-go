package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	ToRune(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	ToRune(points.y)
	z01.PrintRune('\n')
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n/10 != 0 {
		printInt(n / 10)
	}

	// Print the last digit
	digit := n % 10
	ch := '0'
	for i := 0; i < digit; i++ {
		ch++
	}
	z01.PrintRune(ch)
}

func ToRune(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	printInt(n)
}
