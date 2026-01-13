package main

import (
	"os"
	"io/fs"
	"fmt"

	// "piscine"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	inDir := os.Args[1]
	// outDir := os.Args[2]

	root := os.DirFS(".")

	input, err := fs.ReadFile(root, inDir)
	if err != nil {
		fmt.Println("Error occured while reading file : ", err)
		return
	}
	fmt.Println("Input : ", string(input))

	// 1. Split sentence into words with space as separator
	sentence := Split(string(input))

	fmt.Println()
	fmt.Println("Sentence Arr ", len(sentence), " : ", sentence)

	// 2. Initialize string array for the result
	res := []string{}

	// 3. Loop through array of words, looking for mod keys & if found call utility functions on the previous word
	for i, word := range sentence {
		if i == 0 {
			continue
		}

		switch word {
		case "(hex)":
		case "(bin)":
		case "(up)":
		case "(low)":
		case "(cap)":
		}
	}

	fmt.Println("Res : ",res)

	// 4. Join the result string with space as separator
}