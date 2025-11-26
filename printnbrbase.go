package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isBaseValid(base) {
		for _, r := range "NV" {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
		return
	}

	baseRunes := []rune(base)
	baseVal := len(baseRunes)
	digits := []int{}
	num := nbr
	isNegative := false

	if num < 0 {
		isNegative = true
		num = -num
	}

	var mod int
	for quot := num; quot != 0; quot /= baseVal {
		mod = quot % baseVal
		digits = append(digits, mod)
	}

	if num == 0 {
		digits = []int{0}
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if isNegative && i == len(digits)-1 {
			z01.PrintRune('-')
		}
		idxInBase := digits[i] // Elements in digits[] will represent indices of corrensponding rune values in base str
		z01.PrintRune(baseRunes[idxInBase])
	}
}

// modValues = [ 5 2 1 ]
// eg nbr=125, base "10"
// 125 / 10 = 12
// 125 % 10 = 5

// 12 / 10 = 1
// 12 % 10 = 2

// 1 / 10 = 0
// 1 % 10 = 1

func isBaseValid(base string) bool {
	rCounts := map[rune]int{}
	var hasPlusOrMinus bool

	if len(base) < 2 {
		return false
	}

	for _, r := range base {
		if r == '-' || r == '+' {
			hasPlusOrMinus = true
		}

		rCounts[r]++
	}

	if hasPlusOrMinus {
		return false
	}

	isValid := true
	for _, counts := range rCounts {
		if counts > 1 {
			isValid = false
			break
		}
	}

	return isValid
}
