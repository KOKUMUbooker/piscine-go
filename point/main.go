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

func main() {
	incrementRuneVal := func(n int) rune {
		r := '0'
		for i := 1; i <= n; i++ {
			r++
		}
		return r
	}

	prinIntRunes := func(n int) {
		base := 10
		mod := 0

		digits := []int{}
		for quot := n; quot > 0; quot /= base {
			mod = quot % base
			digits = append(digits, mod)
		}

		for i := len(digits) - 1; i >= 0; i-- {
			rVal := incrementRuneVal(digits[i])
			z01.PrintRune(rVal)
		}
	}

	points := &point{}

	setPoint(points)

	for _, r := range "x = " {
		z01.PrintRune(r)
	}
	prinIntRunes(points.x)

	for _, r := range ", y = " {
		z01.PrintRune(r)
	}
	prinIntRunes(points.y)

	z01.PrintRune('\n')
}
