package main

import (
	"fmt"
	"os"
)

const (
	insert = "--insert"
	i      = "-i"
	order  = "--order"
	o      = "-o"
)

func main() {
	params := []string{}

	if len(os.Args) > 1 {
		params = os.Args[1:]
	}

	// If no args provided print usage
	if len(params) == 0 {
		Usage()
		return
	}

	var insertFlag string
	var orderFlag bool
	nonFlagArg := ""
	pLen := len(params)

	for idx, param := range params {
		insFlag := param == insert || param == i
		isOrderFlag := param == order || param == o
		val, iOfEqual, foundEqual := checkStrAfterEqual(param)
		// fmt.Printf("val : %v, iOfEqual : %v, foundEqual : %v\n",val,iOfEqual,foundEqual)

		if insFlag || (foundEqual && (param[:iOfEqual] == i || param[:iOfEqual] == insert)) {
			if foundEqual {
				insertFlag = val
			} else if pLen >= idx+1 && !isFlag(params[idx+1]) {
				insertFlag = params[idx+1]
			}
		} else if isOrderFlag {
			orderFlag = true
		} else if idx >= 0 {
			nonFlagArg = param
		}
	}

	// fmt.Printf("insert flag : %v ,type : %T\n", insertFlag, insertFlag)
	// fmt.Printf("order flag : %v ,type : %T\n", orderFlag, orderFlag)
	// fmt.Printf("nonFlagArg : %v ,type : %T\n", nonFlagArg, nonFlagArg)

	insPresent := len(insertFlag) > 0
	strPresent := len(nonFlagArg) > 0

	strToPrint := ""
	if insPresent && orderFlag && strPresent { // Order the combination of the two
		sR := []rune(nonFlagArg + insertFlag)
		bubbleSort(sR)
		strToPrint = string(sR)
	} else if insPresent && strPresent { // Combine the 2 args
		strToPrint = nonFlagArg + insertFlag
	} else if orderFlag && strPresent { // Order nonFlagArg
		sR := []rune(nonFlagArg)
		bubbleSort(sR)
		strToPrint = string(sR)
	} else { // Print nonFlagArg if its there
		strToPrint = nonFlagArg
	}

	fmt.Fprintf(os.Stdout, "%v\n", strToPrint)
}

func checkStrAfterEqual(str string) (val string, iOfEqual int, found bool) {
	for i, r := range str {
		if r == '=' {
			return str[i+1:], i, true
		}
	}
	return "", -1, false
}

func isFlag(str string) bool {
	return str == o || str == order || str == insert || str == i
}

func Usage() {
	fmt.Fprintf(os.Stderr, "--insert\n")
	fmt.Fprintf(os.Stderr, "  -i\n")
	fmt.Fprintf(os.Stderr, "\t This flag inserts the string into the string passed as argument\n")

	fmt.Fprintf(os.Stderr, "--order\n")
	fmt.Fprintf(os.Stderr, "  -o\n")
	fmt.Fprintf(os.Stderr, "\t This flag will behave like a boolean, if it is called it will order the argument.\n")
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
