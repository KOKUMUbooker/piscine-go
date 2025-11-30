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

	// "x = "
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// x
	tens := points.x / 10
	ones := points.x % 10

	ch := '0'
	for i := 0; i < tens; i++ {
		ch++
	}
	z01.PrintRune(ch)

	ch = '0'
	for i := 0; i < ones; i++ {
		ch++
	}
	z01.PrintRune(ch)

	// ", y = "
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// points.y (21)
	tens = points.y / 10
	ones = points.y % 10

	ch = '0'
	for i := 0; i < tens; i++ {
		ch++
	}
	z01.PrintRune(ch)

	ch = '0'
	for i := 0; i < ones; i++ {
		ch++
	}
	z01.PrintRune(ch)

	z01.PrintRune('\n')
}
