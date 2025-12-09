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
	if l1 == nil && l2 == nil {
		return
	}

	var n1 *NodeL
	var n2 *NodeL

	if l2.Head == nil && l2.Tail == nil {
		return // l2 empty so just return what l1 had
	} else if l1.Head == nil && l1.Tail == nil {
		l1.Head = l2.Head
		n1 = l2.Head
		if l2.Head != nil {
			n2 = l2.Head.Next
		}
	} else {
		n1 = l1.Tail
		n2 = l2.Head
	}

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
