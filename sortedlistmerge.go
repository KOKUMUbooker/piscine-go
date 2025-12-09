package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// listMerge(n1,n2)
	nArr := []*NodeI{}

	n := n1
	// Append n1 nodes
	for n != nil {
		nArr = append(nArr, n)
		n = n.Next
	}

	n = n2
	// Append n2 nodes
	for n != nil {
		nArr = append(nArr, n)
		n = n.Next
	}

	if len(nArr) == 0 {
		return nil
	}

	bSort(nArr)

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

	return nArr[0]
}

func bSort(l []*NodeI) {
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
