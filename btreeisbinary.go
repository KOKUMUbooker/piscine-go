package piscine

func BTreeIsBinaryOg(root *TreeNode) bool {
	return checkIsBSTOg(root, nil, nil)
}

func checkIsBSTOg(node *TreeNode, min *string, max *string) bool {
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
	return checkIsBSTOg(node.Left, min, &val) &&
		checkIsBSTOg(node.Right, &val, max)
}
