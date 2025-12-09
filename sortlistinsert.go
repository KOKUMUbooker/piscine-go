package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newN := &NodeI{Data: data_ref}
	if l == nil {
		return newN
	}

	nArr := []*NodeI{}
	n := l

	for n != nil {
		nArr = append(nArr, n)
		n = n.Next
	}

	// Insert new node in nArr
	nArr = append(nArr, newN)

	bubbleS(nArr)

	for i := 0; i < len(nArr); i++ {
		nextI := i + 1
		curN := nArr[i]

		if nextI < len(nArr) {
			nextN := nArr[nextI]
			curN.Next = nextN
		} else {
			curN.Next = nil
		}
	}

	if len(nArr) == 0 {
		return nil
	}

	return nArr[0]
}

func bubbleS(l []*NodeI) {
	n := len(l)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			a := l[j]
			b := l[j+1]
			if a.Data > b.Data {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
}
