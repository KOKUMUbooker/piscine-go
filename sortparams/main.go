package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := []string{}

	if len(os.Args) > 1 {
		for _, str := range os.Args[1:] {
			if len(str) > 0 {
				params = append(params, str)
			}
		}
	}

	bubbleSort(params)

	for _, str := range params {
		for _, r := range str {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

// Returns true if a > b
func isGreater(a, b string) bool {
	res := true

	if len(a) == 1 && len(b) == 1 {
		res = rune(a[0]) > rune(b[0])
	} else if len(a) == len(b) {
		for i, ar := range a {
			br := rune(b[i])
			if ar != br {
				res = ar > br
				break
			}
		}
	} else if len(a) > len(b) {
		for i, br := range b {
			ar := rune(b[i])
			if ar != br {
				res = ar > br
				break
			}
		}
	} else if len(a) < len(b) {
		for i, ar := range a {
			br := rune(b[i])
			if ar != br {
				res = ar > br
				break
			}
			if res {
				res = false
			}
		}
	}

	return res
}

func bubbleSort(rSlice []string) {
	n := len(rSlice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if isGreater(rSlice[j], rSlice[j+1]) {
				rSlice[j], rSlice[j+1] = rSlice[j+1], rSlice[j]
			}
		}
	}
}
