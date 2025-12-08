package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}

	n := l
	cur := 0
	found := false

	for n != nil {
		if cur == pos {
			found = true
			break
		}

		cur++
		n = n.Next
	}

	if found {
		return n
	}

	return nil
}
