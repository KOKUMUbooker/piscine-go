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

	base := 10
	mod := 0
	count := 0
	tempArr := [2]int{} // Use arrays to avoid using append when using slices
	for quot := points.x; quot > 0; quot /= base {
		mod = quot % base
		tempArr[count] = mod
		count++
	}
	for i := 1; i >= 0; i-- {
		rVal := '0'
		modVal := tempArr[i]
		for j := 1; j <= modVal; j++ {
			rVal++
		}
		z01.PrintRune(rVal)
	}

	for _, r := range ", y = " {
		z01.PrintRune(r)
	}

	// Reset values
	tempArr = [2]int{}
	count = 0
	mod = 0
	for quot := points.y; quot > 0; quot /= base {
		mod = quot % base
		tempArr[count] = mod
		count++
	}
	for i := 1; i >= 0; i-- {
		rVal := '0'
		modVal := tempArr[i]
		for j := 1; j <= modVal; j++ {
			rVal++
		}
		z01.PrintRune(rVal)
	}

	z01.PrintRune('\n')
}
