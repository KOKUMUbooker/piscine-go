package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// At the right most end, we'll get the node with largest data
	if root.Right == nil {
		return root
	}

	return BTreeMax(root.Right) // Other wise continue traversing the right
}
