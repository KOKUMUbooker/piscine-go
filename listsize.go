package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
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
