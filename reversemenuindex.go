package piscine

func ReverseMenuIndex(menu []string) []string {
	l := len(menu)
	out := make([]string, l)

	for i := l - 1; i >= 0; i-- {
		idx := l - (i + 1)
		out[idx] = menu[i]
	}

	return out
}
