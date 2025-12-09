package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}

	nArr := []*NodeI{}

	n := l
	for n != nil {
		nArr = append(nArr, n)
		n = n.Next
	}

	bbleSort(nArr)

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

func bbleSort(l []*NodeI) {
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
