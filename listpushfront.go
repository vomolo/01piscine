package piscine

func ListPushFront(l *List, data interface{}) {
	no := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = no
		return
	}
	no.Next = l.Head
	l.Head = no
}
