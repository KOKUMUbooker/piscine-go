package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	paramsAvailable := len(os.Args) > 1
	if paramsAvailable {
		params := os.Args[1:]

		for i := 0; i < len(params); i++ {
			data, err := os.ReadFile(params[i])
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				return
			}
			fmt.Printf(string(data))
		}

		return
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		return
	}

	fmt.Printf(string(data))
}
