package piscine

func ToUpper(s string) string {
	var capStr string
	// 'a' - 'A' = 32
	// 'a' -> 'A' = 'a' - ('a'-'A')

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			capR := r - ('a' - 'A')
			capStr += string(capR)
		} else {
			capStr += string(r)
		}
	}

	return  capStr
}
