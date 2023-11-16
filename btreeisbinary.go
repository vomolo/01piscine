package piscine

func BTreeIsBinary(root *TreeNode) bool {
	var prevData string
	var isValid bool = true

	var inOrderTraversal func(node *TreeNode)
	inOrderTraversal = func(node *TreeNode) {
		if node != nil {

			inOrderTraversal(node.Left)

			if node.Data <= prevData {
				isValid = false
				return
			}

			prevData = node.Data

			inOrderTraversal(node.Right)
		}
	}

	inOrderTraversal(root)

	return isValid
}
