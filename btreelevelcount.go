package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lHeight := 1 + BTreeLevelCount(root.Left)
	rHeight := 1 + BTreeLevelCount(root.Right)

	return max(lHeight, rHeight)
}
