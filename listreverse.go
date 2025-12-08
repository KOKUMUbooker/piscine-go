package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListReverse(l *List) {
	lArr := []*NodeL{}

	n := l.Head
	for n != nil {
		lArr = append(lArr, n)
		n = n.Next
	}

	for i := len(lArr) - 1; i >= 0; i-- {
		nd := lArr[i]
		if i == len(lArr)-1 {
			nextI := i - 1
			if nextI >= 0 {
				nd.Next = lArr[nextI]
			} else {
				nd.Next = nil
			}

			l.Head = nd
		} else if i == 0 {
			nd.Next = nil
			l.Tail = nd
		} else {
			nextI := i - 1
			if nextI >= 0 {
				nd.Next = lArr[nextI]
			} else {
				nd.Next = nil
			}

			l.Tail = nd
		}
	}
}
