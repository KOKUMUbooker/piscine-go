package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListMerge(l1 *List, l2 *List) {
	if l1 == nil || l2 == nil {
		return
	}

	n1 := l1.Tail
	n2 := l2.Head

	for n2 != nil {
		if n1 == nil {
			break
		}
		n1.Next = n2
		n1 = n1.Next
		n2 = n2.Next
	}

	l1.Tail = l2.Tail
}
