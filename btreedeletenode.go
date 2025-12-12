package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func minNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	if node.Data < root.Data {
		// Node to delete in left subtree
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		// Node to delete in right subtree
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		// Found the node to delete

		// Case 1: No child
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// Case 2: One child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Case 3: Two children
		// Find inorder successor (smallest in right subtree)
		successor := minNode(root.Right)
		root.Data = successor.Data
		// Delete the inorder successor
		root.Right = BTreeDeleteNode(root.Right, successor)
	}

	return root
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}

	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	} else if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else if elem == root.Data {
		return root
	} else {
		return root
	}
}

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

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
