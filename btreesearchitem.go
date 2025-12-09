package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	r := root
	for !(r == nil || r.Data == elem) {
		if elem > r.Data { // Search right
			r = r.Right
		} else { // Search left
			r = r.Left
		}
	}

	return r
}
