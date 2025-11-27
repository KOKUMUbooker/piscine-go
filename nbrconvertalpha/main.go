package main

import (
	"os"

	"github.com/01-edu/z01"
)

// abcdefghijklmnopqrstuvwxyz
func main() {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	aLen := len(alpha)

	params := []string{}
	if len(os.Args) > 1 {
		params = os.Args[1:]
	}

	isUpperFlag := false
	if len(params) > 0 && params[0] == "--upper" {
		isUpperFlag = true
	}

	for _, str := range params {
		var char rune
		if !isNum(str) {
			char = ' '
		} else {
			num := convStrToInt(str) - 1
			if num > aLen {
				char = ' '
			} else {
				r := rune(alpha[num])
				if isUpperFlag {
					char = r - ('a' - 'A')
				} else {
					char = r
				}
			}
		}

		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func isNum(str string) bool {
	isNumeric := true
	for _, r := range str {
		if !(r >= '0' && r <= '9') {
			isNumeric = false
			break
		}
	}

	return isNumeric
}

func convStrToInt(str string) int {
	base := "0123456789"
	num := 0
	digits := []int{}

	for _, r := range str {
		num := baseIdx(r, base)
		digits = append(digits, num)
	}

	for i, n := range digits {
		exp := len(digits) - (i + 1)
		num += n * pow(10, exp)
	}

	return num
}

func baseIdx(br rune, base string) int {
	pos := -1
	for i, r := range base {
		if r == br {
			pos = i
		}
	}

	return pos
}

func pow(base, exp int) int {
	if exp == 0 {
		return 1
	}

	return base * pow(base, exp-1)
}
