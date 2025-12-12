package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	nodeData := LevelOrder(root)

	for _, nStr := range nodeData {
		f(nStr)
	}
}

func LevelOrder(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := []string{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Data)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}
