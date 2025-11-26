package piscine

func BasicJoin(elems []string) string {
	var joinedStr string

	for _, str := range elems {
		joinedStr += str
	}

	return joinedStr
}
