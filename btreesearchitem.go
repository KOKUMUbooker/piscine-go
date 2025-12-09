package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	r := root
	for r != nil {
		if elem == r.Data {
			return r
		} else if elem > r.Data { // Search right
			r = r.Right
		} else { // Search left
			r = r.Left
		}
	}

	return nil
}
