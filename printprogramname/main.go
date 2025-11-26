package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	progName := os.Args[0]
	for _, r := range progName[2:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
