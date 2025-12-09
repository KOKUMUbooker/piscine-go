package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	if root.Left != nil {
		BTreeApplyPostorder(root.Left, f)
	}

	if root.Right != nil {
		BTreeApplyPostorder(root.Right, f)
	}

	// Only print we are at the top of the stack & are about to be popped off
	f(root.Data)
}
