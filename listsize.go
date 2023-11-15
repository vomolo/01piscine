package piscine

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}
	si := 1
	for l.Head.Next != nil {
		si++
		l.Head = l.Head.Next
	}
	return si
}
