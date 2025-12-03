package piscine

func JoinOld(strs []string, sep string) string {
	var joinedStr string

	for i, str := range strs {
		if i == len(strs)-1 {
			joinedStr += str
		} else {
			joinedStr += str + sep
		}
	}

	return joinedStr
}
