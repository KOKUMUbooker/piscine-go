package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// In ascii table lower case letters start from 97 to 122
	for i := 97; i <= 122; i++ {
		z01.PrintRune(rune(i))
	}
}
