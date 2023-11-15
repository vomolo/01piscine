package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	if l.Data > n.Data {
		n.Next = l
		return n
	}
	iterator := l
	for iterator.Next != nil && iterator.Next.Data < n.Data {
		iterator = iterator.Next
	}
	n.Next = iterator.Next
	iterator.Next = n
	return l
}
