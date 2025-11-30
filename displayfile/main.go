package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
	}

	params := os.Args[1:]
	if len(params) > 1 {
		fmt.Println("Too many arguments")
	}

	data, err := os.ReadFile(params[0])
	if err != nil {
		return
	}

	fmt.Printf(string(data))
}
