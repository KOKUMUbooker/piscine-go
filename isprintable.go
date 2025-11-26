package piscine

func IsPrintable(s string) bool {
	isPrintable := true
	for _, r := range s {
		isPrintableRune := r >= '\u0020' && r <= '\u007f'
		if !isPrintableRune {
			isPrintable = false
			break
		}
	}

	return isPrintable
}
