package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1:]
	for _, s := range args {
		if s == "01" || s == "galaxy" || s == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
