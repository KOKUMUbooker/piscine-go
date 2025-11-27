package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Change default usage output that is usually provided by flag.PrintDefaults() which is called by flag.Usage()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "--insert\n")
		fmt.Fprintf(os.Stderr, "  -i\n")
		fmt.Fprintf(os.Stderr, "\t This flag inserts the string into the string passed as argument\n")

		fmt.Fprintf(os.Stderr, "--order\n")
		fmt.Fprintf(os.Stderr, "  -o\n")
		fmt.Fprintf(os.Stderr, "\t This flag will behave like a boolean, if it is called it will order the argument.\n")
	}

	// If no args provided print usage
	if len(os.Args) <= 1 {
		flag.Usage()
		return
	}

	var insert string
	flag.StringVar(&insert, "insert", "", "") // pointer, cmd-name, default-value, usage
	flag.StringVar(&insert, "i", "", "")

	var order bool
	flag.BoolVar(&order, "order", false, "") // pointer, cmd-name, default-value, usage
	flag.BoolVar(&order, "o", false, "")

	flag.Parse() // Parse flags before accessing them

	if !flag.Parsed() {
		flag.Usage()
		return
	}

	nonFlagArg := flag.Arg(0)
	// fmt.Printf("insert flag : %v ,type : %T\n",insert,insert)
	// fmt.Printf("order flag : %v ,type : %T\n",order,order)
	// fmt.Printf("nonFlagArg : %v ,type : %T\n",nonFlagArg,nonFlagArg)

	insPresent := len(insert) > 0
	strPresent := len(nonFlagArg) > 0

	strToPrint := ""
	if insPresent && order && strPresent { // Order the combination of the two
		sR := []rune(nonFlagArg + insert)
		bubbleSort(sR)
		strToPrint = string(sR)
	} else if insPresent && strPresent { // Combine the 2 args
		strToPrint = nonFlagArg + insert
	} else if order && strPresent { // Order nonFlagArg
		sR := []rune(nonFlagArg)
		bubbleSort(sR)
		strToPrint = string(sR)
	} else { // Print nonFlagArg if its there
		strToPrint = nonFlagArg
	}

	fmt.Fprintf(os.Stdout, "%v\n", strToPrint)
}

func bubbleSort(r []rune) {
	n := len(r)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
}
