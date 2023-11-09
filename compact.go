package piscine

func Compact(pt *[]string) int {
	var nlist []string
	for _, w := range *pt {
		if w != "" {
			nlist = append(nlist, w)
		}
	}
	*pt = nlist
	return len(nlist)
}
