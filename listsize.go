package piscine

type NodeL3 struct {
	Data interface{}
	Next *NodeL3
}

type List3 struct {
	Head *NodeL3
	Tail *NodeL3
}

func ListSize(l *List3) int {
	if l.Head == nil {
		return 0
	}

	count := 1
	n := l.Head
	for n.Next != nil {
		count++
		n = n.Next
	}

	return count
}
