package piscine

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		// If the list is empty, set both Head and Tail to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Otherwise, update the Tail's Next to the new node and set Tail to the new node
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func main() {
	// Example usage:
	myList := &List{}

	ListPushBack(myList, 1)
	ListPushBack(myList, 2)
	ListPushBack(myList, 3)

	// Print the list to verify the elements have been added at the end
	current := myList.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
