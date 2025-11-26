package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := []string{}
	if len(os.Args) > 1 {
		params = append([]string{}, os.Args[1:]...)
	}

	for _, str := range params {
		for _, r := range str {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
