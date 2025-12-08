package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	// Get last node address
	last := l.Tail

	// Update the last node's next value to our node
	last.Next = n

	// Update the list's tail value to our node
	l.Tail = n
}
