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

	for _, r := range "x = " {
		z01.PrintRune(r)
	}

	printTwoDigits(points.x)

	for _, r := range ", y = " {
		z01.PrintRune(r)
	}

	printTwoDigits(points.y)

	z01.PrintRune('\n')
}

func printTwoDigits(n int) {
	base := 10
	tens := n / base
	ones := n % base

	printDigit(tens)

	printDigit(ones)
}

func printDigit(d int) {
	ch := '0'
	for i := 0; i < d; i++ {
		ch++
	}
	z01.PrintRune(ch)
}
