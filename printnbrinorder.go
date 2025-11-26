package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	digits := extractIntDigits(n)

	bubbleSort(digits)

	for _, d := range digits {
		r := rune('0' + d)
		z01.PrintRune(r)
	}
}

func bubbleSort(digits []int) {
	n := len(digits)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}
}

func extractIntDigits(n int) []int {
	digits := []int{}
	var mod int

	if n == 0 {
		return []int{0}
	}

	for quot := n; quot != 0; quot /= 10 {
		mod = quot % 10
		digits = append(digits, mod)
	}

	return digits
}

// 125 / 10 = 12
// 125 % 10 = 5

// 12 / 10 = 1
// 12 % 10 = 2

// 1 / 10 = 0
// 1 % 10 = 1
