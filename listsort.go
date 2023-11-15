package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}
	s := l
	f := l
	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	second := s.Next
	s.Next = nil
	return MergeSort(ListSort(l), ListSort(second))
}

func MergeSort(l1 *NodeI, l2 *NodeI) *NodeI {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Data < l2.Data {
		l1.Next = MergeSort(l1.Next, l2)
		return l1
	}
	l2.Next = MergeSort(l1, l2.Next)
	return l2
}
