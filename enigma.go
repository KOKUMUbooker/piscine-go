package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	aTemp := ***a
	bTemp := *b
	cTemp := *******c
	dTemp := ****d

	// a into c
	*******c = aTemp
	// c into d
	****d = cTemp
	// d into b
	*b = dTemp
	// b into a
	***a = bTemp
}
