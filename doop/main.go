package main

import (
	"fmt"
	"os"
)

// For -> invalid operator, value, number of arguments
// or an overflow print nothing
// For -> the modulo and division operations by 0 print
// "No division by 0" or "No modulo by 0"
func main() {
	if len(os.Args) != 4 {
		return
	}

	val1 := os.Args[1]
	op := os.Args[2]
	val2 := os.Args[3]

	if !isNum(val1) || !isNum(val2) || !isValidOp(op) {
		return
	}

	num1, n1Overflow := convStrToInt(val1)
	num2, n2Overflow := convStrToInt(val2)
	if n1Overflow || n2Overflow {
		return
	}

	res := 0
	switch op {
	case "+":
		out, overflow := add64(num1, num2)
		if overflow {
			return
		}
		res = out
	case "-":
		out, overflow := sub64(num1, num2)
		if overflow {
			return
		}
		res = out
	case "/":
		if num2 == 0 {
			// os.Stdin.WriteString("No division by 0\n")
			fmt.Println("No division by 0")
			return
		}
		res = num1 / num2
	case "*":
		out, overflow := mul64(num1, num2)
		if overflow {
			return
		}
		res = out
	case "%":
		if num2 == 0 {
			// os.Stdin.WriteString("No modulo by 0\n")
			fmt.Println("No modulo by 0")
			return
		}
		res = num1 % num2
	default:
		return
	}

	fmt.Println(convIntToStr(res))
	// os.Stdin.WriteString(convIntToStr(res) + "\n")
}

func add64(a, b int) (int, bool) {
	sum := a + b
	overflow := (a > 0 && b > 0 && sum < 0) || (a < 0 && b < 0 && sum > 0)
	return sum, overflow
}

func sub64(a, b int) (int, bool) {
	diff := a - b
	overflow := (a > 0 && b < 0 && a-b < 0) || (a < 0 && b > 0 && a-b > 0)
	return diff, overflow
}

func mul64(a, b int) (int, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}
	prod := a * b
	overflow := prod/b != a
	return prod, overflow
}

func isNum(nS string) bool {
	for i, r := range nS {
		if i == 0 && r == '-' {
			continue
		}
		if !(r >= '0' && r <= '9') {
			return false
		}
	}

	return true
}

func isValidOp(op string) bool {
	if len(op) != 1 {
		return false
	}

	if !(op == "+" || op == "-" || op == "/" || op == "*" || op == "%") {
		return false
	}

	return true
}

func convIntToStr(n int) string {
	digits := []int{}
	if n == 0 {
		return "0"
	}
	numStr := ""
	if n < 0 {
		n = -n
		numStr += "-"
	}
	for quot := n; quot > 0; quot /= 10 {
		mod := quot % 10
		digits = append(digits, mod)
	}

	for i := len(digits) - 1; i >= 0; i-- {
		nD := digits[i]
		nR := rune('0' + nD)
		numStr += string(nR)
	}

	return numStr
}

func convStrToInt(nS string) (int, bool) {
	const maxInt64 = 9223372036854775807
	if len(nS) == 0 {
		return 0, true
	}

	neg := false
	idx := 0

	if nS[0] == '-' {
		neg = true
		idx = 1
	}

	var num int64 = 0

	for ; idx < len(nS); idx++ {
		ch := nS[idx]

		if ch < '0' || ch > '9' {
			return 0, true
		}

		digit := int64(ch - '0')

		if num > (maxInt64-digit)/10 {
			return 0, true // overflow
		}

		num = num*10 + digit
	}

	if neg {
		num = -num
		if num > 0 {
			return 0, true
		}
	}

	return int(num), false
}
