package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	}

	params := os.Args[1:]
	if len(params) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	data, err := os.ReadFile(params[0])
	if err != nil {
		return
	}

	fmt.Printf(string(data))
}
