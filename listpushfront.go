package piscine

type NodeL2 struct {
	Data interface{}
	Next *NodeL2
}

type List2 struct {
	Head *NodeL2
	Tail *NodeL2
}

func ListPushFront(l *List2, data interface{}) {
	n := &NodeL2{Data: data}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	// Get hold of cur head node
	first := l.Head

	// Update my node next val to prev cur node
	n.Next = first

	// Set my node as the head
	l.Head = n
}
