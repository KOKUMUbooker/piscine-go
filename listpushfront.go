package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}

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
