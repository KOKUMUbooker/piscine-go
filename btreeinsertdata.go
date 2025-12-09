package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}
	r := root
	var parent *TreeNode

	for r != nil {
		parent = r
		if data > r.Data || data == r.Data { // Move to right
			r = r.Right
			continue
		}

		if data < r.Data { // Move to left
			r = r.Left
			continue
		}
	}

	if parent == nil {
		return &TreeNode{Data: data}
	} else if data > parent.Data {
		parent.Right = newNode
	} else {
		parent.Left = newNode
	}

	return root
}
