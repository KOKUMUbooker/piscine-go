package piscine

func ConcatParams(args []string) string {
	if args == nil {
		return ""
	}

	concatStr := ""
	for i, str := range args {
		if i == len(args)-1 {
			concatStr += str
		} else {
			concatStr += str + "\n"
		}
	}

	return concatStr
}
