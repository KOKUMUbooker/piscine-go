package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	f(root.Data)
	if root.Left != nil {
		BTreeApplyByLevel(root.Left, f)
	}

	if root.Right != nil {
		BTreeApplyByLevel(root.Right, f)
	}
}
