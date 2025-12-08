package piscine

func ListClear(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	nodes := []*NodeL{}
	n := l.Head
	for n.Next != nil {
		nodes = append(nodes, n)
		n = n.Next
	}

	for i := len(nodes) - 1; i >= 0; i-- {
		nI := nodes[i]
		nI.Data = nil
		nI.Next = nil
	}

	l.Head = nil
	l.Tail = nil
	l = nil
}
