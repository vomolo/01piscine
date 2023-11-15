package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	t := l.Head
	p := l.Head

	for t != nil && t.Data == data_ref {
		l.Head = t.Next
		t = l.Head
	}
	for t != nil {
		if t.Data != data_ref {
			p = t
		}
		p.Next = t.Next
		t = p.Next
	}
}
