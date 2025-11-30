package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func convIntToStr(n int) string {
	base := 10
	quot := n
	mod := 0

	if n < 0 {
		quot = -n
	}

	digits := []int{}
	for ; quot > 0; quot /= base {
		mod = quot % base
		digits = append(digits, mod)
	}

	digitsR := []rune{}
	for i := len(digits) - 1; i >= 0; i-- {
		r := rune('0' + digits[i])
		digitsR = append(digitsR, r)
	}

	return string(digitsR)
}

func main() {
	points := &point{}

	setPoint(points)

	resStr := "x = "
	xStr := convIntToStr(points.x)
	resStr += xStr

	resStr += ", y = "
	yStr := convIntToStr(points.y)
	resStr += yStr

	for _, r := range resStr {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
