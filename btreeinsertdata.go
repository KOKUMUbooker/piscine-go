package piscine

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                string
// }

func BTreeInsertDataOg(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}
	r := root
	var parent *TreeNode

	for r != nil {
		parent = r
		if data >= r.Data { // Move to right
			r = r.Right
			continue
		}

		if data < r.Data { // Move to left
			r = r.Left
			continue
		}
	}

	if parent == nil {
		return newNode
	} else if data > parent.Data {
		parent.Right = newNode
	} else {
		parent.Left = newNode
	}

	// Set parent pointer on the new node
	newNode.Parent = parent

	return root
}
