package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	if node.Left == nil && node.Right == nil {
		if node.Parent == nil {
			return nil
		}

		if node == node.Parent.Left {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
		return root
	}

	if node.Left == nil || node.Right == nil {
		var child *TreeNode
		if node.Left != nil {
			child = node.Left
		} else {
			child = node.Right
		}

		if node.Parent == nil {

			child.Parent = nil
			return child
		}

		if node == node.Parent.Left {
			node.Parent.Left = child
		} else {
			node.Parent.Right = child
		}
		child.Parent = node.Parent
		return root
	}

	successor := BTreeMin(node.Right)
	node.Data = successor.Data

	node.Right = BTreeDeleteNode(node.Right, successor)

	return root
}
