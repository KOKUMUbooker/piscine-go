package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}

	// Prioritise printing left most nodes before all others
	f(root.Data)

	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}
}
