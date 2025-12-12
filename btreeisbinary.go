package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return checkIsBST(root, nil, nil)
}

func checkIsBST(node *TreeNode, min *string, max *string) bool {
	if node == nil {
		return true
	}

	val := node.Data

	// Check against min bound
	if min != nil && val <= *min {
		return false
	}

	// Check against max bound
	if max != nil && val >= *max {
		return false
	}

	// Check subtrees with updated bounds using the val
	return checkIsBST(node.Left, min, &val) &&
		checkIsBST(node.Right, &val, max)
}
