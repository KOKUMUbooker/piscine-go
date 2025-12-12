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
