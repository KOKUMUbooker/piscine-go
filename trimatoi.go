package piscine

func TrimAtoi(s string) int {
	isNegative := false
	firstIntPassed := false
	digits := []int{}

	for _, r := range s {
		isNum := r >= '0' && r <= '9'
		if !firstIntPassed && r == '-' {
			isNegative = true
		}

		if !firstIntPassed && isNum {
			firstIntPassed = true
		}

		if isNum {
			num := int(r - '0')
			digits = append(digits, num)
		}
	}

	num := 0
	for i, d := range digits {
		offIdx := i + 1
		num += (d * pow4(10, len(digits)-offIdx))
	}

	if isNegative {
		num = -num
	}

	return num
}

func pow4(n, exp int) int {
	if exp == 0 {
		return 1
	}

	return n * pow4(n, exp-1)
}
