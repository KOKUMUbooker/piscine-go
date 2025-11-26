package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	paramsR := []rune{}

	if len(os.Args) > 1 {
		for _, str := range os.Args[1:] {
			if len(str) > 0 {
				paramsR = append(paramsR, rune(str[0]))
			}
		}
	}

	bubbleSort(paramsR)

	for _, r := range paramsR {
		z01.PrintRune(r)
		z01.PrintRune('\n')
	}
}

func bubbleSort(rSlice []rune) {
	n := len(rSlice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if rSlice[j] > rSlice[j+1] {
				rSlice[j], rSlice[j+1] = rSlice[j+1], rSlice[j]
			}
		}
	}
}
