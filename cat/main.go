package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	paramsAvailable := len(os.Args) > 1
	if paramsAvailable {
		params := os.Args[1:]

		for i := 0; i < len(params); i++ {
			data, err := os.ReadFile(params[i])
			if err != nil {
				msg := "ERROR: " + err.Error() + "\n"
				printMsg(msg)
				os.Exit(1)
			}
			printMsg(string(data))
		}

		return
	}

	file := os.Stdin
	fileInfo, err := file.Stat()
	if err != nil {
		msg := "ERROR: " + err.Error() + "\n"
		printMsg(msg)
		os.Exit(1)
	}

	if fileInfo.Size() > 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			msg := "ERROR: " + err.Error() + "\n"
			printMsg(msg)
			os.Exit(1)
		}

		printMsg(string(data))
		return
	}

	// printMsg("Hello\n")
}

func printMsg(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
