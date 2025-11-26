package piscine

func ToLower(s string) string {
	var capStr string
	// 'a' - 'A' = 32
	// 'a' -> 'A' = 'a' - ('a'-'A')

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			capR := r + ('a' - 'A')
			capStr += string(capR)
		} else {
			capStr += string(r)
		}
	}

	return capStr
}
