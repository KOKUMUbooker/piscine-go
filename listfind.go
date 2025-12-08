package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	n := l.Head

	for n != nil {
		if CompStr(n.Data, ref) {
			return &n.Data
		}
		n = n.Next
	}

	return nil
}
