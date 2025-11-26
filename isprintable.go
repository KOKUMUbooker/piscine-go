package piscine

func IsPrintable(s string) bool {
	isPrintable := true
	for _, r := range s {
		isPrintableRune := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
		if !isPrintableRune {
			isPrintable = false
			break
		}
	}

	return isPrintable
}
