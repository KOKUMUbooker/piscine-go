package piscine

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	n := l.Head
	var prevN *NodeL = nil

	for n != nil {
		isFirst := n == l.Head
		isLast := n == l.Tail
		isBetween := n != l.Tail || n != l.Head

		if n.Data == data_ref {
			if isFirst {
				prevN = n
				l.Head = n.Next
				n = n.Next
				continue
			}

			if isLast {
				l.Tail = prevN
				prevN.Next = n.Next
				prevN = n
				n = n.Next
				continue
			}

			if isBetween {
				prevN.Next = n.Next
				n = n.Next
				continue
			}
		}
		prevN = n
		n = n.Next
	}
}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}
