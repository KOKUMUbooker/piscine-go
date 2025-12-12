package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}

	res := ""

	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}

	res += "\n"
	return res
}
