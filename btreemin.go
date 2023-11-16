package piscine

import "fmt"

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Traverse to the leftmost leaf
	for root.Left != nil {
		root = root.Left
	}

	return root
}

// BTreeInsertData inserts a new data into the binary tree.
// This function is used in the main function for creating a sample tree.
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
	}

	return root
}

func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")

	min := BTreeMin(root)
	fmt.Println(min.Data)
}
