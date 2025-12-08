package piscine

func ListClear(l *List) {
	if l == nil {
		return
	}

	// Walk through the list whilst breaking all links
	for n := l.Head; n != nil; {
		next := n.Next
		n.Data = nil
		n.Next = nil
		n = next
	}

	// Reset the list
	l.Head = nil
	l.Tail = nil
}
