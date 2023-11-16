package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}

	leftResult := BTreeSearchItem(root.Left, elem)
	if leftResult != nil {
		return leftResult
	}

	rightResult := BTreeSearchItem(root.Right, elem)
	if rightResult != nil {
		return rightResult
	}

	return nil
}
